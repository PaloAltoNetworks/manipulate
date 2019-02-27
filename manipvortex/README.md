# Vortex - Local Cache for Elemental Objects

Vortex is a generic elemental cache, currently backed by MemDB. Other implementations
are possible in the future.

Vortex requires three components:

* A manipulator that allows Vortex to synchronize objects in the cache with a remote backend.

* An optional subscriber that allows vortex to synchronize events from an upstream susbscriber.

* A local datastore where it will cache objects. The datastore must be properly configured with any indexes as required by your queries.

The downstream manipulator is optional and vortex can also be used a standalone embedded
database that obays the manipulate interfaces. Note, that vortex supports both the
manipulator and subscriber interfaces and it will propagate any events from the
downstream to upstream callers.

For every object you must provide a processor configuration that will define any hook configuration. Separate hooks are possible before its transaction:

* RemoteHook will be applied before the transaction is send upstream.

* LocalHost will be applied before the transaction is commited locally.

* RetrieveManyHook is applied before retrieve many operations.

For every identity, you can configure the cache mode:

* Write-through indicates that the object will be first commited downstream and only if it succeeds it will be committed localy.

*Write-back indicates that the object will be stored in a local queue and later it will be send to the downstream database.

You can take a look at the test files for examples.

Here is a simple example:

```go
// NewMemoryDB will create the DB.
func NewMemoryDB(
    ctx context.Context,
    b manipulate.TokenManager,
    api string,
    model elemental.ModelManager,
) (manipulate.Manipulator, context.CancelFunc, error) {

    connectContext, cancel := context.WithCancel(ctx)

    m, err := maniphttp.New(
        connectContext,
        api,
        maniphttp.OptionNamespace(namespace),
        maniphttp.OptionTokenManager(b),
    )

    if err != nil {
        cancel()
        return nil, nil, err
    }

    subscriber := maniphttp.NewSubscriber(m, true)
    subscriber.Start(ctx, nil)

    indexConfig := map[string]*config.MemDBIdentity{
        testmodel.ListIdentity.Category: &config.MemDBIdentity{
            Identity: testmodel.ListIdentity,
            Indexes: []*config.IndexConfig{
                &config.IndexConfig{
                    Name:      "id",
                    Type:      config.String,
                    Unique:    true,
                    Attribute: "ID",
                },
                &config.IndexConfig{
                    Name:      "Name",
                    Type:      config.String,
                    Unique:    false,
                    Attribute: "Name",
                ,
                &config.IndexConfig{
                    Name:      "Slice",
                    Type:      config.Slice,
                    Unique:    false,
                    Attribute: "Slice",
                },
            },
        },
    }

// Create a data store, register the identities and start it.
    datastore, err := memdbvortex.NewDatastore(indexConfig)
    if err != nil {
        return nil, cancel, fmt.Errorf("failed to create local memory db: %s", err)
    }

// Create the processors and the vortex.
    processors := map[string]*config.ProcessorConfiguration{
        testmodel.ListIdentity.Name: &config.ProcessorConfiguration{
            Identity:         testmodel.ListIdentity,
            CommitOnEvent:    true,
        },

    v, err := memdbvortex.NewMemDBVortex(
        ctx,
        datastore,
        processors,
        gaia.Manager(),
        memdbvortex.OptionBackendManipulator(s),
        memdbvortex.OptionBackendSubscriber(subscriber),
        )

    return v, cancel, err
}

```
