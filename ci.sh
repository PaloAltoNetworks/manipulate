#!/bin/bash

set -x

docker build --build-arg GITHUB_TOKEN=${GITHUB_TOKEN} -t manipulabletest .
docker run -t manipulabletest

exit $?
