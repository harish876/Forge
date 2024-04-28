GOBIN := $(shell go env GOPATH)/bin
BIN_NAME := mycli

build:
	go build -o forge main.go && chmod +x forge && mv forge ${GOBIN}