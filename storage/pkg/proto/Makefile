LOCAL_BIN:=$(CURDIR)/bin

install-deps:
	mkdir -p bin
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc

generate-storage:
	protoc --proto_path storage --proto_path vendor.protogen \
		--go_out=storage --go_opt=paths=source_relative \
		--plugin=protoc-gen-go=bin/protoc-gen-go \
		--go-grpc_out=storage --go-grpc_opt=paths=source_relative \
		--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
		storage/storage.proto

vendor-proto:
	@if [ ! -d vendor.protogen/google ]; then \
		git clone https://github.com/googleapis/googleapis vendor.protogen/googleapis && \
		mkdir -p vendor.protogen/google/ &&\
		mv vendor.protogen/googleapis/google/api vendor.protogen/google && \
		rm -rf vendor.protogen/googleapis ;\
	fi