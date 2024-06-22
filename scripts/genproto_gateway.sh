#! /usr/bin/sh


init_file=
if stat ./init_vars.sh &> /dev/null; then
	echo fuck
	init_file='./init_vars.sh';
elif stat ./scripts/init_vars.sh &> /dev/null; then
	init_file='./scripts/init_vars.sh';
else
	echo 'cannot find scripts/init_vars.sh file';
	exit 1
fi;

source "$init_file"

mkdir -p $OUT_DIR

set -euo pipefail

# generate js codes with @grpc/grpc-js
protoc-gen-grpc \
  --js_out=import_style=commonjs,binary:${OUT_DIR} \
  --grpc_out=grpc_js:$OUT_DIR \
  --proto_path $PROTO_DIR \
  $PROTO_DIR/fastrunner.proto

# generate d.ts codes with @grpc/grpc-js
protoc-gen-grpc-ts \
  --ts_out=grpc_js:${OUT_DIR} \
  --proto_path $PROTO_DIR \
  $PROTO_DIR/fastrunner.proto
