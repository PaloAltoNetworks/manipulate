package memdbvortex

import (
	"context"
	"fmt"
	"sync"
	"time"

	"go.aporeto.io/elemental"
	"go.aporeto.io/manipulate"
	"go.aporeto.io/manipulate/manipmemory"
	"go.aporeto.io/manipulate/manipvortex"
	"go.aporeto.io/manipulate/manipvortex/config"
	"go.uber.org/zap"
)

// MemDBVortex is a Vortex based on the memdb implementation.
type MemDBVortex struct {
	m         manipulate.Manipulator
	s         manipulate.Subscriber
	datastore *MemdbDatastore

	model  elemental.ModelManager
	memory manipulate.TransactionalManipulator

	processors          map[string]*config.ProcessorConfiguration
	commitIdentityEvent map[string]struct{}

	subscriberErrorChannel  chan error
	subscriberEventChannel  chan *elemental.Event
	subscriberStatusChannel chan manipulate.SubscriberStatus

	started          bool
	transactionQueue chan *Transaction

	enableLog  bool
	logfile    string
	logChannel chan *Transaction

	sync.RWMutex
}

// updater is the type of all crud functions.
type updater func(mctx manipulate.Context, objects ...elemental.Identifiable) error

// Transaction is the event that captures the transaction for later processing. It is
// also the structure stored in the transaction logs.
type Transaction struct {
	Date     time.Time
	mctx     manipulate.Context
	Objects  []elemental.Identifiable
	Method   elemental.Operation
	Deadline time.Time
}

// NewMemDBVortex will create a new cache. Caller must provide a valid
// backend manipulator and susbscriber. If the manipulator is nil, it will be assumed
// that the cache is standalone (ie there is no backend to synchronize with).
func NewMemDBVortex(
	datastore *MemdbDatastore,
	backendManipulator manipulate.Manipulator,
	backendSubscriber manipulate.Subscriber,
	processors map[string]*config.ProcessorConfiguration,
	model elemental.ModelManager,
	logfile string,
) manipvortex.Cache {

	return &MemDBVortex{
		m:                       backendManipulator,
		s:                       backendSubscriber,
		model:                   model,
		datastore:               datastore,
		processors:              processors,
		transactionQueue:        make(chan *Transaction, 1000),
		subscriberErrorChannel:  make(chan error, 10),
		subscriberEventChannel:  make(chan *elemental.Event, 100),
		subscriberStatusChannel: make(chan manipulate.SubscriberStatus, 10),
		enableLog:               logfile != "",
		logfile:                 logfile,
	}
}

// Run starts the memory cache.
func (v *MemDBVortex) Run(ctx context.Context) error {
	v.Lock()
	defer v.Unlock()

	if v.started {
		return fmt.Errorf("DB is already started")
	}

	if v.enableLog {
		c, err := newLogWriter(ctx, v.logfile, 100)
		if err != nil {
			return fmt.Errorf("cannot open commit log file")
		}
		v.logChannel = c
	}

	if v.datastore == nil || !v.datastore.IsInitialized() {
		return fmt.Errorf("Datastore is nil or not initialized")
	}

	v.memory = manipmemory.NewMemoryManipulatorFromDB(v.datastore.GetDB())

	if v.s != nil {
		go v.monitor(ctx)
	}
	// Start the background thread. It will be blocked
	// when we do resyncs and this is ok. We want it blocked
	// so that resync continues while any updates are buffered.
	go v.backgroundSync(ctx)

	// Do a complete DB resync at this point to download any objects.
	// Note that we are locked down. Any updates coming will be
	// queued waiting for us to finish and they will apply
	// after that. There is a possible race condition here
	// where our read gets a newer object than a pending update.
	// Only way to resolve is to use update times.
	if err := v.resync(ctx); err != nil {
		return err
	}

	// You should not start this thing twice or havoc might happen.
	v.started = true

	return nil
}

// Flush implements the flush interface of the Vortex. It will flush
// all the cache for write-through. For write-back it will wait
// for a maximum of 10 seconds for transactions to complete. When
// done it will flush the channel and create a completely fresh
// db.
func (v *MemDBVortex) Flush(ctx context.Context) error {
	v.Lock()
	defer v.Unlock()

	// Wait for the channel to clean up
	maxDelay := time.Now().Add(10 * time.Second)
	for len(v.transactionQueue) > 0 && time.Now().Before(maxDelay) {
		time.Sleep(1 * time.Second)
	}

	// Flush any outstanding transactions and restart the backgrond sync
	close(v.transactionQueue)
	v.transactionQueue = make(chan *Transaction, 1000)

	if err := v.datastore.Flush(); err != nil {
		return fmt.Errorf("failed to flush the datastore; %s", err)
	}

	v.memory = manipmemory.NewMemoryManipulatorFromDB(v.datastore.GetDB())

	// Restart the background process on the channel.
	go v.backgroundSync(ctx)

	return nil
}

