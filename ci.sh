#!/bin/bash

set -x

export TEST_CONTAINER_IMAGE_NAME="manipulate-test"
make create_test_container
make run_test_container
make clean_test_container

exit $?
