.PHONY: build test docker-build docker-run

build:
    go build -o bin/server ./cmd/server

test:
    go test -v ./...

docker-build:
    docker-compose build

docker-run:
    docker-compose up -d

docker-stop:
    docker-compose down

lint:
    golangci-lint run

proto:
    protoc --go_out=. --go-grpc_out=. proto/user/*.proto