// ReSync implements the ReSync interface of the vortex.
func (v *MemDBVortex) ReSync(ctx context.Context) error {
	v.Lock()
	defer v.Unlock()

	// Call the internal resync after we lock. Need to make
	// sure that everyone is blocked while doing resync.
	return v.resync(ctx)

}

// resync is an internal resync that assumes the caller will
// take the locks. It is called from various places where
// callers already have the lock.
func (v *MemDBVortex) resync(ctx context.Context) error {

	if v.m == nil {
		return nil
	}

	if v.datastore == nil {
		return fmt.Errorf("cannot resync - datastore is not initialized")
	}

	if err := v.datastore.Flush(); err != nil {
		return fmt.Errorf("failed to flush the datastore; %s", err)
	}

	v.memory = manipmemory.NewMemoryManipulatorFromDB(v.datastore.GetDB())

	for _, cfg := range v.processors {

		if err := v.migrateObject(ctx, cfg); err != nil {
			return err
		}
	}

	return nil
}

// RetrieveMany implements the manipulate interface. We are retrieving
// from the cache only.
func (v *MemDBVortex) RetrieveMany(mctx manipulate.Context, dest elemental.Identifiables) error {

	v.RLock()
	defer v.RUnlock()

	if !v.isProcessable(mctx, dest.Identity()) {
		if v.m != nil {
			return v.m.RetrieveMany(mctx, dest)
		}
		return nil
	}

	cfg := v.processors[dest.Identity().Name]

	if cfg.RetrieveManyHook != nil {
		commit, err := cfg.RetrieveManyHook(v.memory, mctx, dest)
		if !commit {
			return err
		}
	}

	return v.memory.RetrieveMany(mctx, dest)
}

// Retrieve implements the manipulate interface. We are retrieving
// from the cache first, unless  if strong consistency is requested
// in which case, we can retrieve from the main node.
func (v *MemDBVortex) Retrieve(mctx manipulate.Context, objects ...elemental.Identifiable) error {

	v.RLock()
	defer v.RUnlock()

	if len(objects) == 0 {
		return nil
	}

	if !v.isCommonIdentity(objects...) {
		return fmt.Errorf("all objects in operation must be of the same identity")
	}

	// If we are not processing the object or the object has a parent
	// send it upstream. We only deal with CRUDs.
	if !v.isProcessable(mctx, objects[0].Identity()) {
		if v.m != nil {
			return v.m.Retrieve(mctx, objects...)
		}
		return nil
	}

	if err := v.memory.Retrieve(mctx, objects...); err != nil {
		// If we can't find it locally, and its strong consistency retrieve
		// we will try the backend if we have one.
		if v.m != nil && (mctx != nil && mctx.ReadConsistency() == manipulate.ReadConsistencyStrong) {
			if err := v.m.Retrieve(mctx, objects...); err != nil {
				return err
			}
			// Make sure that we update our cache for future reference.
			if err := v.memory.Create(mctx, objects...); err != nil {
				return fmt.Errorf("failed to update local cache from backend: %s", err)
			}
			return nil
		}
		return err
	}

	return nil
}

// Create implements the manipulate interface for object creation.
// Depending on the cache mode it will return immediately or after is
// synchronized.
func (v *MemDBVortex) Create(mctx manipulate.Context, objects ...elemental.Identifiable) error {

	v.RLock()
	defer v.RUnlock()

	return v.coreCRUDOperation(elemental.OperationCreate, mctx, objects...)
}

// Update implements the manipulate interface for object updates.
// Depending on the cache mode it will return immediately or after
// it is synchronized.
func (v *MemDBVortex) Update(mctx manipulate.Context, objects ...elemental.Identifiable) error {

	v.RLock()
	defer v.RUnlock()

	return v.coreCRUDOperation(elemental.OperationUpdate, mctx, objects...)
}

