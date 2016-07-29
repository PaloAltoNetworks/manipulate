// Author: Alexandre Wilhelm
// See LICENSE file for full LICENSE
// Copyright 2016 Aporeto.

package manipcassandra

import (
	"bytes"
	"strconv"
	"strings"

	"github.com/aporeto-inc/elemental"
	"github.com/aporeto-inc/manipulate"
)

const sep = ", "

// commandAndValuesFromContext will apply a context to the given command
// it now supports limit, filer (where) and parameters (IF NOT EXISTS)
func commandAndValuesFromContext(command string, c *manipulate.Context, primaryKeys []string) (string, []interface{}) {

	values := []interface{}{}
	numberOfPrimaryKeys := len(primaryKeys)
	hasPrimaryKey := false

	var buffer bytes.Buffer
	buffer.WriteString(command)

	for i := 0; i < numberOfPrimaryKeys; i++ {

		if i == 0 {
			buffer.WriteString(" WHERE ")
		} else {
			buffer.WriteString(" AND ")
		}

		buffer.WriteString(primaryKeys[i])
		buffer.WriteString(" = ?")

		hasPrimaryKey = true
	}

	if c == nil {
		return buffer.String(), []interface{}{}
	}

	if c.Filter != nil {
		buffer.WriteString(` `)

		filterString := c.Filter.Compile()

		if hasPrimaryKey {
			filterString = strings.Replace(filterString, "WHERE", "AND", 1)
		}

		buffer.WriteString(filterString)

		filter := c.Filter.(*Filter)

		for i := 0; i < len(filter.Values); i++ {

			v := getValues(filter.Values[i])
			values = append(values, v...)
		}
	}

	if c.PageSize > 0 {
		buffer.WriteString(` LIMIT `)
		buffer.WriteString(strconv.Itoa(c.PageSize))
	}

	if c.Parameter != nil {
		buffer.WriteString(` `)
		buffer.WriteString(c.Parameter.Compile())
	}

	if c.Filter != nil && c.Filter.(*Filter).AllowFiltering {
		buffer.WriteString(` ALLOW FILTERING`)
	}

	return buffer.String(), values
}

// buildUpdateCollectionCommand build a update command for cassandra
// example : UPDATE policy SET NAME = NAME - ?  WHERE ID = ?
// every values will be replace by a ?
// then it will apply the given context on the query
func buildUpdateCollectionCommand(c *manipulate.Context, tableName string, attributeUpdate *AttributeUpdater, primaryKeys []string, primaryValues []interface{}) (string, []interface{}) {

	var buffer bytes.Buffer
	buffer.WriteString("UPDATE ")
	buffer.WriteString(tableName)
	buffer.WriteString(" SET ")
	buffer.WriteString(attributeUpdate.Key)

	switch attributeUpdate.AssignationType {
	case elemental.AssignationTypeSet:
		buffer.WriteString(" = ")
		break

	case elemental.AssignationTypeAdd:
		buffer.WriteString(" = ")
		buffer.WriteString(attributeUpdate.Key)
		buffer.WriteString(" + ")
		break

	case elemental.AssignationTypeSubstract:
		buffer.WriteString(" = ")
		buffer.WriteString(attributeUpdate.Key)
		buffer.WriteString(" - ")
		break

	default:
		buffer.WriteString(" = ")
		break
	}

	buffer.WriteString("?")

	var v []interface{}

	v = append(v, attributeUpdate.Values)
	v = append(v, primaryValues...)
	command, newValues := commandAndValuesFromContext(buffer.String(), c, primaryKeys)

	return command, append(v, newValues...)
}

