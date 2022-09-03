build:
	go build -o service ./cmd/core-service/main.go

lint:
	gofumpt -w .
	go mod tidy
	golangci-lint run ./...

up:
	docker-compose up -d

down:
	docker-compose down