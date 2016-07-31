#!/bin/bash

# Copyright 2014 Google Inc. All Rights Reserved.
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


SRC=$GOPATH/src
PKGMAP=Mgoogle/protobuf/descriptor.proto=github.com/golang/protobuf/protoc-gen-go/descriptor

rm -f $SRC/google/emulators/broker.*
rm -f $SRC/google/api/*
rm -f $SRC/google/protobuf/*

echo "GO: broker protos"
protoc -I googleapis -I protos \
  protos/google/emulators/broker.proto \
  --go_out=plugins=grpc:$SRC \
  --grpc-gateway_out=logtostderr=true:$SRC

echo "GO: google/api"
protoc -I googleapis \
  googleapis/google/api/*.proto \
  --go_out=$PKGMAP,plugins=grpc:$SRC

echo "GO: google/protobuf"
protoc -I googleapis \
  googleapis/google/protobuf/*.proto \
  --go_out=plugins=grpc:$SRC
