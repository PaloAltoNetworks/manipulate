// Author: Alexandre Wilhelm
// See LICENSE file for full LICENSE
// Copyright 2016 Aporeto.

package manipcassandra

import (
	"bytes"
	"testing"

	"github.com/aporeto-inc/elemental"
	"github.com/aporeto-inc/manipulate"
	. "github.com/smartystreets/goconvey/convey"
)

func TestMethodAddOptionsFromContextWithPrimaryKeys(t *testing.T) {

	Convey("Given I call the method addOptionsFromContext", t, func() {

		query := bytes.NewBufferString(`SELECT * FROM policy`)
		context := manipulate.NewContext()

		command, values := commandAndValuesFromContext(query, elemental.OperationRetrieveMany, context, []string{"name", "ID"})

		So(command, ShouldEqual, `SELECT * FROM policy WHERE name = ? AND ID = ? LIMIT 100`)
		So(values, ShouldResemble, []interface{}{})
	})
}

func TestMethodAddOptionsFromContextWithLimitEqualTo0(t *testing.T) {

	Convey("Given I call the method addOptionsFromContext", t, func() {

		query := bytes.NewBufferString(`SELECT * FROM policy`)
		context := manipulate.NewContext()

		command, values := commandAndValuesFromContext(query, elemental.OperationRetrieveMany, context, []string{})

		So(command, ShouldEqual, `SELECT * FROM policy LIMIT 100`)
		So(values, ShouldResemble, []interface{}{})
	})
}

func TestMethodAddOptionsFromContextWithNilContext(t *testing.T) {

	Convey("Given I call the method addOptionsFromContext", t, func() {

		query := bytes.NewBufferString(`SELECT * FROM policy`)
		command, values := commandAndValuesFromContext(query, elemental.OperationRetrieveMany, nil, []string{})

		So(command, ShouldEqual, `SELECT * FROM policy`)
		So(values, ShouldResemble, []interface{}{})
	})
}

func TestMethodAddOptionsFromContextWithLimitEqualTo20(t *testing.T) {

	Convey("Given I call the method addOptionsFromContext", t, func() {
		query := bytes.NewBufferString(`SELECT * FROM policy`)
		context := manipulate.NewContext()
		context.PageSize = 20
		command, values := commandAndValuesFromContext(query, elemental.OperationRetrieveMany, context, []string{})

		So(command, ShouldEqual, `SELECT * FROM policy LIMIT 20`)
		So(values, ShouldResemble, []interface{}{})
	})
}

func TestMethodAddOptionsFromContextWithParameter(t *testing.T) {

	Convey("Given I call the method addOptionsFromContext", t, func() {
		query := bytes.NewBufferString(`SELECT * FROM policy`)
		context := manipulate.NewContext()
		context.Parameter = &Parameter{IfNotExists: true}
		command, values := commandAndValuesFromContext(query, elemental.OperationRetrieveMany, context, []string{})

		So(command, ShouldEqual, `SELECT * FROM policy LIMIT 100 IF NOT EXISTS `)
		So(values, ShouldResemble, []interface{}{})
	})
}

func TestMethodAddOptionsFromContextWithFilter(t *testing.T) {

	Convey("Given I call the method addOptionsFromContext", t, func() {
		query := bytes.NewBufferString(`SELECT * FROM policy`)
		context := manipulate.NewContext()

		filter := &Filter{}
		filter.Keys = [][]string{[]string{"ID"}}
		filter.Separators = []string{CassandraFilterEqualSeparator}
		filter.Values = [][]interface{}{[]interface{}{"12345"}}

		context.Filter = filter
		command, values := commandAndValuesFromContext(query, elemental.OperationRetrieveMany, context, []string{})

		So(command, ShouldEqual, `SELECT * FROM policy WHERE ID = ? LIMIT 100`)
		So(values, ShouldResemble, []interface{}{"12345"})
	})
}

func TestMethodAddOptionsFromContextWithEverything(t *testing.T) {

	Convey("Given I call the method addOptionsFromContext", t, func() {
		query := bytes.NewBufferString(`SELECT * FROM policy`)
		context := manipulate.NewContext()

		filter := &Filter{}
		filter.Keys = [][]string{[]string{"ID"}}
		filter.Separators = []string{CassandraFilterEqualSeparator}
		filter.Values = [][]interface{}{[]interface{}{"12345"}}
		filter.AllowFiltering = true

		context.Filter = filter
		context.PageSize = 20

		context.Parameter = &Parameter{IfNotExists: true}

		command, values := commandAndValuesFromContext(query, elemental.OperationRetrieveMany, context, []string{"name", "age"})

		So(command, ShouldEqual, `SELECT * FROM policy WHERE name = ? AND age = ? AND ID = ? LIMIT 20 IF NOT EXISTS  ALLOW FILTERING`)
		So(values, ShouldResemble, []interface{}{"12345"})
	})
}

