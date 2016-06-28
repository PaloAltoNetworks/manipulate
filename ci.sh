#!/bin/bash

set -x

if [ ! -n "${BUILD_NUMBER+1}" ]; then
  BUILD_NUMBER="latest"
fi

docker build --build-arg GITHUB_TOKEN=${JENKINS_GITHUB_TOKEN} -t manipulate-test:${BUILD_NUMBER} .

docker run --rm manipulate-test:${BUILD_NUMBER}
TEST_RESULT=$?

docker rmi manipulate-test:${BUILD_NUMBER}

exit $TEST_RESULT
