// Author: Alexandre Wilhelm
// See LICENSE file for full LICENSE
// Copyright 2016 Aporeto.

package manipcassandra

import (
	"fmt"
	"strings"
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
type BatchRegistry map[manipulate.TransactionID]*gocql.Batch

// AttributeUpdater structu used to make request as UPDATE policy SET NAME = NAME - ?
type AttributeUpdater struct {
	Key             string
	Values          interface{}
	AssignationType elemental.AssignationType
}

// CassandraStore needs doc
type CassandraStore struct {
	Servers      []string
	KeySpace     string
	ProtoVersion int

	nativeSession     *gocql.Session
	batchRegistry     BatchRegistry
	batchRegistryLock *sync.Mutex
}

// NewCassandraStore returns a new *CassandraStore
// You can specify the parameters servers, jeyspace and version.
func NewCassandraStore(servers []string, keyspace string, version int) *CassandraStore {

	return &CassandraStore{
		Servers:           servers,
		KeySpace:          keyspace,
		ProtoVersion:      version,
		batchRegistry:     BatchRegistry{},
		batchRegistryLock: &sync.Mutex{},
	}
}

// Stop will close the cassandra session
func (c *CassandraStore) Stop() {
	c.nativeSession.Close()
	c.nativeSession = nil
}

// Start will start the cassandra session
func (c *CassandraStore) Start() error {

	session, err := c.createNativeSession(c.Servers, c.KeySpace, c.ProtoVersion, GocqlTimeout)

	if err != nil {
		return err
	}

	c.nativeSession = session

	return nil
}

// Retrieve will launch a set of select to the cassandra session
// These selects are not batched and won't be, this is not possible in cassandra
// The given Contexts is either an array of Context or a Context, each Context will map with the given objects
// The given objects will be automatically updated
// This method returns an array of Errors if something wrong happens
func (c *CassandraStore) Retrieve(context manipulate.Contexts, objects ...manipulate.Manipulable) elemental.Errors {

	for index, object := range objects {

		context := manipulate.ContextForIndex(context, index)
		primaryKeys, primaryValues, err := cassandra.PrimaryFieldsAndValues(object)

		if err != nil {
			log.WithFields(log.Fields{
				"package": "manipcassandra",
				"context": context,
				"error":   err,
			}).Error("Unable to extract primary keys and values.")
			return makeError(err.Error(), ErrCannotExractPrimaryFieldsAndValues)
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

// Delete will launch a set of delete to the cassandra session
// The delete commands will be in a bacth
// The given Contexts is either an array of Context or a Context, each Context will map with the given objects
// The given objects will be deleted in the database
// This method returns an array of Errors if something wrong happens
func (c *CassandraStore) Delete(context manipulate.Contexts, objects ...manipulate.Manipulable) elemental.Errors {

	var transactionID manipulate.TransactionID
	var batch *gocql.Batch

	for index, object := range objects {

		primaryKeys, primaryValues, err := cassandra.PrimaryFieldsAndValues(object)

		if err != nil {

			log.WithFields(log.Fields{
				"package": "manipcassandra",
				"context": context,
				"error":   err,
			}).Error("Unable to extract primary keys and values.")

			return makeError(err.Error(), ErrCannotExractPrimaryFieldsAndValues)
		}

		context := manipulate.ContextForIndex(context, index)
		transactionID = context.TransactionID

		if transactionID != "" || batch == nil {
			batch = c.batchForID(transactionID)
		}

		command, values := buildDeleteCommand(context, object.Identity().Name, primaryKeys, primaryValues)
		batch.Query(command, values...)
	}

	if transactionID != "" {
		return nil
	}

	if err := c.commitBatch(batch); err != nil {
		return makeError(err.Error(), ErrCannotExecuteBatch)
	}

	return nil
}

// RetrieveChildren will launch a set of delete to the cassandra session
// Onyl one command select will be launched to the database
// The given Contexts is either an array of Context or a Context, each Context will map with the given objects
// The retrieven objects will be stored in the given dest{}. For instance if your a fetching Tag Object, you should pass an array of Tag ([]*Tag)
// The param parent is ignored
// This method returns an array of Errors if something wrong happens
func (c *CassandraStore) RetrieveChildren(context manipulate.Contexts, parent manipulate.Manipulable, identity elemental.Identity, dest interface{}) elemental.Errors {

	ctx := manipulate.ContextForIndex(context, 0)
	command, values := buildGetCommand(ctx, identity.Name, []string{}, []interface{}{})

	log.WithFields(log.Fields{
		"package": "manipcassandra",
		"command": command,
		"values":  values,
		"context": context,
	}).Debug("sending select all command to cassandra")

	iter := c.nativeSession.Query(command, values...).Iter()

	if errs := unmarshalManipulables(iter, dest); errs != nil {
		return errs
	}

	return nil
}

// Create will launch a set of delete to the cassandra session
// The create commands will be in a bacth
// The given Contexts is either an array of Context or a Context, each Context will map with the given objects
// The given objects will be created in the database
// If the object has an field ID, this field will be automatically set to a new UUID generated by gocql.TimeUUID.String() and will be set by calling the method SetIdentifier on the object
// This method returns an array of Errors if something wrong happens
func (c *CassandraStore) Create(context manipulate.Contexts, parent manipulate.Manipulable, objects ...manipulate.Manipulable) elemental.Errors {

	var transactionID manipulate.TransactionID
	var batch *gocql.Batch

	for index, object := range objects {
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

			return makeError(err.Error(), ErrCannotExtractFieldsAndValues)
		}

		context := manipulate.ContextForIndex(context, index)
		transactionID = context.TransactionID

		if transactionID != "" || batch == nil {
			batch = c.batchForID(transactionID)
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

		return makeError(err.Error(), ErrCannotExecuteBatch)
	}

	log.WithFields(log.Fields{
		"package": "manipcassandra",
		"batch":   batch.Entries,
		"context": context,
	}).Debug("create command to cassandra sent")

	return nil
}

// UpdateCollection will launch a set of update to the cassandra session
// The update commands will be not in a batch
// The given Contexts is either an array of Context or a Context, each Context will map with the given objects
// The given objects will be updated in the database
// This method returns an array of Errors if something wrong happens
// This method is used for request as UPDATE policy SET NAME = NAME - ?
func (c *CassandraStore) UpdateCollection(context manipulate.Contexts, attributeUpdate *AttributeUpdater, object manipulate.Manipulable) elemental.Errors {

	primaryKeys, primaryValues, err := cassandra.PrimaryFieldsAndValues(object)

	if err != nil {

		log.WithFields(log.Fields{
			"package": "manipcassandra",
			"context": context,
			"error":   err,
		}).Error("Unable to extract primary fields and values.")

		return makeError(err.Error(), ErrCannotExractPrimaryFieldsAndValues)
	}

	command, values := buildUpdateCollectionCommand(manipulate.ContextForIndex(context, 0), object.Identity().Name, attributeUpdate, primaryKeys, primaryValues)
	if err := c.nativeSession.Query(command, values...).Exec(); err != nil {

		log.WithFields(log.Fields{
			"package": "manipcassandra",
			"context": context,
			"error":   err,
		}).Error("Unable to send update collection command to cassandra.")

		return makeError(err.Error(), ErrCannotExecuteQuery)
	}

	return nil
}

// Update will launch a set of update to the cassandra session
// The update commands will be in a batch
// The given Contexts is either an array of Context or a Context, each Context will map with the given objects
// The given objects will be updated in the database
// This method returns an array of Errors if something wrong happens
func (c *CassandraStore) Update(context manipulate.Contexts, objects ...manipulate.Manipulable) elemental.Errors {

	var transactionID manipulate.TransactionID
	var batch *gocql.Batch

	for index, object := range objects {

		primaryKeys, primaryValues, err := cassandra.PrimaryFieldsAndValues(object)

		if err != nil {

			log.WithFields(log.Fields{
				"package": "manipcassandra",
				"context": context,
				"error":   err,
			}).Error("Unable to extract primary fields and values.")

			return makeError(err.Error(), ErrCannotExractPrimaryFieldsAndValues)
		}

		list, values, err := cassandra.FieldsAndValues(object)

		if err != nil {

			log.WithFields(log.Fields{
				"package": "manipcassandra",
				"context": context,
				"error":   err,
			}).Debug("Unable to extract fields and values.")

			return makeError(err.Error(), ErrCannotExtractFieldsAndValues)
		}

		context := manipulate.ContextForIndex(context, index)
		transactionID = context.TransactionID

		if transactionID != "" || batch == nil {
			batch = c.batchForID(transactionID)
		}

		command, values := buildUpdateCommand(context, object.Identity().Name, list, values, primaryKeys, primaryValues)

		batch.Query(command, values...)
	}

	if transactionID != "" {
		return nil
	}

	if err := c.commitBatch(batch); err != nil {
		return makeError(err.Error(), ErrCannotExecuteBatch)
	}

	return nil
}

// Count will launch a count to the cassandra session
// The count commands will not be in a bacth
// The given Contexts has to be Context
// This method returns an array of Errors if something wrong happens
// This method return an interger, the count asked
func (c *CassandraStore) Count(context manipulate.Contexts, identity elemental.Identity) (int, elemental.Errors) {

	ctx := manipulate.ContextForIndex(context, 0)
	command, values := buildCountCommand(ctx, identity.Name)

	log.WithFields(log.Fields{
		"package": "manipcassandra",
		"query":   command,
		"context": context,
	}).Debug("sending count command to cassandra")

	iter := c.nativeSession.Query(command, values...).Iter()

	var count int
	success := iter.Scan(&count)

	if !success {
		return -1, makeError("Unable to scan collection", ErrCannotScan)
	}

	if err := iter.Close(); err != nil {
		return -1, makeError(err.Error(), ErrCannotCloseIterator)
	}

	log.WithFields(log.Fields{
		"package": "manipcassandra",
		"query":   command,
		"context": context,
	}).Debug("count command to cassandra sent")

	return count, nil
}

// Assign is not yet implemented
func (c *CassandraStore) Assign(contexts manipulate.Contexts, parent manipulate.Manipulable, assignation *elemental.Assignation) elemental.Errors {
	panic("Not implemented")
}

// Commit will execute the batch of the given transaction
// The method will return an error if the batch does not succeed
func (c *CassandraStore) Commit(id manipulate.TransactionID) elemental.Errors {

	defer func() { c.unregisterBatch(id) }()

	if c.registeredBatchWithID(id) == nil {
		log.WithFields(log.Fields{
			"package":       "manipcassandra",
			"store":         c,
			"transactionID": id,
		}).Error("No batch found for the given transaction.")

		return makeError("No batch found for the given transaction.", ErrCannotCommit)
		// return nil
	}

	if err := c.commitBatch(c.registeredBatchWithID(id)); err != nil {
		return makeError(err.Error(), ErrCannotExecuteBatch)
	}

	return nil
}

// Abort aborts the given transaction ID.
func (c *CassandraStore) Abort(id manipulate.TransactionID) bool {

	if c.registeredBatchWithID(id) == nil {
		return false
	}

	c.unregisterBatch(id)

	return true
}

// Increment increments the given counter name from the given counter table name by the given number.
// To Decrement, simply pass a negative value.
func (c *CassandraStore) Increment(contexts manipulate.Contexts, name, counter string, inc int, primaryKeys []string, primaryValues []interface{}) elemental.Errors {

	var transactionID manipulate.TransactionID
	var batch *gocql.Batch

	context := manipulate.ContextForIndex(contexts, 0)
	transactionID = context.TransactionID

	if transactionID != "" || batch == nil {
		batch = c.batchForID(transactionID)
	}

	command, values := buildIncrementCommand(context, name, counter, inc, primaryKeys, primaryValues)

	batch.Query(command, values...)

	if transactionID != "" {
		return nil
	}

	if err := c.commitBatch(batch); err != nil {
		return makeError(err.Error(), ErrCannotExecuteBatch)
	}

	return nil
}

// DoesKeyspaceExist checks if the configured keyspace exists
func (c *CassandraStore) DoesKeyspaceExist() (bool, error) {

	session, err := c.createNativeSession(c.Servers, "", c.ProtoVersion, GocqlTimeout)
	if err != nil {
		return false, err
	}
	defer session.Close()

	info, err := session.KeyspaceMetadata(c.KeySpace)

	if err != nil {
		log.WithFields(log.Fields{
			"package":  "manipcassandra",
			"keyspace": c.KeySpace,
			"error":    err,
		}).Error("unable to get keyspace metadata")

		return false, err
	}

	return len(info.Tables) > 0, nil
}

// CreateKeySpace creates a new keyspace
func (c *CassandraStore) CreateKeySpace(replicationFactor int) error {

	session, err := c.createNativeSession(c.Servers, "", c.ProtoVersion, ExtendedTimeout)
	if err != nil {
		return err
	}
	defer session.Close()

	query := session.Query(
		fmt.Sprintf("CREATE KEYSPACE %s WITH replication = {'class' : 'SimpleStrategy', 'replication_factor': %d}", c.KeySpace, replicationFactor))

	return query.Exec()
}

// DropKeySpace deletes the given keyspace
func (c *CassandraStore) DropKeySpace() error {

	session, err := c.createNativeSession(c.Servers, "", c.ProtoVersion, ExtendedTimeout)
	if err != nil {
		return err
	}
	defer session.Close()

	query := session.Query(fmt.Sprintf("DROP KEYSPACE IF EXISTS %s", c.KeySpace))

	return query.Exec()
}

// ExecuteScript opens a new session, runs the given script in a mode and close the session.
func (c *CassandraStore) ExecuteScript(data string) error {

	session, err := c.createNativeSession(c.Servers, c.KeySpace, c.ProtoVersion, ExtendedTimeout)

	if err != nil {
		return err
	}
	defer session.Close()

	for _, statement := range strings.Split(data, ";\n") {

		if len(statement) == 0 {
			continue
		}

		if err := session.Query(statement).Exec(); err != nil {
			log.WithFields(log.Fields{
				"package": "manipcassandra",
				"error":   err,
			}).Error("unable to execute query. aborting script in the middle. be sure to clean up my mess.")

			return err
		}
	}

	return nil
}

// createNativeSession returns a native gocql.Session.
func (c *CassandraStore) createNativeSession(srvs []string, ks string, v int, timeout time.Duration) (*gocql.Session, error) {
	cluster := gocql.NewCluster(srvs...)
	cluster.Keyspace = ks
	cluster.Consistency = gocql.Quorum
	cluster.ProtoVersion = v
	cluster.Timeout = timeout
	cluster.RetryPolicy = &gocql.SimpleRetryPolicy{NumRetries: 5}

	session, err := cluster.CreateSession()

	if err != nil {
		return nil, err
	}

	return session, nil
}

// batchForID return a gocql.Batch,
// The dev can then do whatever he wants with
// If id is emptyn it will return a new batch
// If id does not match with a batch, it will create a new batch
func (c *CassandraStore) batchForID(id manipulate.TransactionID) *gocql.Batch {

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

// CommitBatch commit the given batch
// The dev can then do whatever he wants with
func (c *CassandraStore) commitBatch(b *gocql.Batch) error {

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

func (c *CassandraStore) registerBatch(id manipulate.TransactionID, batch *gocql.Batch) {

	c.batchRegistryLock.Lock()
	c.batchRegistry[id] = batch
	c.batchRegistryLock.Unlock()
}

func (c *CassandraStore) unregisterBatch(id manipulate.TransactionID) {

	c.batchRegistryLock.Lock()
	delete(c.batchRegistry, id)
	c.batchRegistryLock.Unlock()
}

func (c *CassandraStore) registeredBatchWithID(id manipulate.TransactionID) *gocql.Batch {

	c.batchRegistryLock.Lock()
	b := c.batchRegistry[id]
	c.batchRegistryLock.Unlock()

	return b
}

// unmarshalManipulable this will be called by the method retrieve as we know in which object we will store the retrieved data
func unmarshalManipulable(iter *gocql.Iter, object manipulate.Manipulable) elemental.Errors {

	maps, errs := sliceMaps(iter)
	if errs != nil {
		return errs
	}

	if len(maps) != 1 {
		return makeError(fmt.Sprintf("Could not find object %s.", object), ErrObjectNotFound)
	}

	if err := cassandra.Unmarshal(maps[0], object); err != nil {
		return makeError(err.Error(), ErrCannotUnmarshal)
	}

	return nil
}

// unmarshalManipulables this will be called by the method retrieveChildren, we will store the objects in the empty given array/interface
func unmarshalManipulables(iter *gocql.Iter, objects interface{}) elemental.Errors {

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
		return makeError(err.Error(), ErrCannotUnmarshal)
	}

	return nil
}

// sliceMaps will try to call the method SliceMap on the given iterator
// It will return an array of map ot an array of errors
func sliceMaps(iter *gocql.Iter) ([]map[string]interface{}, elemental.Errors) {

	maps, err := iter.SliceMap()
	if err != nil {
		return nil, makeError(err.Error(), ErrCannotSlice)
	}

	if err = iter.Close(); err != nil {
		return nil, makeError(err.Error(), ErrCannotCloseIterator)
	}

	return maps, nil
}
