.PHONY: fmt test

all: fmt test

fmt:
	go fmt ./...

test:
	go test -v -cover -count=1 -test.cpu=1 ./...