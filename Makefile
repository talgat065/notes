PROJECT_PKG = github.com/talgat065/notes

BUILD_DIR = bin/cli

.DEFAULT_GOAL=all

.PHONY: test

all:
	make setup
	make build
	make test
	make lint

setup:
	go mod verify
	go mod tidy
	go mod vendor

build:
	go build -o ${BUILD_DIR} ./cmd/cli

serve:
	${BUILD_DIR}

lint:
	golangci-lint run --fix

test:
	go test -v -race ./...
