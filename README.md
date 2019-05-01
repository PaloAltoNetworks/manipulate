# Manipulate

[![codecov](https://codecov.io/gh/aporeto-inc/manipulate/branch/master/graph/badge.svg?token=2dEWoQKRO0)](https://codecov.io/gh/aporeto-inc/manipulate)

> README IS A WORK IN PROGRESS AS WE ARE WRITTING MORE DOCUMENTATION ABOUT THIS PACKAGE.

Package manipulate provides everything needed to perform CRUD operations
on an [elemental](https://go.aporeto.io/elemental) based data model.

The main interface is `Manipulator`. This interface provides various
methods for creation, modification, retrieval and so on.

A Manipulator works with `elemental.Identifiable`.

The storage engine used by a Manipulator is abstracted. By default manipulate
provides implementations for Mongo, ReST HTTP, Websocket and a Memory backed datastore. You can of course implement
Your own storage implementation.

Each method of a Manipulator is taking a `manipulate.Context` as argument. The context is used
to pass additional informations like a Filter, or some Parameters.

## Example for creating an object

```go
// Create a User from a generated Elemental model.
user := models.NewUser() // always use the initializer to get various default value correctly set.
user.FullName := "Antoine Mercadal"
user.Login := "primalmotion"

// Create Mongo Manipulator.
m := manipmongo.New("127.0.0.1", "test")

// Then create the User.
m.Create(nil, user)
```

## Example for retreving an object

```go
// Create a Context with a filter.
ctx := manipulate.NewContextWithFilter(manipulate.NewFilterComposer().
    WithKey("login").Equals("primalmotion").
    Done(),
)

// Retrieve the users matching the filter.
var users models.UserLists
m.RetrieveMany(ctx, &users)
```