func TestMethodAddOptionsFromContextWithMultiColumnAndValues(t *testing.T) {

	Convey("Given I call the method addOptionsFromContext", t, func() {
		query := bytes.NewBufferString(`SELECT * FROM policy`)
		context := manipulate.NewContext()

		filter := &Filter{}
		filter.Keys = [][]string{[]string{"ID", "name"}}
		filter.Separators = []string{CassandraFilterEqualSeparator}
		filter.Values = [][]interface{}{[]interface{}{[]interface{}{"20", 0}, []interface{}{"60", 122}}}

		context.Filter = filter

		command, values := commandAndValuesFromContext(query, elemental.OperationRetrieveMany, context, []string{})

		So(command, ShouldEqual, `SELECT * FROM policy WHERE (ID,name) = ((?,?),(?,?)) LIMIT 100`)
		So(values, ShouldResemble, []interface{}{"20", 0, "60", 122})
	})
}

func TestMethodBuildDeleteCommand(t *testing.T) {

	Convey("Given I call the method buildDeleteCommand", t, func() {

		command, values := buildDeleteCommand(nil, "policy", []string{}, []interface{}{})

		So(command, ShouldEqual, `DELETE FROM policy`)
		So(values, ShouldResemble, []interface{}{})
	})
}

func TestMethodBuildDeleteCommandWithPrimaryKeysAnsValues(t *testing.T) {

	Convey("Given I call the method buildDeleteCommand", t, func() {

		command, values := buildDeleteCommand(nil, "policy", []string{"ID", "name"}, []interface{}{"123", "Alexandre"})

		So(command, ShouldEqual, `DELETE FROM policy WHERE ID = ? AND name = ?`)
		So(values, ShouldResemble, []interface{}{"123", "Alexandre"})
	})
}

func TestMethodBuildDeleteCommandWithPrimaryKeysAnsValuesAndFilter(t *testing.T) {

	Convey("Given I call the method buildDeleteCommand", t, func() {

		context := manipulate.NewContext()
		filter := &Filter{}
		filter.Keys = [][]string{[]string{"ID", "name"}}
		filter.Separators = []string{CassandraFilterEqualSeparator}
		filter.Values = [][]interface{}{[]interface{}{[]interface{}{"20", 0}, []interface{}{"60", 122}}}
		context.Filter = filter
		context.PageSize = -1

		command, values := buildDeleteCommand(context, "policy", []string{"ID", "name"}, []interface{}{"123", "Alexandre"})

		So(command, ShouldEqual, `DELETE FROM policy WHERE ID = ? AND name = ? AND (ID,name) = ((?,?),(?,?))`)
		So(values, ShouldResemble, []interface{}{"123", "Alexandre", "20", 0, "60", 122})
	})
}

func TestMethodBuildUpdateCollectionCommandOperationAdditive(t *testing.T) {

	Convey("Given I call the method buildUpdateCollectionCommand", t, func() {

		a := &AttributeUpdater{}
		a.Key = "NAME"
		a.AssignationType = elemental.AssignationTypeAdd
		a.Values = "coucou"

		command, values := buildUpdateCollectionCommand(nil, "policy", a, []string{}, []interface{}{})
		So(command, ShouldEqual, `UPDATE policy SET NAME = NAME + ?`)
		So(values, ShouldResemble, []interface{}{"coucou"})
	})
}

func TestMethodBuildUpdateCollectionCommandOperationDefault(t *testing.T) {

	Convey("Given I call the method buildUpdateCollectionCommand", t, func() {

		a := &AttributeUpdater{}
		a.Key = "NAME"
		a.Values = "coucou"

		command, values := buildUpdateCollectionCommand(nil, "policy", a, []string{}, []interface{}{})
		So(command, ShouldEqual, `UPDATE policy SET NAME = ?`)
		So(values, ShouldResemble, []interface{}{"coucou"})
	})
}

func TestMethodBuildUpdateCollectionCommandOperationSubstractive(t *testing.T) {

	Convey("Given I call the method buildUpdateCollectionCommand", t, func() {

		a := &AttributeUpdater{}
		a.Key = "NAME"
		a.AssignationType = elemental.AssignationTypeSubstract
		a.Values = "coucou"

		command, values := buildUpdateCollectionCommand(nil, "policy", a, []string{}, []interface{}{})
		So(command, ShouldEqual, `UPDATE policy SET NAME = NAME - ?`)
		So(values, ShouldResemble, []interface{}{"coucou"})
	})
}

