GOBIN := $(shell go env GOPATH)/bin
BIN_NAME := forge

cli-tool:
	cd cli && go build -o forge main.go && chmod +x ${BIN_NAME} && mv forge ${GOBIN}