#! /usr/bin/sh

init_file=

if stat ./init_vars.sh &> /dev/null; then
	echo fuck
	init_file='./init_vars.sh';
elif stat ./scripts/init_vars.sh &> /dev/null; then
	init_file='./scripts/init_vars.sh';
else
	echo 'cannot find init_vars.sh file';
	exit 1
fi;

source "$init_file"

# Fail on any error / non-zero exit code / use of uninitialied variable
set -euo pipefail

OUT_DIR="$SERVICES_DIR/fastrunner"

mkdir -p $OUT_DIR

echo "using:"
echo "    output = $OUT_DIR"
echo "    protodir = $PROTO_DIR"

protoc --proto_path=$PROTO_DIR \
	--go_out=$OUT_DIR \
	--go-grpc_out=$OUT_DIR \
	$PROTO_DIR/gorunner.proto

protoc --proto_path=$PROTO_DIR \
	--go_out=$OUT_DIR \
	--go-grpc_out=$OUT_DIR \
	$PROTO_DIR/jsrunner.proto
