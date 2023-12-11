//go:generate protoc --proto_path=api/v1 --go_out=paths=source_relative:pkg/proto --go-grpc_out=paths=source_relative:pkg/proto api/v1/service.proto

package main
