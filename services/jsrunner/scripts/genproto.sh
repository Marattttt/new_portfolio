#! /usr/bin/sh

# Fail on any error / non-zero exit code 
set -eo pipefail

USAGE='This script generates protobuf bindings for the service 

Usage: genproto.sh <path to proto dir> <path to service dir>

Requires "protoc-gen-go" and "protoc" for generation'


if [ $1 ]; then 
  PROTO_DIR=$(realpath "$1")
  echo "Using proto directory = $PROTO_DIR"
else
  echo "$USAGE"
  exit 1
fi

if [ $2 ]; then 
  SERVICE_DIR=$(realpath "$2")
  echo "Using service directory = $SERVICE_DIR"
else
  echo "$USAGE"
  exit 1
fi

# Add fail on use of uninitialied variable
set -u pipefail

OUT_DIR="$SERVICE_DIR/src/protogen"

protoc --proto_path=$PROTO_DIR \
	--go_out=$OUT_DIR \
	--go-grpc_out=$OUT_DIR \
	$PROTO_DIR/jsrunner.proto
