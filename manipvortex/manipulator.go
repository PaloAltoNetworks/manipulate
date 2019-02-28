package manipvortex

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
	"go.aporeto.io/manipulate"
	"go.uber.org/zap"
)

// updater is the type of all crud functions.
type updater func(mctx manipulate.Context, objects ...elemental.Identifiable) error

// vortexManipulator is a Vortex based on the memdb implementation.
type vortexManipulator struct {
	upstreamManipulator     manipulate.Manipulator
	upstreamSubscriber      manipulate.Subscriber
	downstreamManipulator   manipulate.Manipulator
	model                   elemental.ModelManager
	processors              map[string]*Processor
	commitIdentityEvent     map[string]struct{}
	subscribers             []*vortexSubscriber
	transactionQueue        chan *Transaction
	enableLog               bool
	logfile                 string
	logChannel              chan *Transaction
	defaultReadConsistency  manipulate.ReadConsistency
	defaultWriteConsistency manipulate.WriteConsistency
	defaultQueueDuration    time.Duration
	pageSize                int
	prefetcher              Prefetcher

	sync.RWMutex
}

// New will create a new cache. Caller must provide a valid
// backend manipulator and susbscriber. If the manipulator is nil, it will be assumed
// that the cache is standalone (ie there is no backend to synchronize with).
func New(
	ctx context.Context,
	downstreamManipulator manipulate.Manipulator,
	processors map[string]*Processor,
	model elemental.ModelManager,
	options ...Option,
) (manipulate.BufferedManipulator, error) {

	if downstreamManipulator == nil {
		panic("downstreamManipulator must not be nil")
	}

	if model == nil {
		panic("model must not be nil")
	}

	if len(processors) == 0 {
		panic("processors must not be empty")
	}

	cfg := newConfig()
	for _, option := range options {
		option(cfg)
	}

	m := &vortexManipulator{
		downstreamManipulator:   downstreamManipulator,
		upstreamManipulator:     cfg.upstreamManipulator,
		upstreamSubscriber:      cfg.upstreamSubscriber,
		defaultReadConsistency:  cfg.readConsistency,
		defaultWriteConsistency: cfg.writeConsistency,
		enableLog:               cfg.enableLog,
		logfile:                 cfg.logfile,
		pageSize:                cfg.defaultPageSize,
		prefetcher:              cfg.prefetcher,
		processors:              processors,
		model:                   model,
		transactionQueue:        cfg.transactionQueue,
		defaultQueueDuration:    cfg.defaultQueueDuration,
		subscribers:             []*vortexSubscriber{},
		commitIdentityEvent:     map[string]struct{}{},
	}

	if m.enableLog {
		c, err := newLogWriter(ctx, m.logfile, 100)
		if err != nil {
			return nil, fmt.Errorf("unable open commit log file: %s", err)
		}
		m.logChannel = c
	}

	if m.prefetcher != nil {
		if err := m.warmUp(ctx); err != nil {
			return nil, fmt.Errorf("unable to warm up: %s", err)
		}
	}

	if m.upstreamSubscriber != nil {

		filter := elemental.NewPushFilter()
		for identity, cfg := range m.processors {
			if cfg.CommitOnEvent {
				m.commitIdentityEvent[identity] = struct{}{}
			}
			filter.FilterIdentity(identity)
		}

		m.upstreamSubscriber.Start(ctx, filter)

		go m.monitor(ctx)
	}

	// Start the background thread. It will be blocked
	// when we do resyncs and this is ok. We want it blocked
	// so that resync continues while any updates are buffered.
	go m.backgroundSync(ctx)

	return m, nil
}

