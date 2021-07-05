#!/bin/sh
# proto生成命令，将pb.go文件生成到types/目录下, turingchain_path支持引用turingchain框架的proto文件
turingchain_path=$(go list -f '{{.Dir}}' "github.com/turingchain2020/turingchain")
protoc --go_out=plugins=grpc:../types ./*.proto --proto_path=. --proto_path="${turingchain_path}/types/proto/"
