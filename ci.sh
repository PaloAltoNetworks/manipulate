#!/bin/bash

set -x

if [ ! -n "${BUILD_NUMBER+1}" ]; then
  BUILD_NUMBER="latest"
fi

## build the test container
make apodocker
docker build --file .dockerfile-test --build-arg GITHUB_TOKEN=${JENKINS_GITHUB_TOKEN} -t manipulate-test:${BUILD_NUMBER} .

## run the tests in the container
docker run --rm manipulate-test:${BUILD_NUMBER}
TEST_RESULT=$?

## cleanup
docker rmi manipulate-test:${BUILD_NUMBER}
make clean_apodocker

exit $TEST_RESULT
