gen: 
	@protoc \
	 --proto_path=protos "protos/orders.proto" \
	 --go_out=services/common/protoGen \
	 --go_opt=paths=source_relative
	 --go-grpc_out=services/common/protoGen \
	 --go-grpc_opt=paths=source_relative 