// Author: Alexandre Wilhelm
// See LICENSE file for full LICENSE
// Copyright 2016 Aporeto.

package manipcassandra

import (
	"errors"
	"fmt"

	"github.com/aporeto-inc/elemental"
	"github.com/aporeto-inc/gocql"
	"github.com/aporeto-inc/manipulate"
	"github.com/aporeto-inc/manipulate/manipcassandra/encoding"

	log "github.com/Sirupsen/logrus"
)

// AttributeUpdater structu used to make request as UPDATE policy SET NAME = NAME - ?
type AttributeUpdater struct {
	Key       string
	Values    interface{}
	Operation elemental.Operation
}

// CassandraStore needs doc
type CassandraStore struct {
	Servers      []string
	KeySpace     string
	ProtoVersion int

	nativeSession *gocql.Session
}

// NewCassandraStore returns a new *CassandraStore
// You can specify the parameters servers, jeyspace and version
func NewCassandraStore(servers []string, keyspace string, version int) *CassandraStore {

	return &CassandraStore{
		Servers:      servers,
		KeySpace:     keyspace,
		ProtoVersion: version,
	}
}

// Stop will close the cassandra session
func (c *CassandraStore) Stop() {
	c.nativeSession.Close()
	c.nativeSession = nil
}

// Start will start the cassandra session
// This method will panic if the host is not reachable
func (c *CassandraStore) Start() {

	cluster := gocql.NewCluster(c.Servers...)
	cluster.Keyspace = c.KeySpace
	cluster.Consistency = gocql.Quorum
	cluster.ProtoVersion = c.ProtoVersion

	session, err := cluster.CreateSession()

	if err != nil {
		log.WithFields(log.Fields{
			"servers":      c.Servers,
			"keyspace":     c.KeySpace,
			"consistency":  gocql.Quorum,
			"protoVersion": c.ProtoVersion,
		}).Info("Fail : creation of a cassandra session failed")

		panic(err)
	}

	c.nativeSession = session
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
				"error":   err.Error(),
			}).Info("Fail : sending select command to cassandra")

			return []*elemental.Error{elemental.NewError("CassandraStore retrieve failed", err.Error(), fmt.Sprintf("%s", object), 500)}
		}

		context := manipulate.ContextForIndex(context, index)
		command, values := buildGetCommand(context, object.Identity().Name, primaryKeys, primaryValues)

		log.WithFields(log.Fields{
			"command": command,
			"context": context,
		}).Info("About : sending select command to cassandra")

		// There is mabe a way to create only one request based on every context
		// TODO: intern work one day :D
		iter := c.nativeSession.Query(command, values...).Iter()
		unmarshalErr := unmarshalManipulable(iter, []manipulate.Manipulable{object})

		if unmarshalErr != nil {
			return unmarshalErr
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

	batch := c.nativeSession.NewBatch(gocql.LoggedBatch)

	for index, object := range objects {

		primaryKeys, primaryValues, err := cassandra.PrimaryFieldsAndValues(object)

		if err != nil {

			log.WithFields(log.Fields{
				"context": context,
				"error":   err.Error(),
			}).Info("Fail : sending select command to cassandra")

			return []*elemental.Error{elemental.NewError("CassandraStore delete failed", err.Error(), fmt.Sprintf("%s", object), 500)}
		}

		context := manipulate.ContextForIndex(context, index)
		command, values := buildDeleteCommand(context, object.Identity().Name, primaryKeys, primaryValues)
		batch.Query(command, values...)
	}

	log.WithFields(log.Fields{
		"batch":   batch.Entries,
		"context": context,
	}).Info("About : sending delete command to cassandra")

	if err := c.nativeSession.ExecuteBatch(batch); err != nil {

		log.WithFields(log.Fields{
			"batch":   batch.Entries,
			"context": context,
			"error":   err.Error(),
		}).Info("Fail : sending delete command to cassandra")

		return []*elemental.Error{elemental.NewError("CassandraStore delete failed", err.Error(), fmt.Sprintf("%s", objects), 500)}
	}

	log.WithFields(log.Fields{
		"batch":   batch.Entries,
		"context": context,
	}).Info("Success : sending delete command to cassandra")

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
		"query":   command,
		"context": context,
	}).Info("About : sending select all command to cassandra")

	iter := c.nativeSession.Query(command, values...).Iter()

	return unmarshalInterface(iter, dest)
}

