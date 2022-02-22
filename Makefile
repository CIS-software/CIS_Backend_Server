.PHONY: build
build:
	go build -v ./cmd/main/main

.DEFAULT_GOAL := build
