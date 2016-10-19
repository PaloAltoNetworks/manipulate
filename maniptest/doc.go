// Package maniptest contains a Mockable TransactionalManipulator.
// It implements all method of the TransactionalManipulator but do nothing.
//
// Methods can be mocked by using one of the MockXX method.
//
// For example:
//      m := maniptest.NewTestManipulator()
//      m.MockCreate(t, func(context *manipulate.Context, objects ...manipulate.Manipulable) error {
//          return elemental.NewError("title", "description", "subject", 43)
//      })
//
// The next calls to the Create method will use the given method, in the context of the given *testing.T.
// If you need to reset the mocked method in the context of the same test, simply do:
//
//      m.MockCreate(t, nil)
//
package maniptest
