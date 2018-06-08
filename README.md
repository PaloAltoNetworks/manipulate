# Manipulate

[![codecov](https://codecov.io/gh/aporeto-inc/manipulate/branch/master/graph/badge.svg?token=2dEWoQKRO0)](https://codecov.io/gh/aporeto-inc/manipulate)

Package manipulate provides everything needed to perform CRUD operations
on an [elemental](https://github.com/aporeto-inc/elemental) based data model.

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
m := manipmongo.NewMongoManipulator("127.0.0.1", "test")

// Then create the User.
m.Create(nil, user)
```

## Example for retreving an object

```go
// Create a Context with a filter.
ctx := manipulate.NewContextWithFilter(manipulate.NewFilterComposer().
    WithKey("login").
    Equals("primalmotion").
    Done(),
)

// Retrieve the users matching the filter.
var users models.UserLists
m.RetrieveMany(ctx, models.UserIdentity, &users)
```

## Example to retrieve an Aporeto Processing Unit

> Note: this is a specific Aporeto use case.

```go
package main

import (
    "flag"
    "fmt"
    "os"

    "go.aporeto.io/gaia/squallmodels/v1/golang"
    "go.aporeto.io/manipulate"
    "go.aporeto.io/manipulate/manipwebsocket"
)

const aporetoAPIURL = "https://squall.console.aporeto.com"

func main() {

    // Here we get the cli parameters. Nothing exciting.
    var token, namespace string
    flag.StringVar(&token, "token", "", "A valid Aporeto token")
    flag.StringVar(&namespace, "namespace", "", "Your namespace")
    flag.Parse()
    if token == "" || namespace == "" {
        fmt.Println("Please pass both the -token and -namespace paramaters")
        os.Exit(1)
    }

    // We want to recursively retrieve all the running processing units starting
    // from the given namespace.

    // We first create a simple manipulator using an existing token. There are more
    // sophisticated ways of doing this, such as using
    // manipwebsocket.NewWebSocketManipulatorWithMidgardCertAuthentication to handle
    // automatic token renewal based on a certificate, but we'll keep this example simple.
    manipulator, disconnect, err := manipwebsocket.NewWebSocketManipulator(
        "Bearer",
        token,
        aporetoAPIURL,
        namespace,
    )
    if err != nil {
        panic(err)
    }

    // As we want only the Running processing unit, we need to filter them on
    // the operationalstatus status tag.
    // To do so, we need to create a manipulate.Context and give it a filter.
    mctx := manipulate.NewContextWithFilter(
        manipulate.NewFilterComposer().
            WithKey("operationalstatus").Equals("Running").
            Done(),
    )

    // Then as we want to get all processing units recursively, we set
    // the Recursive parameter of the context to true.
    mctx.Recursive = true

    // We create a ProcessingUnitsList to store the results.
    var pus squallmodels.ProcessingUnitsList

    // We call the manipulator.RetrieveMany using our context to retrieve the data.
    if err := manipulator.RetrieveMany(mctx, &pus); err != nil {
        panic(err)
    }

    // We print the results.
    for _, pu := range pus {
        fmt.Println("namespace:", pu.Namespace, "name:", pu.Name)
    }

    // And we nicely disconnect.
    disconnect()
}
```

Build this, retrieve a valid token (for instance using `apoctl`), then execute it:

    ./main -token $TOKEN -namespace /your-namespace
