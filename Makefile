SHELL := /bin/bash # Use bash syntax
.PHONY: install deploy build login push clean tool lint help
.DEFAULT: help

ENV_LIST=$(shell basename -a -s .env env/*.env)
REGISTRY=dev-registry.aralego.com
BUILD_NAME=backend/apigateway
BUILD_TAG=latest

help:
	@echo "make install: compile packages and dependencies"
	@echo "make deploy: build & push docker image"
	@echo "make build: build docker image"
	@echo "make login: login to docker registry"
	@echo "make push: push image to docker registry"
	@echo "make tool: run specified go tool"
	@echo "make lint: golint ./..."
	@echo "make clean: remove object files and cached files"
	@echo "make test: run all unit tests"

install:
ifeq (, $(wildcard $(shell which swag)))
	@go get -u github.com/swaggo/swag/cmd/swag
endif
	swag init
	@go build -v .
ifneq ($(findstring $(ENV),$(ENV_LIST)),)
	cp ./env/$(ENV).env .env
endif

deploy: build login push

build:
	docker build --platform linux/amd64 -t "$(REGISTRY)/$(BUILD_NAME):$(BUILD_TAG)" .

login:
	docker login -u admin -p admin $(REGISTRY)

push:
	docker push $(REGISTRY)/$(BUILD_NAME):$(BUILD_TAG)

tool:
	go vet -composites=false ./...; true
	gofmt -w .

lint:
	docker run --rm -v $(shell pwd):/app -w /app golangci/golangci-lint:v1.42.0 golangci-lint run -v --timeout 300s

clean:
	rm -f APIGateway
	rm -rf docs
	rm -f .env
	go clean -i .

test: install
	go test -v -benchmem -bench=. -run=none ./...
	go test -v -covermode=count -coverprofile=test.out ./...
	go tool cover -html=test.out