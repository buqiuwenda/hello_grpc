#!/usr/bin/env bash
protoc --proto_path=${GOPATH}/src:. --go_out=plugins=grpc:. api/protobuf-spec/example/example.proto
