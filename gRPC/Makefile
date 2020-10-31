install_procbuf:
	go get google.golang.org/protobuf/cmd/protoc-gen-go
	go get google.golang.org/grpc/cmd/protoc-gen-go-grpc

generate_pb:
	protoc --go_out=plugins=grpc:chat chat.proto