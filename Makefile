.PHONY build:
build:
	go build -v ./cmd/gotesserver

.PHONY: test
test:
	go test -v -race -timeout 30s ./internal/...

.DEFAULT_GOAL := build