// Flush implements the flush interface of the Vortex. It will flush
// all the cache for write-through. For write-back it will wait
// for a maximum of 10 seconds for transactions to complete. When
// done it will flush the channel and create a completely fresh
// db.
func (m *vortexManipulator) Flush(ctx context.Context) error {

	m.Lock()
	defer m.Unlock()

	if m.prefetcher != nil {
		m.prefetcher.Flush()
	}

	f, ok := m.downstreamManipulator.(manipulate.FlushableManipulator)
	if !ok {
		return nil
	}

	// Wait for the channel to clean up
	maxDelay := time.Now().Add(10 * time.Second)
	for len(m.transactionQueue) > 0 && time.Now().Before(maxDelay) {
		time.Sleep(1 * time.Second)
	}

	// Flush any outstanding transactions and restart the backgrond sync
	if err := f.Flush(ctx); err != nil {
		return fmt.Errorf("unable to flush the datastore: %s", err)
	}

	return nil
}

func (m *vortexManipulator) RetrieveMany(mctx manipulate.Context, dest elemental.Identifiables) error {

	m.RLock()
	defer m.RUnlock()

	if mctx == nil {
		mctx = manipulate.NewContext(context.Background())
	}

	if m.prefetcher != nil {

		prefetched, err := m.prefetcher.Prefetch(mctx.Context(), elemental.OperationRetrieveMany, dest.Identity(), m.upstreamManipulator, mctx.Derive())
		if err != nil {
			return fmt.Errorf("unable to prefetch data for retrieve many operation for '%s': %s", dest.Identity(), err)
		}

		if err := m.insertPrefetchedData(prefetched); err != nil {
			return fmt.Errorf("unable to insert prefetched data for retrieve many operation for '%s': %s", dest.Identity(), err)
		}
	}

	if !m.shouldProcess(mctx, dest.Identity()) {
		if m.upstreamManipulator != nil {
			return m.upstreamManipulator.RetrieveMany(mctx, dest)
		}
		return nil
	}

	cfg := m.processors[dest.Identity().Name]

	if cfg.RetrieveManyHook != nil {
		commit, err := cfg.RetrieveManyHook(m.downstreamManipulator, mctx, dest)
		if !commit {
			return err
		}
	}

	return m.downstreamManipulator.RetrieveMany(mctx, dest)
}

func (m *vortexManipulator) Retrieve(mctx manipulate.Context, objects ...elemental.Identifiable) error {

	m.RLock()
	defer m.RUnlock()

	if len(objects) == 0 {
		return nil
	}

	if !isCommonIdentity(objects...) {
		return fmt.Errorf("all objects in operation must be of the same identity")
	}

	if mctx == nil {
		mctx = manipulate.NewContext(context.Background())
	}

	if m.prefetcher != nil {
		prefetched, err := m.prefetcher.Prefetch(mctx.Context(), elemental.OperationRetrieve, objects[0].Identity(), m.upstreamManipulator, mctx.Derive())
		if err != nil {
			return fmt.Errorf("unable to prefetch data for retrieve operation for '%s': %s", objects[0].Identity(), err)
		}
		if err := m.insertPrefetchedData(prefetched); err != nil {
			return fmt.Errorf("unable to insert prefetched data for retrieve operation for '%s': %s", objects[0].Identity(), err)
		}
	}

	// If we are not processing the object or the object has a parent
	// send it upstream. We only deal with CRUDs.
	if !m.shouldProcess(mctx, objects[0].Identity()) {
		if m.upstreamManipulator != nil {
			return m.upstreamManipulator.Retrieve(mctx, objects...)
		}
		return nil
	}

	if err := m.downstreamManipulator.Retrieve(mctx, objects...); err != nil {

		// If we can't find it locally, and its strong consistency retrieve
		// we will try the backend if we have one.
		if m.upstreamManipulator != nil && (mctx != nil && mctx.ReadConsistency() == manipulate.ReadConsistencyStrong) {

			if err := m.upstreamManipulator.Retrieve(mctx, objects...); err != nil {
				return err
			}

			// Make sure that we update our cache for future reference.
			if err := m.downstreamManipulator.Create(mctx, objects...); err != nil {
				return fmt.Errorf("unable to update local cache from backend: %s", err)
			}

			return nil
		}

		return err
	}

	return nil
}

