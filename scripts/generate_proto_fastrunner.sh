#! /usr/bin/sh

# protoc --proto_path=proto --go_out=gen --go_opt=paths=source_relative --go-grpc_out=gen --go-grpc_opt=paths=source_relative proto/example.proto
protoc --proto_path=proto \
	--go_out=services/fastrunner/grpc/grpcgen \
	--go_opt=paths=source_relative \
	--go-grpc_out=services/fastrunner/grpc/grpcgen \
	--go-grpc_opt=paths=source_relative \
	proto/fastrunner.proto
