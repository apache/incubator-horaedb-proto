#!/usr/bin/env bash

# Licensed to the Apache Software Foundation (ASF) under one
# or more contributor license agreements.  See the NOTICE file
# distributed with this work for additional information
# regarding copyright ownership.  The ASF licenses this file
# to you under the Apache License, Version 2.0 (the
# "License"); you may not use this file except in compliance
# with the License.  You may obtain a copy of the License at
#
#   http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing,
# software distributed under the License is distributed on an
# "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
# KIND, either express or implied.  See the License for the
# specific language governing permissions and limitations
# under the License.

set -ex

ROOT_DIR=$( dirname "$0" )
GO_DIR="${ROOT_DIR}/golang"
GO_PREFIX_PATH="${GO_DIR}/github.com/apache/incubator-horaedb-proto/golang"

export PATH="$(go env GOPATH)/bin:$PATH"

protoc --proto_path=./protos --go_out="$GO_DIR" --go-grpc_out="$GO_DIR" \
    ./protos/cluster.proto \
    ./protos/common.proto \
    ./protos/meta_event.proto \
    ./protos/meta_service.proto \
    ./protos/meta_storage.proto \
    ./protos/prometheus.proto \
    ./protos/storage.proto

rm -rf "$GO_DIR/pkg" && mv "$GO_PREFIX_PATH/pkg" "$GO_DIR"
rm -rf "$GO_DIR/github.com"
