// Author: Alexandre Wilhelm
// See LICENSE file for full LICENSE
// Copyright 2016 Aporeto.

package manipcassandra

import (
	"fmt"
	"sync"
	"time"

	"github.com/aporeto-inc/elemental"
	"github.com/aporeto-inc/gocql"
	"github.com/aporeto-inc/manipulate"
	"github.com/aporeto-inc/manipulate/manipcassandra/encoding"

	log "github.com/Sirupsen/logrus"
)

// GocqlTimeout changes the timeout that will be used for the next gocql session
var GocqlTimeout = 600 * time.Millisecond

// ExtendedTimeout is used  for creating/dropping tables since gocql gives timeouts
var ExtendedTimeout = 2000 * time.Millisecond

// BatchRegistry is used to store the batches used in the store
type batchRegistry map[manipulate.TransactionID]*gocql.Batch

// AttributeUpdater structu used to make request as UPDATE policy SET NAME = NAME - ?
type attributeUpdater struct {
	Key             string
	Values          interface{}
	AssignationType elemental.AssignationType
}

// CassandraStore needs doc
type cassandraManipulator struct {
	Servers      []string
	KeySpace     string
	ProtoVersion int

	nativeSession     *gocql.Session
	batchRegistry     batchRegistry
	batchRegistryLock *sync.Mutex
}

// NewCassandraManipulator returns a new TransactionalManipulator backed by Cassandra.
func NewCassandraManipulator(servers []string, keyspace string, version int) manipulate.TransactionalManipulator {

	session, err := createNativeSession(servers, keyspace, version, GocqlTimeout)
	if err != nil {
		log.WithFields(log.Fields{
			"package":  "manipcassandra",
			"servers":  servers,
			"keyspace": keyspace,
			"version":  version,
			"error":    err.Error(),
		}).Fatal("Cannot connect to cassandra.")
	}

	return &cassandraManipulator{
		Servers:           servers,
		KeySpace:          keyspace,
		ProtoVersion:      version,
		batchRegistry:     batchRegistry{},
		batchRegistryLock: &sync.Mutex{},
		nativeSession:     session,
	}
}

func (c *cassandraManipulator) Retrieve(context *manipulate.Context, objects ...manipulate.Manipulable) error {

	if context == nil {
		context = manipulate.NewContext()
	}

	for _, object := range objects {

		primaryKeys, primaryValues, err := cassandra.PrimaryFieldsAndValues(object)

		if err != nil {
			log.WithFields(log.Fields{
				"package": "manipcassandra",
				"context": context,
				"error":   err,
			}).Error("Unable to extract primary keys and values.")
			return manipulate.NewError(err.Error(), manipulate.ErrCannotExractPrimaryFieldsAndValues)
		}

		command, values := buildGetCommand(context, object.Identity().Name, primaryKeys, primaryValues)

		log.WithFields(log.Fields{
			"package": "manipcassandra",
			"command": command,
			"values":  values,
			"context": context,
		}).Debug("sending select command to cassandra")

		if errs := unmarshalManipulable(c.nativeSession.Query(command, values...).Iter(), object); errs != nil {
			return errs
		}
	}

	return nil
}

func (c *cassandraManipulator) Delete(context *manipulate.Context, objects ...manipulate.Manipulable) error {

	if context == nil {
		context = manipulate.NewContext()
	}

	transactionID := context.TransactionID
	batch := c.batchForID(transactionID)

	for _, object := range objects {

		primaryKeys, primaryValues, err := cassandra.PrimaryFieldsAndValues(object)

		if err != nil {

			log.WithFields(log.Fields{
				"package": "manipcassandra",
				"context": context,
				"error":   err,
			}).Error("Unable to extract primary keys and values.")

			return manipulate.NewError(err.Error(), manipulate.ErrCannotExractPrimaryFieldsAndValues)
		}

		command, values := buildDeleteCommand(context, object.Identity().Name, primaryKeys, primaryValues)
		batch.Query(command, values...)
	}

	if transactionID != "" {
		return nil
	}

	if err := c.commitBatch(batch); err != nil {
		return manipulate.NewError(err.Error(), manipulate.ErrCannotExecuteBatch)
	}

	return nil
}

