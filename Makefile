LOCAL_BIN:=$(CURDIR)/app/bin

install-golangci-lint:
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.53.3

lint:
	GOBIN=$(LOCAL_BIN) golangci-lint run ./... --config .golangci.pipeline.yaml

install-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc

generate:
	make generate-chat-api

generate-chat-api:
	mkdir -p app/pkg/chat_v1
	protoc --proto_path app/api/chat_v1 \
	--go_out=app/pkg/chat_v1 --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=app/bin/protoc-gen-go \
	--go-grpc_out=app/pkg/chat_v1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=app/bin/protoc-gen-go-grpc \
	app/api/chat_v1/chat.proto