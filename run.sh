#!/bin/bash

make install_dependencies

echo "Installing mock"
pushd manipcassandra
make apomock
popd

echo "Lauching the tests"
make test

exit $?
