#!/usr/bin/env bash

set -ex

GO_PREFIX_PATH=github.com/CeresDB/ceresdbproto/golang

protoc --proto_path=../protos --go_out=. --go-grpc_out=. ../protos/*

rm -rf pkg && mv $GO_PREFIX_PATH/pkg .
rm -rf github.com
