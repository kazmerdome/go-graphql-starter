NAME=go-graphql-starter
VERSION=1.0.0

.PHONY: gateway
gateway:
	go run cmd/gateway/main.go

.PHONY: generate
generate:
	cd pkg/gateway && go run ../../scripts/gqlgen.go

.PHONY: build
build:
	@go build -o build/$(NAME) cmd/gateway/main.go

.PHONY: run
run: build
	@./build/$(NAME)