// Create will launch a set of delete to the cassandra session
// The create commands will be in a bacth
// The given Contexts is either an array of Context or a Context, each Context will map with the given objects
// The given objects will be created in the database
// If the object has an field ID, this field will be automatically set to a new UUID generated by gocql.TimeUUID.String() and will be set by calling the method SetIdentifier on the object
// This method returns an array of Errors if something wrong happens
func (c *CassandraStore) Create(context manipulate.Contexts, parent manipulate.Manipulable, objects ...manipulate.Manipulable) elemental.Errors {

	batch := c.nativeSession.NewBatch(gocql.LoggedBatch)

	for index, object := range objects {
		object.SetIdentifier(gocql.TimeUUID().String())
		list, values, err := cassandra.FieldsAndValues(object)

		if err != nil {

			for _, object := range objects {
				object.SetIdentifier("")
			}

			log.WithFields(log.Fields{
				"batch":   batch.Entries,
				"context": context,
				"error":   err.Error(),
			}).Info("Fail : sending update command to cassandra")

			return []*elemental.Error{elemental.NewError("CassandraStore create failed", err.Error(), fmt.Sprintf("%s", object), 500)}
		}

		context := manipulate.ContextForIndex(context, index)
		command, values := buildInsertCommand(context, object.Identity().Name, list, values)
		batch.Query(command, values...)
	}

	log.WithFields(log.Fields{
		"batch":   batch.Entries,
		"context": context,
	}).Info("About : sending create command to cassandra")

	if err := c.nativeSession.ExecuteBatch(batch); err != nil {

		log.WithFields(log.Fields{
			"batch":   batch.Entries,
			"context": context,
			"error":   err.Error(),
		}).Info("Fail : sending create command to cassandra")

		for _, object := range objects {
			object.SetIdentifier("")
		}

		return []*elemental.Error{elemental.NewError("CassandraStore create failed", err.Error(), fmt.Sprintf("%s", objects), 500)}
	}

	log.WithFields(log.Fields{
		"batch":   batch.Entries,
		"context": context,
	}).Info("Success : sending create command to cassandra")

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
			"error":   err.Error(),
		}).Info("Fail : sending update collection command to cassandra")

		return []*elemental.Error{elemental.NewError("CassandraStore UpdateCollection failed", err.Error(), fmt.Sprintf("%s", object), 500)}
	}

	command, values := buildUpdateCollectionCommand(manipulate.ContextForIndex(context, 0), object.Identity().Name, attributeUpdate, primaryKeys, primaryValues)

	fmt.Println(command, values)
	err = c.nativeSession.Query(command, values...).Exec()

	if err != nil {

		log.WithFields(log.Fields{
			"context": context,
			"error":   err.Error(),
		}).Info("Fail : sending update collection command to cassandra")

		return []*elemental.Error{elemental.NewError("CassandraStore UpdateCollection failed", err.Error(), fmt.Sprintf("%s", object), 500)}
	}

	return nil
}

// Update will launch a set of update to the cassandra session
// The update commands will be in a batch
// The given Contexts is either an array of Context or a Context, each Context will map with the given objects
// The given objects will be updated in the database
// This method returns an array of Errors if something wrong happens
func (c *CassandraStore) Update(context manipulate.Contexts, objects ...manipulate.Manipulable) elemental.Errors {

	batch := c.nativeSession.NewBatch(gocql.LoggedBatch)

	for index, object := range objects {

		primaryKeys, primaryValues, err := cassandra.PrimaryFieldsAndValues(object)

		if err != nil {

			log.WithFields(log.Fields{
				"context": context,
				"error":   err.Error(),
			}).Info("Fail : sending select command to cassandra")

			return []*elemental.Error{elemental.NewError("CassandraStore update failed", err.Error(), fmt.Sprintf("%s", object), 500)}
		}

		list, values, err := cassandra.FieldsAndValues(object)

		if err != nil {

			log.WithFields(log.Fields{
				"batch":   batch.Entries,
				"context": context,
				"error":   err.Error(),
			}).Info("Fail : sending update command to cassandra")

			return []*elemental.Error{elemental.NewError("CassandraStore update failed", err.Error(), fmt.Sprintf("%s", object), 500)}
		}

		context := manipulate.ContextForIndex(context, index)
		command, values := buildUpdateCommand(context, object.Identity().Name, list, values, primaryKeys, primaryValues)

		batch.Query(command, values...)
	}

	log.WithFields(log.Fields{
		"batch":   batch.Entries,
		"context": context,
	}).Info("About : sending update command to cassandra")

	if err := c.nativeSession.ExecuteBatch(batch); err != nil {

		log.WithFields(log.Fields{
			"batch":   batch.Entries,
			"context": context,
			"error":   err.Error(),
		}).Info("Fail : sending update command to cassandra")

		return []*elemental.Error{elemental.NewError("CassandraStore update failed", err.Error(), fmt.Sprintf("%s", objects), 500)}
	}

	log.WithFields(log.Fields{
		"batch":   batch.Entries,
		"context": context,
	}).Info("Success : sending update command to cassandra")

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
	}).Info("About : sending count command to cassandra")

	iter := c.nativeSession.Query(command, values...).Iter()

	var count int
	success := iter.Scan(&count)

	if !success {
		return -1, []*elemental.Error{elemental.NewError("CassandraStore Iter Close error", errors.New("Error when scanning the iterator of a count command").Error(), identity.Name, 500)}
	}

	if err := iter.Close(); err != nil {
		return -1, []*elemental.Error{elemental.NewError("CassandraStore Iter Close error", errors.New("Error when closing the iterator of a count command").Error(), identity.Name, 500)}
	}

	log.WithFields(log.Fields{
		"query":   command,
		"context": context,
	}).Info("Success : sending count command to cassandra")

	return count, nil
}

