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
	"github.com/aporeto-inc/manipulate/manipcassandra/compiler"
)

const sep = ", "

// commandAndValuesFromContext applies a context to the given command buffer.
// it supports limits, filters (where) and parameters (IF NOT EXISTS)
func commandAndValuesFromContext(buffer *bytes.Buffer, operation elemental.Operation, c *manipulate.Context, primaryKeys []string) (string, []interface{}) {

	values := []interface{}{}
	numberOfPrimaryKeys := len(primaryKeys)
	hasPrimaryKey := false

	for i := 0; i < numberOfPrimaryKeys; i++ {

		if i == 0 {
			manipulate.WriteString(buffer, " WHERE ")
		} else {
			manipulate.WriteString(buffer, " AND ")
		}

		manipulate.WriteString(buffer, primaryKeys[i])
		manipulate.WriteString(buffer, " = ?")

		hasPrimaryKey = true
	}

	if c == nil {
		return buffer.String(), []interface{}{}
	}

	if c.Filter != nil {
		manipulate.WriteString(buffer, ` `)

		filterString := compiler.CompileFilter(c.Filter)

		if hasPrimaryKey {
			filterString = strings.Replace(filterString, "WHERE", "AND", 1)
		}

		manipulate.WriteString(buffer, filterString)

		filter := c.Filter

		for i := 0; i < len(filter.Values()); i++ {

			v := getValues(filter.Values()[i])
			values = append(values, v...)
		}
	}

	if c.PageSize > 0 && (operation == elemental.OperationRetrieveMany || operation == elemental.OperationInfo) {
		manipulate.WriteString(buffer, ` LIMIT `)
		manipulate.WriteString(buffer, strconv.Itoa(c.PageSize))
	}

	parameterString := compiler.CompileParameters(c.Parameters)
	if len(parameterString) > 0 {
		manipulate.WriteString(buffer, ` `+parameterString)
	}

	return buffer.String(), values
}

// buildUpdateCollectionCommand build a update command for cassandra
// example : UPDATE policy SET NAME = NAME - ?  WHERE ID = ?
// every values will be replace by a ?
// then it will apply the given context on the query
func buildUpdateCollectionCommand(c *manipulate.Context, tableName string, attributeUpdate *attributeUpdater, primaryKeys []string, primaryValues []interface{}) (string, []interface{}) {

	var buffer bytes.Buffer
	manipulate.WriteString(&buffer, "UPDATE ")
	manipulate.WriteString(&buffer, tableName)
	manipulate.WriteString(&buffer, " SET ")
	manipulate.WriteString(&buffer, attributeUpdate.Key)

	switch attributeUpdate.AssignationType {
	case elemental.AssignationTypeSet:
		manipulate.WriteString(&buffer, " = ")
		break

	case elemental.AssignationTypeAdd:
		manipulate.WriteString(&buffer, " = ")
		manipulate.WriteString(&buffer, attributeUpdate.Key)
		manipulate.WriteString(&buffer, " + ")
		break

	case elemental.AssignationTypeSubstract:
		manipulate.WriteString(&buffer, " = ")
		manipulate.WriteString(&buffer, attributeUpdate.Key)
		manipulate.WriteString(&buffer, " - ")
		break

	default:
		manipulate.WriteString(&buffer, " = ")
		break
	}

	manipulate.WriteString(&buffer, "?")

	var v []interface{}

	v = append(v, attributeUpdate.Values)
	v = append(v, primaryValues...)
	command, newValues := commandAndValuesFromContext(&buffer, elemental.OperationUpdate, c, primaryKeys)

	return command + " ALLOW FILTERING", append(v, newValues...)
}

// buildUpdateCommand build a update command for cassandra
// example : Update tag set name = ? id = ?
// every values will be replace by a ?
// then it will apply the given context on the query
func buildUpdateCommand(c *manipulate.Context, tableName string, p []string, v []interface{}, primaryKeys []string, primaryValues []interface{}) (string, []interface{}) {

	var buffer bytes.Buffer
	manipulate.WriteString(&buffer, "UPDATE ")
	manipulate.WriteString(&buffer, tableName)
	manipulate.WriteString(&buffer, " SET ")

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
			manipulate.WriteString(&buffer, sep)
		} else {
			valuesInserted++
		}

		manipulate.WriteString(&buffer, k)
		manipulate.WriteString(&buffer, " = ?")
	}

	v = append(v, primaryValues...)
	command, newValues := commandAndValuesFromContext(&buffer, elemental.OperationUpdate, c, primaryKeys)

	return command, append(v, newValues...)
}

