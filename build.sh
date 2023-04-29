#!/bin/bash
# Build script for the project
# set -x if flag DEBUG is set
[ ! -z "$DEBUG" ] && set -x

# set pipefail to fail fast
set -o pipefail

echo "Building the project"
go build -o bin/ | tee bin/build.log
echo "Build completed successfully"
