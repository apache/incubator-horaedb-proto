#!/usr/bin/env bash

set -ex

GO_PREFIX_PATH=github.com/CeresDB/ceresdbproto

protoc --proto_path=protos --go_out=. --go-grpc_out=. \
    ./protos/common.proto  ./protos/prometheus.proto \
    ./protos/meta/* ./protos/storage.proto

rm -rf pkg && mv $GO_PREFIX_PATH/pkg .
rm -rf github.com