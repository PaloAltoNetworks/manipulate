package testdata

import (
	"fmt"

	"github.com/aporeto-inc/elemental"
)

// ListIdentity represents the Identity of the object
var ListIdentity = elemental.Identity{
	Name:     "list",
	Category: "lists",
}

// ListsList represents a list of Lists
type ListsList []*List

// ContentIdentity is test data
func (o ListsList) ContentIdentity() elemental.Identity {
	return ListIdentity
}

// List is test data
func (o ListsList) List() elemental.IdentifiablesList {
	return nil
}

// Version is test data
func (o ListsList) Version() int {
	return 1
}

// List represents the model of a list
type List struct {
	ID          string `json:"ID,omitempty"`
	ParentID    string `json:"parentID,omitempty"`
	ParentType  string `json:"parentType,omitempty"`
	Owner       string `json:"owner,omitempty"`
	Description string `json:"description,omitempty"`
	Name        string `json:"name,omitempty"`

	ModelVersion int `json:"-"`
}

// NewList returns a new *List
func NewList() *List {

	return &List{ModelVersion: 1}
}

// Identity returns the Identity of the object.
func (o *List) Identity() elemental.Identity {

	return ListIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *List) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *List) SetIdentifier(ID string) {

	o.ID = ID
}

// Version is test data
func (o *List) Version() int {

	return o.ModelVersion
}

// Validate valides the current information stored into the structure.
func (o *List) Validate() error {

	errors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

// TaskIdentity represents the Identity of the object
var TaskIdentity = elemental.Identity{
	Name:     "task",
	Category: "tasks",
}

// TasksList represents a list of Tasks
type TasksList []*Task

// ContentIdentity is test data
func (o TasksList) ContentIdentity() elemental.Identity {
	return TaskIdentity
}

// List is test data
func (o TasksList) List() elemental.IdentifiablesList {
	return nil
}

// Version is test data
func (o TasksList) Version() int {
	return 1
}

// Task represents the model of a task
type Task struct {
	ID          string `json:"ID,omitempty"`
	ParentID    string `json:"parentID,omitempty"`
	ParentType  string `json:"parentType,omitempty"`
	Owner       string `json:"owner,omitempty"`
	Description string `json:"description,omitempty"`
	Name        string `json:"name,omitempty"`
	Status      string `json:"status,omitempty"`

	ModelVersion int `json:"-"`
}

// NewTask returns a new *Task
func NewTask() *Task {

	return &Task{
		ModelVersion: 1,
		Status:       "TODO",
	}
}

// Identity returns the Identity of the object.
func (o *Task) Identity() elemental.Identity {

	return TaskIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Task) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Task) SetIdentifier(ID string) {

	o.ID = ID
}

// Version is test data
func (o *Task) Version() int {

	return o.ModelVersion
}

// Validate valides the current information stored into the structure.
func (o *Task) Validate() error {

	errors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateStringInList("status", o.Status, []string{"DONE", "PROGRESS", "TODO"}, false); err != nil {
		errors = append(errors, err)
	}

	return errors
}

// UnmarshalableListIdentity is test data
var UnmarshalableListIdentity = elemental.Identity{Name: "list", Category: "lists"}

// UnmarshalableList is test data
type UnmarshalableList struct {
	List
}

// NewUnmarshalableList is test data
func NewUnmarshalableList() *UnmarshalableList {
	return &UnmarshalableList{List: List{}}
}

// Identity is test data
func (o *UnmarshalableList) Identity() elemental.Identity { return UnmarshalableListIdentity }

// UnmarshalJSON is test data
func (o *UnmarshalableList) UnmarshalJSON([]byte) error {
	return fmt.Errorf("error unmarshalling")
}

// MarshalJSON is test data
func (o *UnmarshalableList) MarshalJSON() ([]byte, error) {
	return nil, fmt.Errorf("error marshalling")
}

// Validate is test data
func (o *UnmarshalableList) Validate() error { return nil }
