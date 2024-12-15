.PHONY: build test docker-build docker-run docker-stop lint proto clean

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
	protoc --go_out=. --go_opt=paths=source_relative \
    	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
    	proto/user/*.proto

clean:
	rm -rf bin/
	go clean
	rm -f proto/user/*.pb.go

dev: docker-stop docker-build docker-run
	@echo "Development environment is ready"

migrate:
	migrate -path migrations -database "postgresql://postgres:postgres@localhost:5432/userdb?sslmode=disable" up

generate: proto
	go generate ./...