// Delete implements the manipulate interface for object deletes.
// Depending on the cache mode it will return immediately or after
// it is synchronized.
func (v *MemDBVortex) Delete(mctx manipulate.Context, objects ...elemental.Identifiable) error {

	v.RLock()
	defer v.RUnlock()

	return v.coreCRUDOperation(elemental.OperationDelete, mctx, objects...)
}

// DeleteMany implements the corresponding interface method.
func (v *MemDBVortex) DeleteMany(mctx manipulate.Context, identity elemental.Identity) error {

	v.RLock()
	defer v.RUnlock()

	if v.m != nil {
		return v.m.DeleteMany(mctx, identity)
	}

	return fmt.Errorf("delete many not supported by memdbvortex")
}

// Count implements the corresponding interface method.
func (v *MemDBVortex) Count(mctx manipulate.Context, identity elemental.Identity) (int, error) {

	v.RLock()
	defer v.RUnlock()

	if v.memory == nil {
		return 0, fmt.Errorf("datastore is not initialized")
	}
	return v.memory.Count(mctx, identity)
}

// Start connects to the websocket and starts collecting events
// until the given context is canceled or any non communication error is
// received. The eventual error will be received in the Errors() channel.
// If not nil, the given filter will be applied right away.
func (v *MemDBVortex) Start(ctx context.Context, e *elemental.PushFilter) {

	v.UpdateFilter(e)
}

// Implementation of the manipulate subscriber interface.

// UpdateFilter updates the current filter.
func (v *MemDBVortex) UpdateFilter(e *elemental.PushFilter) {

	v.RLock()
	defer v.RUnlock()

	if v.s == nil {
		return
	}

	v.commitIdentityEvent = map[string]struct{}{}

	filter := elemental.NewPushFilter()
	for identity := range v.processors {
		v.commitIdentityEvent[identity] = struct{}{}
		filter.FilterIdentity(identity)
	}

	for callerIdentity := range e.Identities {

		cfg, ok := v.processors[callerIdentity]
		if ok {
			// If we are processing this event and there is a client
			// subscription, we will only commit if the corresponding
			// flag is set. Otherwise, the client will have to handle
			// the update, so we removed it from the list.
			if !cfg.CommitOnEvent {
				delete(v.commitIdentityEvent, callerIdentity)
			}
			continue
		}
		// If it is not one of the registered identites and the client
		// has subscribed anyway, we still register and forward it to the
		// client.
		filter.FilterIdentity(callerIdentity)
	}

	v.s.UpdateFilter(filter)
}

// Subscriber returns itself as a subscriber since it implements
// the interface.
func (v *MemDBVortex) Subscriber() manipulate.Subscriber {
	return v
}

// Events returns the events channel.
func (v *MemDBVortex) Events() chan *elemental.Event { return v.subscriberEventChannel }

// Errors returns the errors channel.
func (v *MemDBVortex) Errors() chan error { return v.subscriberErrorChannel }

// Status returns the status channel.
func (v *MemDBVortex) Status() chan manipulate.SubscriberStatus { return v.subscriberStatusChannel }

// coreCRUDOperation implements the basic operation for updates of the db. This is used
// by create, update, and delete.
func (v *MemDBVortex) coreCRUDOperation(operation elemental.Operation, mctx manipulate.Context, objects ...elemental.Identifiable) error {

	if !v.isCommonIdentity(objects...) {
		return fmt.Errorf("all objects in operation must be of the same identity")
	}

	if mctx == nil {
		mctx = manipulate.NewContext(context.Background())
	}

	// If the identity is not registered or the request has a parent
	// send upstream. We are not dealing with this locally.
	if !v.isProcessable(mctx, objects[0].Identity()) {
		return v.commitUpstream(mctx.Context(), operation, mctx, objects...)
	}

	reconcile, err := v.genericUpdater(operation, mctx, objects...)
	if err != nil {
		return err
	}
	if !reconcile {
		return nil
	}

	return v.commitLocal(operation, mctx, objects)
}

// isCommonIdentity will validate that all objects in the operation have the same identity.
// We do not support calls with different identities.
func (v *MemDBVortex) isCommonIdentity(objects ...elemental.Identifiable) bool {
	if len(objects) == 0 {
		return false
	}

	first := objects[0].Identity()
	for _, obj := range objects {
		if !first.IsEqual(obj.Identity()) {
			return false
		}
	}

	return true
}

