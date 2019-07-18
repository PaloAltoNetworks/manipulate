// Copyright 2019 Aporeto Inc.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//     http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
	"golang.org/x/time/rate"
)

// updater is the type of all crud functions.
type updater func(mctx manipulate.Context, object elemental.Identifiable) error

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
	upstreamReconciler      Reconciler
	downstreamReconciler    Reconciler

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
		upstreamReconciler:      cfg.upstreamReconciler,
		downstreamReconciler:    cfg.downstreamReconciler,
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
	go m.backgroundSync(ctx, cfg.rateLimiter)

	return m, nil
}

// Flush implements the flush interface of the Vortex. It will flush
// all the cache for write-through. For write-back it will wait
// for a maximum of 10 seconds for transactions to complete. When
// done it will flush the channel and create a completely fresh
// db.
func (m *vortexManipulator) Flush(ctx context.Context) error {

	m.RLock()
	defer m.RUnlock()

	if m.prefetcher != nil {
		m.prefetcher.Flush()
	}

	f, ok := m.downstreamManipulator.(manipulate.FlushableManipulator)
	if ok {
		// Wait for the channel to clean up
		maxDelay := time.Now().Add(10 * time.Second)
		for len(m.transactionQueue) > 0 && time.Now().Before(maxDelay) {
			time.Sleep(1 * time.Second)
		}

		// Flush any outstanding transactions and restart the backgrond sync
		if err := f.Flush(ctx); err != nil {
			return fmt.Errorf("unable to flush the datastore: %s", err)
		}
	}

	if m.prefetcher != nil {
		if err := m.warmUp(ctx); err != nil {
			return err
		}
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

	if cfg := m.processors[dest.Identity().Name]; cfg != nil && cfg.RetrieveManyHook != nil {
		commit, err := cfg.RetrieveManyHook(m.downstreamManipulator, mctx, dest)
		if !commit {
			return err
		}
	}

	return m.downstreamManipulator.RetrieveMany(mctx, dest)
}

func (m *vortexManipulator) Retrieve(mctx manipulate.Context, object elemental.Identifiable) error {

	m.RLock()
	defer m.RUnlock()

	if mctx == nil {
		mctx = manipulate.NewContext(context.Background())
	}

	if m.prefetcher != nil {
		prefetched, err := m.prefetcher.Prefetch(mctx.Context(), elemental.OperationRetrieve, object.Identity(), m.upstreamManipulator, mctx.Derive())
		if err != nil {
			return fmt.Errorf("unable to prefetch data for retrieve operation for '%s': %s", object.Identity(), err)
		}
		if err := m.insertPrefetchedData(prefetched); err != nil {
			return fmt.Errorf("unable to insert prefetched data for retrieve operation for '%s': %s", object.Identity(), err)
		}
	}

	// If we are not processing the object, we send it upstream.
	// We only deal with CRUDs.
	if !m.shouldProcess(mctx, object.Identity()) {
		if m.upstreamManipulator != nil {
			return m.upstreamManipulator.Retrieve(mctx, object)
		}
		return nil
	}

	if err := m.downstreamManipulator.Retrieve(mctx, object); err != nil {

		// If we can't find it locally, and its strong consistency retrieve
		// we will try the backend if we have one.
		if m.upstreamManipulator == nil || !isStrongReadConsistency(mctx, m.processors[object.Identity().Name], m.defaultReadConsistency) {
			return err
		}

		if err := m.upstreamManipulator.Retrieve(mctx, object); err != nil {
			return err
		}

		// Make sure that we update our cache for future reference.
		if err := m.downstreamManipulator.Create(mctx, object); err != nil {
			return fmt.Errorf("unable to update local cache from backend: %s", err)
		}
	}

	return nil
}

func (m *vortexManipulator) Create(mctx manipulate.Context, object elemental.Identifiable) error {

	m.RLock()
	defer m.RUnlock()

	return m.coreCRUDOperation(elemental.OperationCreate, mctx, object)
}

func (m *vortexManipulator) Update(mctx manipulate.Context, object elemental.Identifiable) error {

	m.RLock()
	defer m.RUnlock()

	return m.coreCRUDOperation(elemental.OperationUpdate, mctx, object)
}

func (m *vortexManipulator) Delete(mctx manipulate.Context, object elemental.Identifiable) error {

	m.RLock()
	defer m.RUnlock()

	return m.coreCRUDOperation(elemental.OperationDelete, mctx, object)
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
func (m *vortexManipulator) coreCRUDOperation(operation elemental.Operation, mctx manipulate.Context, object elemental.Identifiable) error {

	if mctx == nil {
		mctx = manipulate.NewContext(context.Background())
	}

	// If the identity is not registered or the request has a parent
	// send upstream. We are not dealing with this locally.
	if !m.shouldProcess(mctx, object.Identity()) {
		return m.commitUpstream(mctx.Context(), operation, mctx, object)
	}

	reconcile, err := m.genericUpdater(operation, mctx, object)
	if err != nil {
		return err
	}
	if !reconcile {
		return nil
	}

	return m.commitLocal(operation, mctx, object)
}

// shouldProcess returns true if the request can be processed by the cache. If false,
// it must be forwarded to the upstream.
func (m *vortexManipulator) shouldProcess(mctx manipulate.Context, identity elemental.Identity) bool {

	_, ok := m.processors[identity.Name]
	if !ok {
		return false
	}

	return mctx == nil || mctx.Parent() == nil
}

// commitUpstream will commit a transaction to the upstream if it is not nil. It will
// return the upstream error.
func (m *vortexManipulator) commitUpstream(ctx context.Context, operation elemental.Operation, mctx manipulate.Context, object elemental.Identifiable) error {

	var reconcile bool
	var err error

	if m.upstreamManipulator == nil {
		return nil
	}

	// If we have an accepter, we see if it accepts the write
	if m.upstreamReconciler != nil {
		object, reconcile, err = m.upstreamReconciler.Reconcile(mctx, operation, object)
		if err != nil {
			return err
		}
		if !reconcile {
			return nil
		}
	}

	// If it is managed object we apply the pre-hook.
	cfg, ok := m.processors[object.Identity().Name]
	if ok && cfg.UpstreamReconciler != nil {
		object, reconcile, err = cfg.UpstreamReconciler.Reconcile(mctx, operation, object)
		if err != nil {
			return err
		}
		if !reconcile {
			return nil
		}
	}

	// We always commit if prehook says ok or it is not a managed object.
	if err := m.methodFromType(operation)(mctx, object); err != nil {
		return err
	}

	return nil
}

// commitLocal will commit a transaction locally after processing any
// hooks. It will return error if either the hook or the local commit
// fail for some reason.
func (m *vortexManipulator) commitLocal(operation elemental.Operation, mctx manipulate.Context, object elemental.Identifiable) error {

	var reconcile bool
	var err error

	if mctx == nil {
		mctx = manipulate.NewContext(context.Background())
	}

	// If we have a global Reconciler, we see if it accepts the write.
	if m.downstreamReconciler != nil {
		object, reconcile, err = m.downstreamReconciler.Reconcile(mctx, operation, object)
		if err != nil {
			return err
		}
		if !reconcile {
			return nil
		}
	}

	cfg, ok := m.processors[object.Identity().Name]
	if !ok {
		return nil
	}

	// If we have a processor Reconciler, we see if it accepts the write.
	if cfg.DownstreamReconciler != nil {
		object, reconcile, err = cfg.DownstreamReconciler.Reconcile(mctx, operation, object)
		if err != nil {
			return err
		}
		if !reconcile {
			return nil
		}
	}

	if err := m.localMethodFromType(operation)(mctx, object); err != nil {
		return err
	}

	if m.enableLog {
		m.logChannel <- &Transaction{
			Date:   time.Now(),
			mctx:   mctx,
			Object: object,
			Method: operation,
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

// genericUpdate will implement the updates. It takes as parameters the methods
// to be used (update, create, delete) and avoids repeating code. It will
// return true if the transaction has to be committed in the local DB. It will
// return an error if the backend fails. Specifically:
// For WriteThrough: it will return an error if the backend fails.
// For WriteBack it will cache it and return commit=false. The commit will happen
// later after the object is stored in the backend.
func (m *vortexManipulator) genericUpdater(method elemental.Operation, mctx manipulate.Context, object elemental.Identifiable) (bool, error) {

	if m.upstreamManipulator == nil {
		return true, nil
	}

	// We are guaranteed that there is at least one object and the identity is processable.
	processor := m.processors[object.Identity().Name]

	// In Strong consistency we make sure that the backend gets the create.
	// Only then store in the cache.

	if isStrongWriteConsistency(mctx, processor, m.defaultWriteConsistency) {
		return true, m.commitUpstream(mctx.Context(), method, mctx, object)
	}

	tdeadline := processor.QueueingDuration
	if tdeadline == 0 {
		tdeadline = m.defaultQueueDuration
	}

	select {

	case m.transactionQueue <- &Transaction{
		mctx:     mctx,
		Object:   object,
		Method:   method,
		Deadline: time.Now().Add(tdeadline),
	}:
		return false, nil

	default:
		return false, fmt.Errorf("commit queue is full: %d", len(m.transactionQueue))
	}
}

// backgroundSync will empty the transaction queue and try to sync it
// with the backend.
func (m *vortexManipulator) backgroundSync(ctx context.Context, limiter *rate.Limiter) {

	if m.upstreamManipulator == nil {
		return
	}

	for {
		select {
		case t := <-m.transactionQueue:
			// Rate limit.
			if err := limiter.Wait(ctx); err != nil {
				zap.L().Warn("unable to rate limit", zap.Error(err))
			}

			// If the dealine is exceeded we just drop the request
			// no matter what. This allows us to clean up the queue
			// if there is a problem.
			if time.Now().After(t.Deadline) {
				continue
			}

			if t.Object == nil {
				continue
			}

			// We first try to update the backend. If this succeeds
			// then we also update the local db. At this point
			// the object can be accessible through our API since
			// the ID has been populated.
			m.RLock()

			if _, ok := m.processors[t.Object.Identity().Name]; !ok {
				m.RUnlock()
				continue
			}

			retryCtx, cancel := context.WithDeadline(ctx, t.Deadline)
			defer cancel()

			if err := m.commitUpstream(retryCtx, t.Method, t.mctx, t.Object); err != nil {
				m.RUnlock()
				zap.L().Error("failed to commit object upstream", zap.Error(err))
				continue
			}

			// Update the local copy of the object now.
			if err := m.commitLocal(t.Method, t.mctx, t.Object); err != nil {
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
				if err := m.eventHandler(ctx, evt); err != nil {
					m.pushErrors(fmt.Errorf("unable to handle event: %s", err))
					continue
				}
			}

			m.pushEvent(evt)

		case err := <-m.upstreamSubscriber.Errors():
			m.pushErrors(fmt.Errorf("upstream error: %s", err))

		case status := <-m.upstreamSubscriber.Status():

			switch status {

			case manipulate.SubscriberStatusReconnection:

				// We resync everything
				if err := m.Flush(ctx); err != nil {
					m.pushErrors(fmt.Errorf("unable to flush: %s", err))
				}

			case manipulate.SubscriberStatusFinalDisconnection:
				m.pushStatus(status)
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

func (m *vortexManipulator) eventHandler(ctx context.Context, evt *elemental.Event) error {

	if m.upstreamManipulator == nil {
		return nil
	}

	obj := m.model.IdentifiableFromString(evt.Identity)

	if err := evt.Decode(obj); err != nil {
		return fmt.Errorf("unable to unmarshal received event: %s", err)
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
		return fmt.Errorf("unsupported event received: %s", evt.Type)
	}

	if err := m.commitLocal(method, nil, obj); err != nil {
		if method != elemental.OperationDelete {
			return fmt.Errorf("unable to commit event of type '%s': %s", evt.Type, err)
		}
	}

	return nil
}

func (m *vortexManipulator) insertPrefetchedData(prefetched elemental.Identifiables) error {

	if prefetched == nil {
		return nil
	}

	lst := prefetched.List()
	if len(lst) == 0 {
		return nil
	}

	for _, item := range lst {
		if err := m.commitLocal(elemental.OperationCreate, nil, item); err != nil {
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

		if err := m.insertPrefetchedData(prefetched); err != nil {
			return fmt.Errorf("unable to insert prefetched '%s' for warm up operation: %s", proc.Identity, err)
		}
	}

	return nil
}
