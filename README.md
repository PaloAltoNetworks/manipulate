# Manipulate

Package manipulate provides everything needed to perform CRUD operations
on an https://github.com/aporeto-inc/elemental based data model.

The main interface is Manipulator. This interface provides various
methods for creation, modification, retrieval and so on. TransactionalManipulator,
which is an extension of the Manipulator add methods to manage transactions, like
Commit and Abort.

A Manipulator works with a Manipulable which is a combination of the interfaces
elemental.Identifiable and elemental.Validatable.

The storage engine used by a Manipulator is abstracted. By default manipulate
provides implementations for Cassandra, Mongo, ReST API. You can of course implement
Your own storage implementation.

Each method of a Manipulator is taking a Context as argument. The context is used
to pass additional informations like a Filter, or some Parameters.

### Example for creating an object

```go
// Create a User from a generated Elemental model.
user := models.NewUser()
user.FullName, user.Login := "Antoine Mercadal", "primalmotion"

// Create Mongo Manipulator.
m := manipmongo.NewMongoManipulator("127.0.0.1", "test")

// Then create the User.
m.Create(nil, nil, user)
```

### Example for retreving an object

```go
// Create a Context with a filter.
ctx := manipulate.NewContextWithFilter(manipulate.NewFilterComposer().WithKey("login").Equals("primalmotion"))

// Retrieve the users matching the filter.
var users models.UserLists
m.RetrieveMany(ctx, nil, models.UserIdentity, &users)
```