func (m *vortexManipulator) Create(mctx manipulate.Context, objects ...elemental.Identifiable) error {

	m.RLock()
	defer m.RUnlock()

	return m.coreCRUDOperation(elemental.OperationCreate, mctx, objects...)
}

func (m *vortexManipulator) Update(mctx manipulate.Context, objects ...elemental.Identifiable) error {

	m.RLock()
	defer m.RUnlock()

	return m.coreCRUDOperation(elemental.OperationUpdate, mctx, objects...)
}

func (m *vortexManipulator) Delete(mctx manipulate.Context, objects ...elemental.Identifiable) error {

	m.RLock()
	defer m.RUnlock()

	return m.coreCRUDOperation(elemental.OperationDelete, mctx, objects...)
}

func (m *vortexManipulator) DeleteMany(mctx manipulate.Context, identity elemental.Identity) error {

	m.RLock()
	defer m.RUnlock()

	if mctx == nil {
		mctx = manipulate.NewContext(context.Background())
	}

	if m.upstreamManipulator == nil {
		return fmt.Errorf("delete many not supported by vortexManipulator")
	}

	return m.upstreamManipulator.DeleteMany(mctx, identity)
}

func (m *vortexManipulator) Count(mctx manipulate.Context, identity elemental.Identity) (int, error) {

	m.RLock()
	defer m.RUnlock()

	if mctx == nil {
		mctx = manipulate.NewContext(context.Background())
	}

	if m.downstreamManipulator == nil {
		return 0, fmt.Errorf("datastore is not initialized")
	}

	return m.downstreamManipulator.Count(mctx, identity)
}

func (m *vortexManipulator) hasBackendSubscriber() bool {

	m.RLock()
	defer m.RUnlock()

	return m.upstreamSubscriber != nil
}

func (m *vortexManipulator) registerSubscriber(s manipulate.Subscriber) {

	m.Lock()
	defer m.Unlock()

	m.subscribers = append(m.subscribers, s.(*vortexSubscriber))
}

// UpdateFilter updates the current filter.
func (m *vortexManipulator) updateFilter() {

	m.RLock()
	defer m.RUnlock()

	if m.upstreamSubscriber == nil {
		return
	}

	m.commitIdentityEvent = map[string]struct{}{}

	filter := elemental.NewPushFilter()
	for identity := range m.processors {
		m.commitIdentityEvent[identity] = struct{}{}
		filter.FilterIdentity(identity)
	}

	for _, subscriber := range m.subscribers {
		subscriber.RLock()
		for callerIdentity := range subscriber.filter.Identities {

			cfg, ok := m.processors[callerIdentity]
			if ok {
				// If we are processing this event and there is a client
				// subscription, we will only commit if the corresponding
				// flag is set. Otherwise, the client will have to handle
				// the update, so we remove it from the list.
				if !cfg.CommitOnEvent {
					delete(m.commitIdentityEvent, callerIdentity)
				}
				continue
			}
			// If it is not one of the registered identites and the client
			// has subscribed anyway, we still register and forward it to the
			// client.
			filter.FilterIdentity(callerIdentity)
		}
		subscriber.RUnlock()
	}

	// Update the downstream filter.
	m.upstreamSubscriber.UpdateFilter(filter)
}

// coreCRUDOperation implements the basic operation for updates of the db. This is used
// by create, update, and delete.
func (m *vortexManipulator) coreCRUDOperation(operation elemental.Operation, mctx manipulate.Context, objects ...elemental.Identifiable) error {

	if !isCommonIdentity(objects...) {
		return fmt.Errorf("all objects in operation must be of the same identity")
	}

	if mctx == nil {
		mctx = manipulate.NewContext(context.Background())
	}

	// If the identity is not registered or the request has a parent
	// send upstream. We are not dealing with this locally.
	if !m.shouldProcess(mctx, objects[0].Identity()) {
		return m.commitUpstream(mctx.Context(), operation, mctx, objects...)
	}

	reconcile, err := m.genericUpdater(operation, mctx, objects...)
	if err != nil {
		return err
	}
	if !reconcile {
		return nil
	}

	return m.commitLocal(operation, mctx, objects)
}

