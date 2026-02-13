.PHONY: build test lint run clean docker-build

build:
	go build -o bin/server ./cmd/server

test:
	go test ./... -v

lint:
	golangci-lint run

run:
	go run ./cmd/server

clean:
	rm -rf bin/

docker-build:
	docker build -t go-ci-books .