func (c *cassandraManipulator) RetrieveMany(context *manipulate.Context, identity elemental.Identity, dest interface{}) error {

	if context == nil {
		context = manipulate.NewContext()
	}

	command, values := buildGetCommand(context, identity.Name, []string{}, []interface{}{})

	log.WithFields(log.Fields{
		"package": "manipcassandra",
		"command": command,
		"values":  values,
		"context": context,
	}).Debug("sending select all command to cassandra")

	iter := c.nativeSession.Query(command, values...).Iter()

	return unmarshalManipulables(iter, dest)
}

func (c *cassandraManipulator) Create(context *manipulate.Context, objects ...manipulate.Manipulable) error {

	if context == nil {
		context = manipulate.NewContext()
	}

	transactionID := context.TransactionID
	batch := c.batchForID(transactionID)

	for _, object := range objects {
		object.SetIdentifier(gocql.TimeUUID().String())
		list, values, err := cassandra.FieldsAndValues(object)

		if err != nil {

			for _, object := range objects {
				object.SetIdentifier("")
			}

			log.WithFields(log.Fields{
				"package": "manipcassandra",
				"context": context,
				"error":   err,
			}).Error("Unable to extract fields and values.")

			return manipulate.NewError(err.Error(), manipulate.ErrCannotExtractFieldsAndValues)
		}

		command, values := buildInsertCommand(context, object.Identity().Name, list, values)
		batch.Query(command, values...)
	}

	if transactionID != "" {
		return nil
	}

	log.WithFields(log.Fields{
		"package": "manipcassandra",
		"batch":   batch.Entries,
		"context": context,
	}).Debug("sending create command to cassandra")

	if err := c.nativeSession.ExecuteBatch(batch); err != nil {

		log.WithFields(log.Fields{
			"package": "manipcassandra",
			"batch":   batch.Entries,
			"context": context,
			"error":   err,
		}).Error("Unabled to send create command to cassandra.")

		for _, object := range objects {
			object.SetIdentifier("")
		}

		return manipulate.NewError(err.Error(), manipulate.ErrCannotExecuteBatch)
	}

	log.WithFields(log.Fields{
		"package": "manipcassandra",
		"batch":   batch.Entries,
		"context": context,
	}).Debug("create command to cassandra sent")

	return nil
}

func (c *cassandraManipulator) Update(context *manipulate.Context, objects ...manipulate.Manipulable) error {

	if context == nil {
		context = manipulate.NewContext()
	}

	transactionID := context.TransactionID
	batch := c.batchForID(transactionID)

	for _, object := range objects {

		primaryKeys, primaryValues, err := cassandra.PrimaryFieldsAndValues(object)

		if err != nil {

			log.WithFields(log.Fields{
				"package": "manipcassandra",
				"context": context,
				"error":   err,
			}).Error("Unable to extract primary fields and values.")

			return manipulate.NewError(err.Error(), manipulate.ErrCannotExractPrimaryFieldsAndValues)
		}

		list, values, err := cassandra.FieldsAndValues(object)

		if err != nil {

			log.WithFields(log.Fields{
				"package": "manipcassandra",
				"context": context,
				"error":   err,
			}).Debug("Unable to extract fields and values.")

			return manipulate.NewError(err.Error(), manipulate.ErrCannotExtractFieldsAndValues)
		}

		command, values := buildUpdateCommand(context, object.Identity().Name, list, values, primaryKeys, primaryValues)

		batch.Query(command, values...)
	}

	if transactionID != "" {
		return nil
	}

	if err := c.commitBatch(batch); err != nil {
		return manipulate.NewError(err.Error(), manipulate.ErrCannotExecuteBatch)
	}

	return nil
}