// shouldProcess returns true if the request can be processed by the cache. If false,
// it must be forwarded to the upstream.
func (m *vortexManipulator) shouldProcess(mctx manipulate.Context, identity elemental.Identity) bool {

	_, ok := m.processors[identity.Name]
	if !ok {
		return false
	}

	return mctx == nil || (mctx != nil && mctx.Parent() == nil)
}

// commitUpstream will commit a transaction to the upstream if it is not nil. It will
// return the upstream error.
func (m *vortexManipulator) commitUpstream(ctx context.Context, method elemental.Operation, mctx manipulate.Context, objects ...elemental.Identifiable) error {

	if m.upstreamManipulator == nil {
		return nil
	}

	// If it is managed object we apply the pre-hook.
	cfg, ok := m.processors[objects[0].Identity().Name]
	if ok {
		reconcile, err := m.processHook(method, cfg.UpstreamHook, mctx, objects...)
		if !reconcile {
			return err
		}
	}

	// We always commit if prehook says ok or it is not a managed object.
	if err := manipulate.Retry(
		ctx,
		func() error {
			return m.methodFromType(method)(mctx, objects...)
		},
		nil,
	); err != nil {
		return err
	}

	return nil
}

// commitLocal will commit a transaction locally after processing any
// hooks. It will return error if either the hook or the local commit
// fail for some reason.
func (m *vortexManipulator) commitLocal(method elemental.Operation, mctx manipulate.Context, objects []elemental.Identifiable) error {

	if objects == nil || len(objects) == 0 {
		return nil
	}

	cfg, ok := m.processors[objects[0].Identity().Name]
	if !ok {
		return nil
	}

	reconcile, err := m.processHook(method, cfg.LocalHook, mctx, objects...)
	if !reconcile {
		return err
	}

	if err := m.localMethodFromType(method)(mctx, objects...); err != nil {
		return err
	}

	if m.enableLog {
		m.logChannel <- &Transaction{
			Date:    time.Now(),
			mctx:    mctx,
			Objects: objects,
			Method:  method,
		}
	}

	return nil
}

// localMethodFromType will return a pointer to the corresponding function
// based  on the elemental method type.
func (m *vortexManipulator) localMethodFromType(method elemental.Operation) updater {

	switch method {

	case elemental.OperationCreate:
		return m.downstreamManipulator.Create

	case elemental.OperationUpdate:
		return m.downstreamManipulator.Update

	default:
		return m.downstreamManipulator.Delete
	}
}

// methodFromType it will return an upstream function pointer based on the method.
func (m *vortexManipulator) methodFromType(method elemental.Operation) updater {

	switch method {

	case elemental.OperationCreate:
		return m.upstreamManipulator.Create

	case elemental.OperationUpdate:
		return m.upstreamManipulator.Update

	default:
		return m.upstreamManipulator.Delete
	}
}

func (m *vortexManipulator) processHook(method elemental.Operation, hook Hook, mctx manipulate.Context, objects ...elemental.Identifiable) (reconcile bool, err error) {

	if hook != nil {
		return hook(method, mctx, objects)
	}

	return true, nil
}

