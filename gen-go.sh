#!/usr/bin/env bash

set -ex

protoc --proto_path=protos --go_out=. --go-grpc_out=. \
    ./protos/common.proto  ./protos/prometheus.proto \
    ./protos/meta_cluster.proto ./protos/meta_service.proto ./protos/meta_storage.proto ./protos/storage.proto
