#!/bin/sh

turingchain_path=$(go list -f '{{.Dir}}' "github.com/turingchain2020/turingchain")
protoc --go_out=plugins=grpc:../types ./*.proto --proto_path=. --proto_path="${turingchain_path}/types/proto/"
