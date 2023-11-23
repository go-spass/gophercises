.DEFAULT_GOAL := all

BUILD_DIR ?= ./build

fmt:
	go fmt ./...
.PHONY:fmt

lint: fmt
	golangci-lint run
#	golint ./...
.PHONY:lint

vet: lint
	go vet ./...
	shadow ./...
.PHONY: vet

test:
	go test -cover ./...
.PHONY: test

bench:
	go test -tags=benchmark -bench=. ./...
.PHONY: bench

hello: vet
	go build -o build/hello cmd/hello/main.go
.PHONY: hello

quiz: vet
	go build -o build/quiz cmd/quiz/main.go
.PHONY: quiz

urlshort: vet
	go build -o build/urlshort cmd/urlshort/main.go
.PHONY: urlshort

all: hello quiz urlshort
.PHONY: all

clean:
	rm -r $(BUILD_DIR)
.PHONY: clean