func (c *cassandraManipulator) Count(context *manipulate.Context, identity elemental.Identity) (int, error) {

	if context == nil {
		context = manipulate.NewContext()
	}

	command, values := buildCountCommand(context, identity.Name)

	log.WithFields(log.Fields{
		"package": "manipcassandra",
		"query":   command,
		"context": context,
	}).Debug("sending count command to cassandra")

	iter := c.nativeSession.Query(command, values...).Iter()

	var count int
	success := iter.Scan(&count)

	if !success {
		return -1, manipulate.NewError("Unable to scan collection", manipulate.ErrCannotScan)
	}

	if err := iter.Close(); err != nil {
		return -1, manipulate.NewError(err.Error(), manipulate.ErrCannotCloseIterator)
	}

	log.WithFields(log.Fields{
		"package": "manipcassandra",
		"query":   command,
		"context": context,
	}).Debug("count command to cassandra sent")

	return count, nil
}

func (c *cassandraManipulator) Assign(*manipulate.Context, *elemental.Assignation) error {
	panic("Not implemented")
}

func (c *cassandraManipulator) Increment(context *manipulate.Context, identity elemental.Identity, counter string, inc int) error {

	if context == nil {
		context = manipulate.NewContext()
	}

	var transactionID manipulate.TransactionID
	var batch *gocql.Batch

	transactionID = context.TransactionID

	if transactionID != "" || batch == nil {
		batch = c.batchForID(transactionID)
	}

	var primaryKeys manipulate.FilterKey
	var primaryValues manipulate.FilterValue

	if filter := context.Filter; filter != nil {
		primaryKeys = filter.Keys()[0]
		primaryValues = filter.Values()[0]
	}

	command, values := buildIncrementCommand(context, identity.Name, counter, inc, primaryKeys, primaryValues)

	batch.Query(command, values...)

	if transactionID != "" {
		return nil
	}

	if err := c.commitBatch(batch); err != nil {
		return manipulate.NewError(err.Error(), manipulate.ErrCannotExecuteBatch)
	}

	return nil
}

func (c *cassandraManipulator) Commit(id manipulate.TransactionID) error {

	defer func() { c.unregisterBatch(id) }()

	if c.registeredBatchWithID(id) == nil {
		log.WithFields(log.Fields{
			"package":       "manipcassandra",
			"store":         c,
			"transactionID": id,
		}).Error("No batch found for the given transaction.")

		return manipulate.NewError("No batch found for the given transaction.", manipulate.ErrCannotCommit)
		// return nil
	}

	if err := c.commitBatch(c.registeredBatchWithID(id)); err != nil {
		return manipulate.NewError(err.Error(), manipulate.ErrCannotExecuteBatch)
	}

	return nil
}

func (c *cassandraManipulator) Abort(id manipulate.TransactionID) bool {

	if c.registeredBatchWithID(id) == nil {
		return false
	}

	c.unregisterBatch(id)

	return true
}

// UpdateCollection seems to be useless.
func (c *cassandraManipulator) UpdateCollection(context *manipulate.Context, attributeUpdate *attributeUpdater, object manipulate.Manipulable) error {

	if context == nil {
		context = manipulate.NewContext()
	}

	primaryKeys, primaryValues, err := cassandra.PrimaryFieldsAndValues(object)

	if err != nil {

		log.WithFields(log.Fields{
			"package": "manipcassandra",
			"context": context,
			"error":   err,
		}).Error("Unable to extract primary fields and values.")

		return manipulate.NewError(err.Error(), manipulate.ErrCannotExractPrimaryFieldsAndValues)
	}

	command, values := buildUpdateCollectionCommand(context, object.Identity().Name, attributeUpdate, primaryKeys, primaryValues)
	if err := c.nativeSession.Query(command, values...).Exec(); err != nil {

		log.WithFields(log.Fields{
			"package": "manipcassandra",
			"context": context,
			"error":   err,
		}).Error("Unable to send update collection command to cassandra.")

		return manipulate.NewError(err.Error(), manipulate.ErrCannotExecuteQuery)
	}

	return nil
}

