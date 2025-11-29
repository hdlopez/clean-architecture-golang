.PHONY: all run test lint clean

all: test lint build

run:
	go run main.go

test:
	go test -v ./...

lint:
	golangci-lint run

clean:
	go clean
	rm -f clean-architecture-golang
