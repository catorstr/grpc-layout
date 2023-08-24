GOPATH:=$(shell go env GOPATH)
APP_NAME="app-grpc-layout"
# Release variables
# ------------------
GIT_COMMIT?=$(shell git rev-parse "HEAD^{commit}" 2>/dev/null)
GIT_TAG?=$(shell git describe --abbrev=0 --tags 2>/dev/null)
BUILD_DATE:=$(shell date -u +'%Y-%m-%dT%H:%M:%SZ')

LDFLAGS:=-X 'main.gitVersion=$(GIT_TAG)' -X 'main.gitCommit=$(GIT_COMMIT)' -X 'main.buildDate=$(BUILD_DATE)'


.PHONY: init
init:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest


.PHONY: proto
# protobuf协议转换为所需的源代码
# 映射gw 默认提供的http服务
# 生成反向代理并使protoc-gen-openapiv2支持的自定义protobuf注释
proto:
	protoc -I . --go-grpc_out=../ --go_out=../ api/*/*.proto
	protoc -I . --grpc-gateway_out=allow_delete_body=true:. --grpc-gateway_opt logtostderr=true --grpc-gateway_opt paths=source_relative --grpc-gateway_opt generate_unbound_methods=true  api/*/*.proto
	protoc -I . --openapiv2_out . --openapiv2_opt logtostderr=true --openapiv2_opt allow_delete_body=true  api/*/*.proto

.PHONY: build
build: proto test
	go build -ldflags "$(LDFLAGS)" -o ${APP_NAME} .

.PHONY: test
test:
	go test -v ./... -cover

docker-build:
	go build -ldflags "$(LDFLAGS)" -o ${APP_NAME} .

# docker:
# 	docker build ../ -t "${REGISTRY}/${APP_NAME}:${GIT_COMMIT}" -f ../grpc-layout/Dockerfile
# 	docker push "${REGISTRY}/${APP_NAME}:${GIT_COMMIT}"
