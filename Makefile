BINARY_NAME=spirex
MAIN=cmd/server/main.go

run:
	go run $(MAIN)

up:
	docker compose up -d

down:
	docker compose down -v --rmi all && yes | docker system prune --all

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o bin/$(BINARY_NAME) $(MAIN) 

test:
	go test ./... -v

cover:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html

fmt:
	go fmt ./...

lint:
	golangci-lint run

clean:
	rm -rf bin/ $(BINARY_NAME) coverage.out coverage.html

.PHONY: run build test cover fmt lint clean docs

