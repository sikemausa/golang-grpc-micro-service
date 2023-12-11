//go:generate protoc --proto_path=pb/v1 --go_out=paths=source_relative:pkg/proto --go-grpc_out=require_unimplemented_servers=false,paths=source_relative:pkg/proto pb/v1/service.proto

package main
