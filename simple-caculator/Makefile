.PHONY: all test cov

all: fmt test cov

test:
	go test -v ./...

cov:
	go test -coverprofile /tmp/c.out

fmt:
	gofmt -s -w .