func TestMethodBuildUpdateCollectionCommandOperationSet(t *testing.T) {

	Convey("Given I call the method buildUpdateCollectionCommand", t, func() {

		a := &AttributeUpdater{}
		a.Key = "NAME"
		a.AssignationType = elemental.AssignationTypeSet
		a.Values = "coucou"

		command, values := buildUpdateCollectionCommand(nil, "policy", a, []string{}, []interface{}{})
		So(command, ShouldEqual, `UPDATE policy SET NAME = ?`)
		So(values, ShouldResemble, []interface{}{"coucou"})
	})
}

func TestMethodBuildUpdateCollectionCommandOperationSubstractiveWithPrimaryKeys(t *testing.T) {

	Convey("Given I call the method buildUpdateCollectionCommand", t, func() {

		a := &AttributeUpdater{}
		a.Key = "NAME"
		a.AssignationType = elemental.AssignationTypeSubstract
		a.Values = "coucou"

		command, values := buildUpdateCollectionCommand(nil, "policy", a, []string{"ID"}, []interface{}{"123"})
		So(command, ShouldEqual, `UPDATE policy SET NAME = NAME - ? WHERE ID = ?`)
		So(values, ShouldResemble, []interface{}{"coucou", "123"})
	})
}

func TestMethodBuildUpdateCollectionCommandOperationSubstractiveWithPrimaryKeysAndContext(t *testing.T) {

	Convey("Given I call the method buildUpdateCollectionCommand", t, func() {

		a := &AttributeUpdater{}
		a.Key = "NAME"
		a.AssignationType = elemental.AssignationTypeSubstract
		a.Values = "coucou"

		context := manipulate.NewContext()
		context.Attributes = []string{"CITY"}
		context.PageSize = 0

		filter := &Filter{}
		filter.Keys = [][]string{[]string{"ID", "name"}}
		filter.Separators = []string{CassandraFilterEqualSeparator}
		filter.Values = [][]interface{}{[]interface{}{[]interface{}{"20", 0}, []interface{}{"60", 122}}}
		context.Filter = filter

		command, values := buildUpdateCollectionCommand(context, "policy", a, []string{"ID"}, []interface{}{"123"})
		So(command, ShouldEqual, `UPDATE policy SET NAME = NAME - ? WHERE ID = ? AND (ID,name) = ((?,?),(?,?))`)
		So(values, ShouldResemble, []interface{}{"coucou", "123", "20", 0, "60", 122})
	})
}

func TestMethodBuildUpdateCommand(t *testing.T) {

	Convey("Given I call the method buildUpdateCommand", t, func() {

		command, values := buildUpdateCommand(nil, "policy", []string{"NAME", "CITY"}, []interface{}{}, []string{}, []interface{}{})
		So(command, ShouldEqual, `UPDATE policy SET NAME = ?, CITY = ?`)
		So(values, ShouldResemble, []interface{}{})
	})
}

func TestMethodBuildUpdateCommandWithPrimaryKeysAndValues(t *testing.T) {

	Convey("Given I call the method buildUpdateCommand", t, func() {

		command, values := buildUpdateCommand(nil, "policy", []string{"NAME", "CITY"}, []interface{}{}, []string{"ID"}, []interface{}{"123"})
		So(command, ShouldEqual, `UPDATE policy SET NAME = ?, CITY = ? WHERE ID = ?`)
		So(values, ShouldResemble, []interface{}{"123"})
	})
}

func TestMethodBuildUpdateCommandWithAttributes(t *testing.T) {

	Convey("Given I call the method buildUpdateCommand", t, func() {

		context := manipulate.NewContext()
		context.Attributes = []string{"CITY"}
		context.PageSize = 0

		command, values := buildUpdateCommand(context, "policy", []string{"NAME", "CITY", "DESCRIPTION"}, []interface{}{"Alexandre", "Sarralbe", "God"}, []string{}, []interface{}{})
		So(command, ShouldEqual, `UPDATE policy SET CITY = ?`)
		So(values, ShouldResemble, []interface{}{"Sarralbe"})
	})
}

func TestMethodBuildUpdateCommandWithAttributesWithPrimaryKeysAndValues(t *testing.T) {

	Convey("Given I call the method buildUpdateCommand", t, func() {

		context := manipulate.NewContext()
		context.Attributes = []string{"CITY", "ID"}
		context.PageSize = 0

		command, values := buildUpdateCommand(context, "policy", []string{"NAME", "CITY", "DESCRIPTION", "ID"}, []interface{}{"Alexandre", "Sarralbe", "God", "567"}, []string{"ID"}, []interface{}{"123"})
		So(command, ShouldEqual, `UPDATE policy SET CITY = ? WHERE ID = ?`)
		So(values, ShouldResemble, []interface{}{"Sarralbe", "123"})
	})
}

