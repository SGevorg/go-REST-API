#!/bin/bash

# Usage!!
# cd /path/to/repo/go-rest-api
# ./ops/build.sh <version>

# Build the docker container with the given version

VERSION=${1:-latest}

docker build -t go-rest-api:${VERSION} .