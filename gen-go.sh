#!/usr/bin/env bash

set -ex

protoc --proto_path=protos --go_out=. --go-grpc_out=. \
    ./protos/common.proto  ./protos/storage.proto  ./protos/prometheus.proto
