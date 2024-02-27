#!/bin/bash

set -e

ENV_FILE=".env"
BINARY_NAME="app"

if [ ! -f "$ENV_FILE" ];
then
    echo "Missing .env file at $(pwd)/$ENV_FILE"
    # do not exit 1, can run without .env file
else
  export $(grep -v '^#' $ENV_FILE | xargs)
fi

echo "Building the Go application..."
go build -o $BINARY_NAME main.go

echo "Running the application..."
./$BINARY_NAME