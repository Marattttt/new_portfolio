#! /usr/bin/sh

# Fail on any error / non-zero exit code 
set -eo pipefail

USAGE='This script generates protobuf bindings for the service 

Usage: genproto.sh <path to proto dir> <path to service dir>

Requires "@grpc/grpc-js and @grpc/grpc-js-ts" and "protoc-gen-grpc" for generation'


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

# generate js codes with @grpc/grpc-js
protoc-gen-grpc \
  --js_out=import_style=commonjs,binary:${OUT_DIR} \
  --grpc_out=grpc_js:$OUT_DIR \
  --proto_path $PROTO_DIR \
  $PROTO_DIR/gorunner.proto

# generate d.ts codes with @grpc/grpc-js
protoc-gen-grpc-ts \
  --ts_out=grpc_js:${OUT_DIR} \
  --proto_path $PROTO_DIR \
  $PROTO_DIR/jsrunner.proto
