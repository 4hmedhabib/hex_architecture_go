build_proto:
	protoc --proto_path=internal/adapters/framework/left/grpc/proto \
 		--go_out=internal/adapters/framework/left/grpc/proto \
 		--go-grpc_out=internal/adapters/framework/left/grpc/proto \
		internal/adapters/framework/left/grpc/proto/*.proto
