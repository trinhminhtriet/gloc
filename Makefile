VERSION := $(shell git describe --tags)
COMMIT := $(shell git rev-parse --short HEAD)

.PHONY: test build

build:
	mkdir -p bin
	GO111MODULE=on go build -ldflags="-s -w -X main.Version=${VERSION} -X main.GitCommit=${COMMIT}" -o ./bin/gloc cmd/gloc/main.go

update-package:
	GO111MODULE=on go get -u github.com/trinhminhtriet/gloc

cleanup-package:
	GO111MODULE=on go mod tidy

run-example:
	GO111MODULE=on go run examples/languages/main.go
	GO111MODULE=on go run examples/files/main.go

test:
	GO111MODULE=on go test -v

test-cover:
	GO111MODULE=on go test -v -coverprofile=coverage.out
