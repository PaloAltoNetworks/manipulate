// Author: Alexandre Wilhelm
// See LICENSE file for full LICENSE
// Copyright 2016 Aporeto.

package manipcassandra

import (
	"fmt"
	"strings"
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

	nativeSession *gocql.Session
	batchRegistry BatchRegistry
}

// NewCassandraStore returns a new *CassandraStore
// You can specify the parameters servers, jeyspace and version.
func NewCassandraStore(servers []string, keyspace string, version int) *CassandraStore {

	return &CassandraStore{
		Servers:       servers,
		KeySpace:      keyspace,
		ProtoVersion:  version,
		batchRegistry: BatchRegistry{},
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

		primaryKeys, primaryValues, err := cassandra.PrimaryFieldsAndValues(object)

		if err != nil {

			log.WithFields(log.Fields{
				"context": context,
				"error":   err,
			}).Debug("Unable to send select command to cassandra.")

			return makeManipCassandraErrors(err.Error(), ManipCassandraPrimaryFieldsAndValuesErrorCode)
		}

		context := manipulate.ContextForIndex(context, index)
		command, values := buildGetCommand(context, object.Identity().Name, primaryKeys, primaryValues)

		log.WithFields(log.Fields{
			"command": command,
			"values":  values,
			"context": context,
		}).Debug("sending select command to cassandra")

		// There is mabe a way to create only one request based on every context
		// TODO: intern work one day :D
		iter := c.nativeSession.Query(command, values...).Iter()
		err = unmarshalManipulable(iter, []manipulate.Manipulable{object})

		if err != nil {
			return makeManipCassandraErrors(err.Error(), ManipCassandraUnmarshalErrorCode)
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
				"context": context,
				"error":   err,
			}).Error("Unable to extract primary keys and values.")

			return makeManipCassandraErrors(err.Error(), ManipCassandraPrimaryFieldsAndValuesErrorCode)
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
		return makeManipCassandraErrors(err.Error(), ManipCassandraExecuteBatchErrorCode)
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
		"command": command,
		"values":  values,
		"context": context,
	}).Debug("sending select all command to cassandra")

	iter := c.nativeSession.Query(command, values...).Iter()

	if err := unmarshalInterface(iter, dest); err != nil {
		return makeManipCassandraErrors(err.Error(), ManipCassandraUnmarshalErrorCode)
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
				"context": context,
				"error":   err,
			}).Error("Unable to extract fields and values.")

			return makeManipCassandraErrors(err.Error(), ManipCassandraFieldsAndValuesErrorCode)
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
		"batch":   batch.Entries,
		"context": context,
	}).Debug("sending create command to cassandra")

	if err := c.nativeSession.ExecuteBatch(batch); err != nil {

		log.WithFields(log.Fields{
			"batch":   batch.Entries,
			"context": context,
			"error":   err,
		}).Error("Unabled to send create command to cassandra.")

		for _, object := range objects {
			object.SetIdentifier("")
		}

		return makeManipCassandraErrors(err.Error(), ManipCassandraExecuteBatchErrorCode)
	}

	log.WithFields(log.Fields{
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
			"context": context,
			"error":   err,
		}).Error("Unable to extract primary fields and values.")

		return makeManipCassandraErrors(err.Error(), ManipCassandraPrimaryFieldsAndValuesErrorCode)
	}

	command, values := buildUpdateCollectionCommand(manipulate.ContextForIndex(context, 0), object.Identity().Name, attributeUpdate, primaryKeys, primaryValues)
	if err := c.nativeSession.Query(command, values...).Exec(); err != nil {

		log.WithFields(log.Fields{
			"context": context,
			"error":   err,
		}).Error("Unable to send update collection command to cassandra.")

		return makeManipCassandraErrors(err.Error(), ManipCassandraQueryErrorCode)
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
				"context": context,
				"error":   err,
			}).Error("Unable to extract primary fields and values.")

			return makeManipCassandraErrors(err.Error(), ManipCassandraPrimaryFieldsAndValuesErrorCode)
		}

		list, values, err := cassandra.FieldsAndValues(object)

		if err != nil {

			log.WithFields(log.Fields{
				"context": context,
				"error":   err,
			}).Debug("Unable to extract fields and values.")

			return makeManipCassandraErrors(err.Error(), ManipCassandraFieldsAndValuesErrorCode)
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
		return makeManipCassandraErrors(err.Error(), ManipCassandraExecuteBatchErrorCode)
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
		"query":   command,
		"context": context,
	}).Debug("sending count command to cassandra")

	iter := c.nativeSession.Query(command, values...).Iter()

	var count int
	success := iter.Scan(&count)

	if !success {
		return -1, makeManipCassandraErrors("Unable to scan collection", ManipCassandraIteratorScanErrorCode)
	}

	if err := iter.Close(); err != nil {
		return -1, makeManipCassandraErrors(err.Error(), ManipCassandraIteratorCloseErrorCode)
	}

	log.WithFields(log.Fields{
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

	defer func() { delete(c.batchRegistry, id) }()

	if c.batchRegistry[id] == nil {
		log.WithFields(log.Fields{
			"store":         c,
			"transactionID": id,
		}).Error("No batch found for the given transaction.")

		return makeManipCassandraErrors("No batch found for the given transaction.", ManipCassandraCommitTransactionErrorCode)
	}

	if err := c.commitBatch(c.batchRegistry[id]); err != nil {
		return makeManipCassandraErrors(err.Error(), ManipCassandraExecuteBatchErrorCode)
	}

	return nil
}

// Abort aborts the given transaction ID.
func (c *CassandraStore) Abort(id manipulate.TransactionID) elemental.Errors {

	if c.batchRegistry[id] == nil {
		log.WithFields(log.Fields{
			"store":         c,
			"transactionID": id,
		}).Error("No batch found for the given transaction.")

		return makeManipCassandraErrors("No batch found for the given transaction.", ManipCassandraCommitTransactionErrorCode)
	}

	delete(c.batchRegistry, id)

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
				"error": err,
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

	session, err := cluster.CreateSession()

	if err != nil {
		return nil, err
	}

	return session, nil
}

// query return a gocql.Query.
// The dev can then do whatever he wants with
func (c *CassandraStore) query(query string, values []interface{}) *gocql.Query {
	return c.nativeSession.Query(query, values...)
}

// batchForID return a gocql.Batch,
// The dev can then do whatever he wants with
// If id is emptyn it will return a new batch
// If id does not match with a batch, it will create a new batch
func (c *CassandraStore) batchForID(id manipulate.TransactionID) *gocql.Batch {

	if id == "" {
		return c.nativeSession.NewBatch(gocql.UnloggedBatch)
	}

	batch := c.batchRegistry[id]

	if batch == nil {
		c.batchRegistry[id] = c.nativeSession.NewBatch(gocql.UnloggedBatch)
	}

	return c.batchRegistry[id]
}

// CommitBatch commit the given batch
// The dev can then do whatever he wants with
func (c *CassandraStore) commitBatch(b *gocql.Batch) error {

	log.WithFields(log.Fields{
		"batch": b.Entries,
	}).Debug("Commiting batch to cassandra.")

	if err := c.nativeSession.ExecuteBatch(b); err != nil {

		log.WithFields(log.Fields{
			"batch": b.Entries,
			"error": err,
		}).Debug("Unable to send batch command.")

		return err
	}

	return nil
}

// sliceMaps will try to call the method SliceMap on the given iterator
// It will return an array of map ot an array of errors
func sliceMaps(iter *gocql.Iter) ([]map[string]interface{}, error) {

	sliceMaps, err := iter.SliceMap()

	if err != nil {
		return nil, err
	}

	if err = iter.Close(); err != nil {
		return nil, err
	}

	return sliceMaps, nil
}

// unmarshalManipulable this will be called by the method retrieve as we know in which object we will store the retrieved data
func unmarshalManipulable(iter *gocql.Iter, objects []manipulate.Manipulable) error {

	sliceMaps, err := sliceMaps(iter)

	if err != nil {
		return err
	}

	if len(objects) != len(sliceMaps) {
		return fmt.Errorf("Unexpected number of objects in unmarshaled data.")
	}

	for index, sliceMap := range sliceMaps {

		// We access the object here and give it to Unmarshal
		// We can't give the array objects as the type is manipulate.Manipulable which is an interface
		// It won't be possible to decode that...or in an easy (or not) way...
		// TODO : intern work one day :D
		if err := cassandra.Unmarshal(sliceMap, objects[index]); err != nil {
			return err
		}
	}

	return nil
}

// unmarshalManipulable this will be called by the method retrieveChildren, we will store the objects in the empty given array/interface
func unmarshalInterface(iter *gocql.Iter, objects interface{}) error {

	sliceMaps, err := sliceMaps(iter)

	if err != nil {
		return err
	}

	if len(sliceMaps) < 1 {
		return nil
	}

	if err := cassandra.Unmarshal(sliceMaps, objects); err != nil {
		return err
	}

	return nil
}
