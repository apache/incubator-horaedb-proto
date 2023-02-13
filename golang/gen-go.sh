#!/usr/bin/env bash

set -ex

GO_PREFIX_PATH=github.com/CeresDB/ceresdbproto/golang

protoc --proto_path=../protos --go_out=. --go-grpc_out=. ../protos/cluster.proto ../protos/common.proto ../protos/meta_event.proto ../protos/meta_service.proto ../protos/meta_storage.proto ../protos/prometheus.proto ../protos/storage.proto

rm -rf pkg && mv $GO_PREFIX_PATH/pkg .
rm -rf github.com
