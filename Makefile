SHELL := /bin/bash

BINARY_NAME=loop

export GO111MODULE=on

.PHONY: check format vet lint build tidy

check: format vet lint

format:
	@echo "Formatting packages using gofmt..."
	@gofmt -w .
	@echo "Done"

generate:
	@echo "Generate grpc code..."
	@protoc --go_out=plugins=grpc:. service/grpc/loopd.proto
	@echo "Done"

vet:
	@echo "Checking packages using go vet"
	@go vet ./...
	@echo "Done"

lint:
	@echo "Checking packages using golint"
	@golint ./...
	@echo "Done"

build: generate tidy format
	@mkdir -p ./bin
	@echo "Building loop && loopd..."
	@go build -o ./bin/loop ./cmd/loop
	@go build -o ./bin/loopd ./cmd/loopd
	@echo "Done"

install: build
	@echo "Installing loop && loopd..."
	@cp ./bin/loop* ${GOPATH}/bin
	@echo "Done"

tidy:
	@echo "Tidy and check the go mod files"
	@go mod tidy
	@go mod verify
	@echo "Done"