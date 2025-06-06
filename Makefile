LOCAL_BIN:=$(CURDIR)/bin

# install-golangci-lint:
# 	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.61.0

install-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc

lint:
	golangci-lint run ./... --config .golangci.pipeline.yaml
	# $(LOCAL_BIN)/golangci-lint run ./... --config .golangci.pipeline.yaml

generate:
	make generate-user-api

generate-user-api:
	mkdir -p pkg/user/v1
	protoc --proto_path=api/user/v1 \
		--go_out=pkg/user/v1 --go_opt=paths=source_relative \
		--plugin=protoc-gen-go=bin/protoc-gen-go \
		--go-grpc_out=pkg/user/v1 --go-grpc_opt=paths=source_relative \
		--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
		api/user/v1/user.proto

build:
	GOOS=linux GOARCH=amd64 go build -o service_linux cmd/grpc_server/main.go