// genericUpdate will implement the updates. It takes as parameters the methods
// to be used (update, create, delete) and avoids repeating code. It will
// return true if the transaction has to be committed in the local DB. It will
// return an error if the backend fails. Specifically:
// For WriteThrough: it will return an error if the backend fails.
// For WriteBack it will cache it and return commit=false. The commit will happen
// later after the object is stored in the backend.
func (m *vortexManipulator) genericUpdater(method elemental.Operation, mctx manipulate.Context, objects ...elemental.Identifiable) (bool, error) {

	if m.upstreamManipulator == nil {
		return true, nil
	}

	// We are guaranteed that there is at least one object and the identity is processable.
	cfg := m.processors[objects[0].Identity().Name]

	wc := cfg.WriteConsistency
	if mctx.WriteConsistency() != manipulate.WriteConsistencyDefault {
		wc = mctx.WriteConsistency()
	} else if wc == manipulate.WriteConsistencyDefault || wc == "" {
		wc = m.defaultWriteConsistency
	}

	tdeadline := cfg.QueueingDuration
	if tdeadline == 0 {
		tdeadline = m.defaultQueueDuration
	}

	switch wc {

	case manipulate.WriteConsistencyStrong, manipulate.WriteConsistencyStrongest:
		// In Strong consistency we make sure that the backend gets the create.
		// Only then store in the cache.
		return true, m.commitUpstream(mctx.Context(), method, mctx, objects...)

	default:

		select {

		case m.transactionQueue <- &Transaction{
			mctx:     mctx,
			Objects:  objects,
			Method:   method,
			Deadline: time.Now().Add(tdeadline),
		}:
			return false, nil

		default:
			return false, fmt.Errorf("commit queue is full: %d", len(m.transactionQueue))
		}
	}
}

// backgroundSync will empty the transaction queue and try to sync it
// with the backend.
func (m *vortexManipulator) backgroundSync(ctx context.Context) {

	if m.upstreamManipulator == nil {
		return
	}

	for {
		select {
		case t := <-m.transactionQueue:

			// If the dealine is exceeded we just drop the request
			// no matter what. This allows us to clean up the queue
			// if there is a problem.
			if time.Now().After(t.Deadline) {
				continue
			}

			if len(t.Objects) == 0 {
				continue
			}

			// We first try to update the backend. If this succeeds
			// then we also update the local db. At this point
			// the object can be accessible through our API since
			// the ID has been populated.
			m.RLock()

			if _, ok := m.processors[t.Objects[0].Identity().Name]; !ok {
				m.RUnlock()
				continue
			}

			retryCtx, cancel := context.WithDeadline(ctx, t.Deadline)
			cancel()

			if err := m.commitUpstream(retryCtx, t.Method, t.mctx, t.Objects...); err != nil {
				m.RUnlock()
				zap.L().Error("failed to commit object upstream", zap.Error(err))
				continue
			}

			// Update the local copy of the object now.
			if err := m.commitLocal(t.Method, t.mctx, t.Objects); err != nil {
				zap.L().Error("failed to delete local object after failed resync", zap.Error(err))
			}

			m.RUnlock()

		case <-ctx.Done():

			// TODO: If we get killed with objects in the queue, then what ?
			// Do we ignore it and try to empty all objects or what ????
			return
		}
	}
}

// monitor registers for events for all the identities of interest
// and keeps the local cache up-to-date with the backend.
func (m *vortexManipulator) monitor(ctx context.Context) {

	for {

		select {

		case evt := <-m.upstreamSubscriber.Events():

			m.RLock()
			_, commit := m.commitIdentityEvent[evt.Identity]
			m.RUnlock()

			if commit {
				m.eventHandler(ctx, evt)
			}

			m.pushEvent(evt)

		case err := <-m.upstreamSubscriber.Errors():
			zap.L().Error("Received error from the push channel", zap.Error(err))
			// Push event upstream.
			m.pushErrors(err)

		case status := <-m.upstreamSubscriber.Status():

			switch status {

			case manipulate.SubscriberStatusDisconnection:
				zap.L().Warn("Upstream event channel interrupted. Reconnecting...")

			case manipulate.SubscriberStatusInitialConnection:
				zap.L().Info("Upstream event channel connected")

			case manipulate.SubscriberStatusReconnection:
				zap.L().Info("Upstream event channel restored")

				// We flush everything
				if err := m.Flush(ctx); err != nil {
					zap.L().Error("Unable to flush", zap.Error(err))
				}

				// We flush everything
				if m.prefetcher != nil {
					if err := m.warmUp(ctx); err != nil {
						zap.L().Error("Unable to warm up after reconnection", zap.Error(err))
					}
				}

			case manipulate.SubscriberStatusFinalDisconnection:
				return
			}

			m.pushStatus(status)

		case <-ctx.Done():
			return
		}
	}
}

