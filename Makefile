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

day2: vet
	go build -o build/day2 cmd/day2/main.go
.PHONY: day2

day3: vet
	go build -o build/day3 cmd/day3/main.go
.PHONY: day3

day4: vet
	go build -o build/day4 cmd/day4/main.go
.PHONY: day4

all: hello
.PHONY: all

clean:
	rm -r $(BUILD_DIR)
.PHONY: clean

