MAIN_PACKAGE_PATH := ./cmd/authGrpc
BINARY_NAME := authService

.PHONY: build
build:
	go build -o=/tmp/bin/${BINARY_NAME} ${MAIN_PACKAGE_PATH}

.PHONY: run
run: build
	/tmp/bin/${BINARY_NAME}

.PROXY: generate
generate:
	mkdir -p internal/auth_v1
	protoc -I/usr/include -I. --go_out=internal/auth_v1 --go_opt=paths=import \
		--go-grpc_out=internal/auth_v1 --go-grpc_opt=paths=import \
		api/authGrpc/auth.proto
	mv internal/auth_v1/github.com/2pizzzza/authGrpc/internal/auth_v1/* internal/auth_v1/
	rm -rf internal/auth_v1/github.com
