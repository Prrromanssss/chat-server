LOCAL_BIN:=$(CURDIR)/app/bin

lint:
	GOBIN=$(LOCAL_BIN) golangci-lint run ./... --config .golangci.pipeline.yaml

install-deps:
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.53.3
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@v3.14.0
	GOBIN=$(LOCAL_BIN) go install github.com/envoyproxy/protoc-gen-validate@v1.0.4


get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc

generate:
	make generate-chat-api

generate-chat-api:
	mkdir -p app/pkg/chat_v1
	protoc --proto_path app/api/chat_v1 \
	--proto_path app/vendor.protogen  \
	--go_out=app/pkg/chat_v1 \
	--go_opt=paths=source_relative \
	--plugin=protoc-gen-go=app/bin/protoc-gen-go \
	--go-grpc_out=app/pkg/chat_v1 \
	--go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=app/bin/protoc-gen-go-grpc \
	--validate_out lang=go:app/pkg/chat_v1 \
	--validate_opt=paths=source_relative \
	--plugin=protoc-gen-validate=app/bin/protoc-gen-validate \
	app/api/chat_v1/chat.proto

test-coverage:
	@cd app && \
	go clean -testcache && \
	go test ./... -coverprofile=coverage.tmp.out -covermode count -coverpkg=github.com/Prrromanssss/chat-server/internal/service/...,github.com/Prrromanssss/chat-server/internal/api/... -count 5  && \
	grep -v 'mocks\|config' coverage.tmp.out  > coverage.out  && \
	rm coverage.tmp.out
	cd app && go tool cover -html=coverage.out;
	cd app && go tool cover -func=./coverage.out | grep "total";
	grep -sqFx "/coverage.out" .gitignore || echo "/coverage.out" >> .gitignore

vendor-proto:
	@if [ ! -d app/vendor.protogen/validate ]; then \
		mkdir -p app/vendor.protogen/validate &&\
		git clone https://github.com/envoyproxy/protoc-gen-validate app/vendor.protogen/protoc-gen-validate &&\
		mv app/vendor.protogen/protoc-gen-validate/validate/*.proto app/vendor.protogen/validate &&\
		rm -rf app/vendor.protogen/protoc-gen-validate ;\
	fi