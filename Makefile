generate_proto:
	protoc --proto_path=proto/v1 --go_out=paths=source_relative:pkg/proto --go-grpc_out=require_unimplemented_servers=false,paths=source_relative:pkg/proto proto/v1/service.proto

generate_docs:
	protoc --doc_out=doc --doc_opt=html,index.html proto/v1/service.proto