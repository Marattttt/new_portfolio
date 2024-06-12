#! /usr/bin/sh


PROTO_DIR="../proto"
OUT_DIR="../services/gateway/src/protogen"

mkdir -p $OUT_DIR

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