// isProcessable returns true if the request can be processed by the cache. If false,
// it must be forwarded to the upstream.
func (v *MemDBVortex) isProcessable(mctx manipulate.Context, identity elemental.Identity) bool {

	_, ok := v.processors[identity.Name]
	if !ok {
		return false
	}

	return mctx == nil || (mctx != nil && mctx.Parent() == nil)
}

// commitUpstream will commit a transaction to the upstream if it is not nil. It will
// return the upstream error.
func (v *MemDBVortex) commitUpstream(ctx context.Context, method elemental.Operation, mctx manipulate.Context, objects ...elemental.Identifiable) error {

	if v.m == nil {
		return nil
	}

	// If it is managed object we apply the pre-hook.
	cfg, ok := v.processors[objects[0].Identity().Name]
	if ok {
		reconcile, err := v.processHook(method, cfg.UpstreamHook, mctx, objects...)
		if !reconcile {
			return err
		}
	}

	// We always commit if prehook says ok or it is not a managed object.
	if err := manipulate.Retry(
		ctx,
		func() error {
			return v.methodFromType(method)(mctx, objects...)
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
func (v *MemDBVortex) commitLocal(method elemental.Operation, mctx manipulate.Context, objects []elemental.Identifiable) error {

	if objects == nil || len(objects) == 0 {
		return nil
	}

	cfg, ok := v.processors[objects[0].Identity().Name]
	if !ok {
		return nil
	}

	reconcile, err := v.processHook(method, cfg.LocalHook, mctx, objects...)
	if !reconcile {
		return err
	}

	if err := v.localMethodFromType(method)(mctx, objects...); err != nil {
		return err
	}

	if v.enableLog {
		v.logChannel <- &Transaction{
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
func (v *MemDBVortex) localMethodFromType(method elemental.Operation) updater {

	switch method {

	case elemental.OperationCreate:
		return v.memory.Create

	case elemental.OperationUpdate:
		return v.memory.Update

	default:
		return v.memory.Delete
	}
}

// methodFromType it will return an upstream function pointer based on the method.
func (v *MemDBVortex) methodFromType(method elemental.Operation) updater {

	switch method {

	case elemental.OperationCreate:
		return v.m.Create

	case elemental.OperationUpdate:
		return v.m.Update

	default:
		return v.m.Delete
	}
}

func (v *MemDBVortex) processHook(method elemental.Operation, hook config.Hook, mctx manipulate.Context, objects ...elemental.Identifiable) (reconcile bool, err error) {

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
func (v *MemDBVortex) genericUpdater(method elemental.Operation, mctx manipulate.Context, objects ...elemental.Identifiable) (bool, error) {

	if v.m == nil {
		return true, nil
	}

	// We are guaranteed that there is at least one object and the identity is processable.
	cfg := v.processors[objects[0].Identity().Name]

	switch cfg.Mode {

	case config.WriteThrough:
		// In WriteThrough mode make sure that the backend gets the create.
		// Only then store in the cache.
		return true, v.commitUpstream(mctx.Context(), method, mctx, objects...)

	case config.WriteBack:

		select {

		case v.transactionQueue <- &Transaction{
			mctx:     mctx,
			Objects:  objects,
			Method:   method,
			Deadline: time.Now().Add(cfg.QueueingDuration),
		}:
			return false, nil

		default:
			return false, fmt.Errorf("commit queue is full: %d", len(v.transactionQueue))
		}

	default:
		return false, fmt.Errorf("unknown caching mode: %d", cfg.Mode)
	}
}

// migrateObject will read all the objects from the backend and store them in
// the internal database.
func (v *MemDBVortex) migrateObject(ctx context.Context, cfg *config.ProcessorConfiguration) error {

	// We use pagination. We might 1000s of objects that are
	// reading at this point.

	page := 1
	pageSize := 100

	for {
		objList := v.model.Identifiables(cfg.Identity)

		mctx := manipulate.NewContext(
			ctx,
			manipulate.ContextOptionPage(page, pageSize),
		)

		subctx, cancel := context.WithDeadline(ctx, time.Now().Add(10*time.Second))
		defer cancel()

		if err := manipulate.Retry(
			subctx,
			func() error {
				return v.m.RetrieveMany(mctx, objList)
			},
			nil,
		); err != nil {
			return fmt.Errorf("unable to retrieve objects from backend: %s", err)
		}

		objects := objList.List()

		if len(objects) == 0 {
			return nil
		}

		if err := v.commitLocal(elemental.OperationCreate, nil, objList.List()); err != nil {
			return fmt.Errorf("unable to write objects to local db: %s", err)
		}

		page = page + 1

	}
}

// backgroundSync will empty the transaction queue and try to sync it
// with the backend.
func (v *MemDBVortex) backgroundSync(ctx context.Context) {

	if v.m == nil {
		return
	}

	for {
		select {
		case t, ok := <-v.transactionQueue:

			// If the channel is closed, then exit.
			if !ok {
				return
			}

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
			v.RLock()

			if _, ok := v.processors[t.Objects[0].Identity().Name]; !ok {
				v.RUnlock()
				continue
			}

			retryCtx, cancel := context.WithDeadline(ctx, t.Deadline)
			cancel()

			if err := v.commitUpstream(retryCtx, t.Method, t.mctx, t.Objects...); err != nil {
				v.RUnlock()
				zap.L().Error("failed to commit object upstream", zap.Error(err))
				continue
			}

			// Update the local copy of the object now.
			if err := v.commitLocal(t.Method, t.mctx, t.Objects); err != nil {
				zap.L().Error("failed to delete local object after failed resync", zap.Error(err))
			}

			v.RUnlock()

		case <-ctx.Done():

			// TODO: If we get killed with objects in the queue, then what ?
			// Do we ignore it and try to empty all objects or what ????
			return
		}
	}
}

// monitor registers for events for all the identities of interest
// and keeps the local cache up-to-date with the backend.
func (v *MemDBVortex) monitor(ctx context.Context) {
	if v.s == nil {
		return
	}

	filter := elemental.NewPushFilter()

	v.commitIdentityEvent = map[string]struct{}{}

	for identity, cfg := range v.processors {
		if cfg.CommitOnEvent {
			v.commitIdentityEvent[identity] = struct{}{}
		}
		filter.FilterIdentity(identity)
	}

	subctx, cancel := context.WithCancel(ctx)
	defer cancel()

	v.s.Start(subctx, filter)

	for {

		select {

		case evt := <-v.s.Events():

			v.RLock()
			_, commit := v.commitIdentityEvent[evt.Identity]
			v.RUnlock()

			if commit {
				v.eventHandler(ctx, evt)
			}

			// Push event upstream.
			select {

			case v.subscriberEventChannel <- evt:

			default:
				zap.L().Warn("Failed to push event upstream - subscriber busy")
			}

		case err := <-v.s.Errors():
			zap.L().Error("Received error from the push channel", zap.Error(err))
			// Push event upstream.
			select {

			case v.subscriberErrorChannel <- err:

			default:
				zap.L().Warn("Failed to push error upstream - subscriber busy")
			}

		case status := <-v.s.Status():

			switch status {

			case manipulate.SubscriberStatusDisconnection:
				zap.L().Warn("Upstream event channel interrupted. Reconnecting...")

			case manipulate.SubscriberStatusInitialConnection:
				zap.L().Info("Upstream event channel connected")

			case manipulate.SubscriberStatusReconnection:
				zap.L().Info("Upstream event channel restored")
				v.reconnectionHandler(subctx)

			case manipulate.SubscriberStatusFinalDisconnection:
				return
			}

			// Push event upstream.
			select {

			case v.subscriberStatusChannel <- status:

			default:
				zap.L().Warn("Failed to push status upstream - subscriber busy")
			}

		case <-ctx.Done():
			return
		}
	}
}

func (v *MemDBVortex) eventHandler(ctx context.Context, evt *elemental.Event) {

	if v.m == nil {
		return
	}

	obj := v.model.IdentifiableFromString(evt.Identity)

	if err := evt.Decode(obj); err != nil {
		zap.L().Error("Unable to unmarshal received event", zap.Error(err))
		return
	}

	v.RLock()
	defer v.RUnlock()

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

	list := elemental.IdentifiablesList{obj}
	if err := v.commitLocal(method, nil, list); err != nil {
		if method != elemental.OperationDelete {
			zap.L().Error("failed to commit locally an event notification", zap.String("event", evt.String()), zap.Error(err))
		}
	}
}

// reconnectionHandler will kick a re-sync when the push channel is
// restored. This might be heavy, but unclear if we have better
// mechanisms to react on a bad push channel.
func (v *MemDBVortex) reconnectionHandler(ctx context.Context) {
	if err := v.ReSync(ctx); err != nil {
		zap.L().Error("Failed to resync DB", zap.Error(err))
	}
}