// buildInsertCommand build an insert command for cassandra
// example : INSERT INTO tag (ID, name) VALUES (?, ?)
// every values will be replace by a ?
// then it will apply the given context on the query
func buildInsertCommand(c *manipulate.Context, tableName string, p []string, v []interface{}) (string, []interface{}) {

	var buffer bytes.Buffer
	manipulate.WriteString(&buffer, "INSERT INTO ")
	manipulate.WriteString(&buffer, tableName)
	manipulate.WriteString(&buffer, " (")

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
			manipulate.WriteString(&buffer, sep)
		}

		manipulate.WriteString(&buffer, k)
		valuesInserted++
	}

	manipulate.WriteString(&buffer, ") VALUES (")

	for index := 0; index < valuesInserted; index++ {

		if index > 0 {
			manipulate.WriteString(&buffer, sep)
		}

		manipulate.WriteString(&buffer, "?")
	}

	manipulate.WriteString(&buffer, ")")
	command, newValues := commandAndValuesFromContext(&buffer, elemental.OperationCreate, c, []string{})

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
		manipulate.WriteString(&buffer, `SELECT * FROM `)
	} else {
		manipulate.WriteString(&buffer, `SELECT `)

		for i := 0; i < len(attributes); i++ {
			attribute := attributes[i]

			if i > 0 {
				manipulate.WriteString(&buffer, sep)
			}

			manipulate.WriteString(&buffer, attribute)
		}

		manipulate.WriteString(&buffer, ` FROM `)
	}

	manipulate.WriteString(&buffer, tableName)

	command, values := commandAndValuesFromContext(&buffer, elemental.OperationRetrieveMany, c, primaryKeys)

	return command + " ALLOW FILTERING", append(primaryValues, values...)
}

// buildDeleteCommand build a delete command for cassandra
// example : DELETE FROM tag
// every values will be replace by a ?
// then it will apply the given context on the query
func buildDeleteCommand(c *manipulate.Context, tableName string, primaryKeys []string, primaryValues []interface{}) (string, []interface{}) {

	var buffer bytes.Buffer
	manipulate.WriteString(&buffer, `DELETE FROM `)
	manipulate.WriteString(&buffer, tableName)

	command, values := commandAndValuesFromContext(&buffer, elemental.OperationDelete, c, primaryKeys)

	return command, append(primaryValues, values...)
}

// buildCountCommand build a count command for cassandra
// example : SELECT * FROM tag
// every values will be replace by a ?
// then it will apply the given context on the query
func buildCountCommand(c *manipulate.Context, tableName string) (string, []interface{}) {

	var buffer bytes.Buffer
	manipulate.WriteString(&buffer, `SELECT COUNT(*) FROM `)
	manipulate.WriteString(&buffer, tableName)

	command, values := commandAndValuesFromContext(&buffer, elemental.OperationInfo, c, []string{})

	return command + " ALLOW FILTERING", values
}

// buildIncrementCommand build a counter incrementation command for cassandra
// example : UPDATE counter_table_name SET count = count + n WHERE k = x
// every values will be replace by a ?
// then it will apply the given context on the query
func buildIncrementCommand(c *manipulate.Context, tableName, counterName string, inc int, primaryKeys []string, primaryValues []interface{}) (string, []interface{}) {

	var buffer bytes.Buffer
	manipulate.WriteString(&buffer, `UPDATE `)
	manipulate.WriteString(&buffer, tableName)
	manipulate.WriteString(&buffer, ` SET `)
	manipulate.WriteString(&buffer, counterName)
	manipulate.WriteString(&buffer, ` = `)
	manipulate.WriteString(&buffer, counterName)
	manipulate.WriteString(&buffer, ` + `)
	manipulate.WriteString(&buffer, strconv.Itoa(inc))

	command, values := commandAndValuesFromContext(&buffer, elemental.OperationUpdate, c, primaryKeys)

	return command, append(primaryValues, values...)
}