func TestMethodBuildUpdateCommandWithAttributesWithPrimaryKeysAndValuesAndFilter(t *testing.T) {

	Convey("Given I call the method buildUpdateCommand", t, func() {

		context := manipulate.NewContext()
		context.Attributes = []string{"CITY", "ID"}
		context.PageSize = 0
		filter := &Filter{}
		filter.Keys = [][]string{[]string{"ID", "name"}}
		filter.Separators = []string{CassandraFilterEqualSeparator}
		filter.Values = [][]interface{}{[]interface{}{[]interface{}{"20", 0}, []interface{}{"60", 122}}}
		context.Filter = filter
		context.PageSize = -1

		command, values := buildUpdateCommand(context, "policy", []string{"NAME", "CITY", "DESCRIPTION", "ID"}, []interface{}{"Alexandre", "Sarralbe", "God", "567"}, []string{"ID"}, []interface{}{"123"})
		So(command, ShouldEqual, `UPDATE policy SET CITY = ? WHERE ID = ? AND (ID,name) = ((?,?),(?,?))`)
		So(values, ShouldResemble, []interface{}{"Sarralbe", "123", "20", 0, "60", 122})
	})
}

func TestMethodBuildInsertCommandWithAttributes(t *testing.T) {

	Convey("Given I call the method buildInsertCommand", t, func() {

		context := manipulate.NewContext()
		context.Attributes = []string{"CITY"}
		context.PageSize = 0

		command, v := buildInsertCommand(context, "policy", []string{"NAME", "CITY", "DESCRIPTION"}, []interface{}{"Alexandre", "Sarralbe", "GOD"})
		So(command, ShouldEqual, `INSERT INTO policy (CITY) VALUES (?)`)
		So(v, ShouldResemble, []interface{}{"Sarralbe"})

	})
}

func TestMethodBuildInsertCommand(t *testing.T) {

	Convey("Given I call the method buildInsertCommand", t, func() {

		command, v := buildInsertCommand(nil, "policy", []string{"NAME", "CITY"}, []interface{}{"Alexandre", "Sarralbe"})
		So(command, ShouldEqual, `INSERT INTO policy (NAME, CITY) VALUES (?, ?)`)
		So(v, ShouldResemble, []interface{}{"Alexandre", "Sarralbe"})

	})
}

func TestMethodBuildGetCommand(t *testing.T) {

	Convey("Given I call the method buildGetCommand", t, func() {
		command, values := buildGetCommand(nil, "policy", []string{}, []interface{}{})
		So(command, ShouldEqual, `SELECT * FROM policy`)
		So(values, ShouldResemble, []interface{}{})
	})
}

func TestMethodBuildGetCommandWithAttributes(t *testing.T) {

	Convey("Given I call the method buildGetCommand", t, func() {

		context := manipulate.NewContext()
		context.Attributes = []string{"CITY", "DESCRIPTION"}
		context.PageSize = 0

		command, values := buildGetCommand(context, "policy", []string{}, []interface{}{})
		So(command, ShouldEqual, `SELECT CITY, DESCRIPTION FROM policy`)
		So(values, ShouldResemble, []interface{}{})
	})
}

func TestMethodBuildGetCommandWithAttributesAndPrimaryKeysAndValues(t *testing.T) {

	Convey("Given I call the method buildGetCommand", t, func() {

		context := manipulate.NewContext()
		context.Attributes = []string{"CITY", "DESCRIPTION"}
		context.PageSize = 0

		command, values := buildGetCommand(context, "policy", []string{"ID", "Name"}, []interface{}{"123", "Alexandre"})
		So(command, ShouldEqual, `SELECT CITY, DESCRIPTION FROM policy WHERE ID = ? AND Name = ?`)
		So(values, ShouldResemble, []interface{}{"123", "Alexandre"})
	})
}

func TestMethodBuildCountCommand(t *testing.T) {

	Convey("Given I call the method buildCountCommand", t, func() {
		command, values := buildCountCommand(nil, "policy")
		So(command, ShouldEqual, `SELECT COUNT(*) FROM policy`)
		So(values, ShouldResemble, []interface{}{})
	})
}

func TestMethodBuildIncrementCommand(t *testing.T) {

	Convey("Given I call the method buildIncrementCommand with two primary keys", t, func() {

		command, values := buildIncrementCommand(nil, "thecounter", "count", 2, []string{"id", "name"}, []interface{}{"12", "toto"})
		So(command, ShouldEqual, `UPDATE thecounter SET count = count + 2 WHERE id = ? AND name = ?`)
		So(values, ShouldResemble, []interface{}{"12", "toto"})
	})
}
