
.DEFAULT_GOAL := build

fmt:
	go fmt ./

vet: fmt
	go vet ./

build: vet
	go build

clean:
	go clean

.PHONY: fmt vet build clean