func (c *CassandraStore) Assign(contexts manipulate.Contexts, parent manipulate.Manipulable, assignation *elemental.Assignation) elemental.Errors {
	panic("Not implemented")
}

// Query return a gocql.Query.
// The dev can then do whatever he wants with
func (c *CassandraStore) Query(query string, values []interface{}) *gocql.Query {
	return c.nativeSession.Query(query, values...)
}

// Batch return a new gocql.Batch
// The dev can then do whatever he wants with
func (c *CassandraStore) Batch() *gocql.Batch {
	return c.nativeSession.NewBatch(gocql.LoggedBatch)
}

// ExecuteBatch execute the given batch
// The dev can then do whatever he wants with
func (c *CassandraStore) ExecuteBatch(b *gocql.Batch) error {

	log.WithFields(log.Fields{
		"batch": b.Entries,
	}).Info("Success : sending commands to cassandra")

	return c.nativeSession.ExecuteBatch(b)
}

// sliceMaps will try to call the method SliceMap on the given iterator
// It will return an array of map ot an array of errors
func sliceMaps(iter *gocql.Iter) ([]map[string]interface{}, elemental.Errors) {

	sliceMaps, err := iter.SliceMap()

	if err != nil {
		return nil, []*elemental.Error{elemental.NewError("CassandraStore SliceMap error", err.Error(), "", 500)}
	}

	if err = iter.Close(); err != nil {
		return nil, []*elemental.Error{elemental.NewError("CassandraStore Iter Close error", err.Error(), "", 500)}
	}

	return sliceMaps, nil
}

// unmarshalManipulable this will be called by the method retrieve as we know in which object we will store the retrieved data
func unmarshalManipulable(iter *gocql.Iter, objects []manipulate.Manipulable) elemental.Errors {

	sliceMaps, err := sliceMaps(iter)

	if err != nil {
		log.Info(err)
		return err
	}

	if len(objects) != len(sliceMaps) {
		log.Info("The number of the given objects and the number of results fetched is different")
		return []*elemental.Error{elemental.NewError("CassandraStore Unmarshal error", errors.New("The number of the given objects and the number of results fetched is different").Error(), fmt.Sprintf("%s", objects), 500)}
	}

	for index, sliceMap := range sliceMaps {

		// We access the object here and give it to Unmarshal
		// We can't give the array objects as the type is manipulate.Manipulable which is an interface
		// It won't be possible to decode that...or in an easy (or not) way...
		// TODO : intern work one day :D
		err := cassandra.Unmarshal(sliceMap, objects[index])

		if err != nil {
			log.Info(err)
			return []*elemental.Error{elemental.NewError("CassandraStore Unmarshal error", err.Error(), fmt.Sprintf("%s", objects[index]), 500)}
		}
	}

	return nil
}

// unmarshalManipulable this will be called by the method retrieveChildren, we will store the objects in the empty given array/interface
func unmarshalInterface(iter *gocql.Iter, objects interface{}) elemental.Errors {

	sliceMaps, err := sliceMaps(iter)

	if err != nil {
		log.Info(err)
		return err
	}

	if len(sliceMaps) < 1 {
		return nil
	}

	if err := cassandra.Unmarshal(sliceMaps, objects); err != nil {
		log.Info(err)
		return []*elemental.Error{elemental.NewError("CassandraStore Unmarshal error", err.Error(), fmt.Sprintf("%s", objects), 500)}
	}

	return nil
}
