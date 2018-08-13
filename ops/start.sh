#!/bin/bash

# Usage!!
# cd /path/to/repo/go-rest-api
# ./ops/start.sh <version>

# Build the docker container with the given version

VERSION=${1:-latest}

docker run -d --rm --name go-rest-api -p 8000:8000 go-rest-api:$VERSION