// buildUpdateCommand build a update command for cassandra
// example : Update tag set name = ? id = ?
// every values will be replace by a ?
// then it will apply the given context on the query
func buildUpdateCommand(c *manipulate.Context, tableName string, p []string, v []interface{}, primaryKeys []string, primaryValues []interface{}) (string, []interface{}) {

	var buffer bytes.Buffer
	buffer.WriteString("UPDATE ")
	buffer.WriteString(tableName)
	buffer.WriteString(" SET ")

	var attributes []string

	if c != nil {
		attributes = c.Attributes
	}

	shouldCheckAttributes := len(attributes) > 0
	valuesDeleted := 0
	valuesInserted := 0

	for index, k := range p {

		if (shouldCheckAttributes && !stringInSlice(k, attributes)) || stringInSlice(k, primaryKeys) {
			position := index - valuesDeleted
			v = append(v[:position], v[position+1:]...)
			valuesDeleted++
			continue
		}

		if valuesInserted > 0 {
			buffer.WriteString(sep)
		} else {
			valuesInserted++
		}

		buffer.WriteString(k)
		buffer.WriteString(" = ?")
	}

	v = append(v, primaryValues...)
	command, newValues := commandAndValuesFromContext(buffer.String(), c, primaryKeys)

	return command, append(v, newValues...)
}

// buildInsertCommand build an insert command for cassandra
// example : INSERT INTO tag (ID, name) VALUES (?, ?)
// every values will be replace by a ?
// then it will apply the given context on the query
func buildInsertCommand(c *manipulate.Context, tableName string, p []string, v []interface{}) (string, []interface{}) {

	var buffer bytes.Buffer
	buffer.WriteString("INSERT INTO ")
	buffer.WriteString(tableName)
	buffer.WriteString(" (")

	var attributes []string

	if c != nil {
		attributes = c.Attributes
	}

	shouldCheckAttributes := len(attributes) > 0
	valuesDeleted := 0
	valuesInserted := 0

	for index, k := range p {

		if shouldCheckAttributes && !stringInSlice(k, attributes) {
			position := index - valuesDeleted
			v = append(v[:position], v[position+1:]...)
			valuesDeleted++
			continue
		}

		if valuesInserted > 0 {
			buffer.WriteString(sep)
		}

		buffer.WriteString(k)
		valuesInserted++
	}

	buffer.WriteString(") VALUES (")

	for index := 0; index < valuesInserted; index++ {

		if index > 0 {
			buffer.WriteString(sep)
		}

		buffer.WriteString("?")
	}

	buffer.WriteString(")")
	command, newValues := commandAndValuesFromContext(buffer.String(), c, []string{})

	return command, append(v, newValues...)
}

// buildGetCommand build a delete command for cassandra
// example : SELECT * FROM tag
// every values will be replace by a ?
// then it will apply the given context on the query
func buildGetCommand(c *manipulate.Context, tableName string, primaryKeys []string, primaryValues []interface{}) (string, []interface{}) {

	var buffer bytes.Buffer

	var attributes []string

	if c != nil {
		attributes = c.Attributes
	}

	if len(attributes) == 0 {
		buffer.WriteString(`SELECT * FROM `)
	} else {
		buffer.WriteString(`SELECT `)

		for i := 0; i < len(attributes); i++ {
			attribute := attributes[i]

			if i > 0 {
				buffer.WriteString(sep)
			}

			buffer.WriteString(attribute)
		}

		buffer.WriteString(` FROM `)
	}

	buffer.WriteString(tableName)

	command, values := commandAndValuesFromContext(buffer.String(), c, primaryKeys)

	return command, append(primaryValues, values...)
}

// buildDeleteCommand build a delete command for cassandra
// example : DELETE FROM tag
// every values will be replace by a ?
// then it will apply the given context on the query
func buildDeleteCommand(c *manipulate.Context, tableName string, primaryKeys []string, primaryValues []interface{}) (string, []interface{}) {

	var buffer bytes.Buffer
	buffer.WriteString(`DELETE FROM `)
	buffer.WriteString(tableName)

	command, values := commandAndValuesFromContext(buffer.String(), c, primaryKeys)

	return command, append(primaryValues, values...)
}

// buildCountCommand build a count command for cassandra
// example : SELECT * FROM tag
// every values will be replace by a ?
// then it will apply the given context on the query
func buildCountCommand(c *manipulate.Context, tableName string) (string, []interface{}) {

	var buffer bytes.Buffer
	buffer.WriteString(`SELECT COUNT(*) FROM `)
	buffer.WriteString(tableName)

	return commandAndValuesFromContext(buffer.String(), c, []string{})
}