func (m *vortexManipulator) pushEvent(evt *elemental.Event) {

	for _, s := range m.subscribers {
		sevent, err := copystructure.Copy(evt)
		if err != nil {
			zap.L().Error("failed to copy event", zap.Error(err))
			continue
		}

		if !s.filter.IsFilteredOut(evt.Identity, evt.Type) {
			select {
			case s.subscriberEventChannel <- sevent.(*elemental.Event):
			default:
				zap.L().Error("Subscriber channel is full")
			}
		}
	}
}

func (m *vortexManipulator) pushStatus(status manipulate.SubscriberStatus) {

	for _, s := range m.subscribers {
		select {
		case s.subscriberStatusChannel <- status:
		default:
			zap.L().Error("Subscriber channel is full")
		}
	}
}

func (m *vortexManipulator) pushErrors(err error) {
	for _, s := range m.subscribers {
		select {
		case s.subscriberErrorChannel <- err:
		default:
			zap.L().Error("Subscriber channel is full")
		}
	}
}

func (m *vortexManipulator) eventHandler(ctx context.Context, evt *elemental.Event) {

	if m.upstreamManipulator == nil {
		return
	}

	obj := m.model.IdentifiableFromString(evt.Identity)

	if err := evt.Decode(obj); err != nil {
		zap.L().Error("Unable to unmarshal received event", zap.Error(err))
		return
	}

	m.RLock()
	defer m.RUnlock()

	// DO WE FORCE COMPLETE RESYNCS HERE IF THERE ARE FAILURES?
	// ERROR HANDLING NEEDS WORK. Since errors here are extremely
	// unlikely, provided that the schema is correct, probably the
	// right thing to do is to force a re-sync or completely panic.
	var method elemental.Operation

	switch evt.Type {

	case elemental.EventCreate:
		method = elemental.OperationCreate

	case elemental.EventUpdate:
		method = elemental.OperationUpdate

	case elemental.EventDelete:
		method = elemental.OperationDelete

	default:
		zap.L().Error("unsupported event received", zap.String("Event", string(evt.Type)))
		return
	}

	if err := m.commitLocal(method, nil, elemental.IdentifiablesList{obj}); err != nil {
		if method != elemental.OperationDelete {
			zap.L().Error("failed to commit locally an event notification", zap.String("event", evt.String()), zap.Error(err))
		}
	}
}

func (m *vortexManipulator) insertPrefetchedData(prefetched []elemental.Identifiables) error {

	for _, identifiables := range prefetched {

		if identifiables == nil {
			continue
		}

		lst := identifiables.List()
		if len(lst) == 0 {
			continue
		}

		if err := m.commitLocal(elemental.OperationCreate, nil, lst); err != nil {
			return err
		}
	}

	return nil
}

func (m *vortexManipulator) warmUp(ctx context.Context) error {

	if m.upstreamManipulator == nil {
		return nil
	}

	for _, proc := range m.processors {

		prefetched, err := m.prefetcher.WarmUp(ctx, m.upstreamManipulator, m.model, proc.Identity)
		if err != nil {
			return fmt.Errorf("unable to prefetch '%s for warm up operation: %s", proc.Identity, err)
		}

		if err := m.insertPrefetchedData(append([]elemental.Identifiables{}, prefetched)); err != nil {
			return fmt.Errorf("unable to insert prefetched '%s' for warm up operation: %s", proc.Identity, err)
		}
	}

	return nil
}
