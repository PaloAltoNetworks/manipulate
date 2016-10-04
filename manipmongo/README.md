# manipmongo

manipmongo is the Manipulate driver to interface an elemental model to a Mongo database.


## Example

Here is a simple example on how to use it.

```go
package main

import (
	"fmt"

	"github.com/aporeto-inc/manipulate"
	"github.com/aporeto-inc/manipulate/manipmongo"

	gaia "github.com/aporeto-inc/gaia/golang"
)

func main() {

	// create a new Store and start it.
	store := manipmongo.NewMongoStore("127.0.0.1", "testdb")
	store.Start()

	// create an empty context.
	mctx := manipulate.NewContext()

	// Create one FilePath
	filepath1 := gaia.NewFilePath()
	filepath1.Name = "My File Path 1"
	filepath1.Filepath = "/etc/passwd"
	filepath1.Server = "localhost"
	store.Create(mctx, nil, filepath1)

	// Create a second FilePath
	filepath2 := gaia.NewFilePath()
	filepath2.Name = "My File Path 2"
	filepath2.Filepath = "/opt/*"
	filepath2.Server = "localhost"
	store.Create(mctx, nil, filepath2)

	// Update File Path 2's Name
	filepath2.Name = "New Name"
	store.Update(mctx, filepath2)

	// Retrieve file path 1 in a separate object.
	rfp1 := gaia.NewFilePath()
	rfp1.SetIdentifier(filepath1.Identifier())
	store.Retrieve(mctx, rfp1)
	fmt.Println("Retrieved filepath 1 (ID, Name): ", rfp1.ID, rfp1.Name)

	// Retrieve file path 2 in a separate object.
	rfp2 := gaia.NewFilePath()
	rfp2.SetIdentifier(filepath2.Identifier())
	store.Retrieve(mctx, rfp2)
	fmt.Println("Retrieved filepath 2 (ID, Name): ", rfp2.ID, rfp2.Name)

	// Delete file path 1 and 2
	store.Delete(mctx, rfp1)
	store.Delete(mctx, rfp2)

	// Stop the store.
	store.Stop()
}
```
