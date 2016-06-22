## Manipulate

Manipulate is a library which allows you to manipulate Manipulate with a set of store.
The current store available are a http store and a cassandra store.

They come with the the following interface :

```go
// Manipulator is the interface of a storage backend.
type Manipulator interface {
	RetrieveChildren(contexts Contexts, parent Manipulable, identity elemental.Identity, dest interface{}) elemental.Errors
	Retrieve(contexts Contexts, objects ...Manipulable) elemental.Errors
	Create(contexts Contexts, parent Manipulable, objects ...Manipulable) elemental.Errors
	Update(contexts Contexts, objects ...Manipulable) elemental.Errors
	Delete(contexts Contexts, objects ...Manipulable) elemental.Errors
	Count(contexts Contexts, identity elemental.Identity) (int, elemental.Errors)
	Assign(contexts Contexts, parent Manipulable, assignation *elemental.Assignation) elemental.Errors
}
```

## Launch the test

First install the testing dependencies with the command

```bash
make install_dependencies
pushd manipcassandra; make apomock; popd
```

Then simply do :

```bash
make test
```

or

```bash
make convey
```
