.PHONY: test fmt

all: test fmt

test:
	go test -cover ./...

fmt:
	go fmt ./...