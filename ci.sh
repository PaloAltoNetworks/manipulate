#!/bin/bash

set -x

export BUILD_CONTAINER_IMAGE_NAME="manipulate-build"
make create_test_container
make run_test_container
make clean_test_container

exit $?
