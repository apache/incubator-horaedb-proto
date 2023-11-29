#!/usr/bin/env bash

# Copyright 2023 The HoraeDB Authors
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.


set -ex

GO_PREFIX_PATH=github.com/CeresDB/horaedbproto/golang

protoc --proto_path=../protos --go_out=. --go-grpc_out=. ../protos/cluster.proto ../protos/common.proto ../protos/meta_event.proto ../protos/meta_service.proto ../protos/meta_storage.proto ../protos/prometheus.proto ../protos/storage.proto

rm -rf pkg && mv $GO_PREFIX_PATH/pkg .
rm -rf github.com