func (c *cassandraManipulator) batchForID(id manipulate.TransactionID) *gocql.Batch {

	if id == "" {
		return c.nativeSession.NewBatch(gocql.UnloggedBatch)
	}

	batch := c.registeredBatchWithID(id)

	if batch == nil {
		batch = c.nativeSession.NewBatch(gocql.UnloggedBatch)
		c.registerBatch(id, batch)
	}

	return batch
}

func (c *cassandraManipulator) commitBatch(b *gocql.Batch) error {

	log.WithFields(log.Fields{
		"package": "manipcassandra",
		"batch":   b.Entries,
	}).Debug("Commiting batch to cassandra.")

	if err := c.nativeSession.ExecuteBatch(b); err != nil {

		log.WithFields(log.Fields{
			"package": "manipcassandra",
			"batch":   b.Entries,
			"error":   err,
		}).Debug("Unable to send batch command.")

		return err
	}

	return nil
}

func (c *cassandraManipulator) registerBatch(id manipulate.TransactionID, batch *gocql.Batch) {

	c.batchRegistryLock.Lock()
	c.batchRegistry[id] = batch
	c.batchRegistryLock.Unlock()
}

func (c *cassandraManipulator) unregisterBatch(id manipulate.TransactionID) {

	c.batchRegistryLock.Lock()
	delete(c.batchRegistry, id)
	c.batchRegistryLock.Unlock()
}

func (c *cassandraManipulator) registeredBatchWithID(id manipulate.TransactionID) *gocql.Batch {

	c.batchRegistryLock.Lock()
	b := c.batchRegistry[id]
	c.batchRegistryLock.Unlock()

	return b
}

func unmarshalManipulable(iter *gocql.Iter, object manipulate.Manipulable) error {

	maps, errs := sliceMaps(iter)
	if errs != nil {
		return errs
	}

	if len(maps) != 1 {
		return manipulate.NewError(fmt.Sprintf("Could not find object %s.", object), manipulate.ErrObjectNotFound)
	}

	if err := cassandra.Unmarshal(maps[0], object); err != nil {
		return manipulate.NewError(err.Error(), manipulate.ErrCannotUnmarshal)
	}

	return nil
}

func unmarshalManipulables(iter *gocql.Iter, objects interface{}) error {

	if iter.NumRows() == 0 {
		return nil
	}

	maps, errs := sliceMaps(iter)
	if errs != nil {
		return errs
	}

	if len(maps) < 1 {
		return nil
	}

	if err := cassandra.Unmarshal(maps, objects); err != nil {
		return manipulate.NewError(err.Error(), manipulate.ErrCannotUnmarshal)
	}

	return nil
}

func sliceMaps(iter *gocql.Iter) ([]map[string]interface{}, error) {

	maps, err := iter.SliceMap()
	if err != nil {
		return nil, manipulate.NewError(err.Error(), manipulate.ErrCannotSlice)
	}

	if err = iter.Close(); err != nil {
		return nil, manipulate.NewError(err.Error(), manipulate.ErrCannotCloseIterator)
	}

	return maps, nil
}

func createNativeSession(servers []string, keyspace string, version int, timeout time.Duration) (*gocql.Session, error) {

	cluster := gocql.NewCluster(servers...)
	cluster.Keyspace = keyspace
	cluster.Consistency = gocql.Quorum
	cluster.ProtoVersion = version
	cluster.Timeout = timeout
	cluster.RetryPolicy = &gocql.SimpleRetryPolicy{NumRetries: 5}

	return cluster.CreateSession()
}
