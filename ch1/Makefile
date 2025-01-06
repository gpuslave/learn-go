
.DEFAULT_GOAL := build

fmt:
	go fmt ./

vet: fmt
	go vet ./

build: vet
	go build -o test1

clean:
	go clean

.PHONY: fmt vet build clean