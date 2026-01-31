APP_NAME=e-commerce-api
MAIN_PATH=cmd/api/main.go

.PHONY: help run build test tidy lint clean

help:
	@echo "Available commands:"
	@echo "  make run     - Run application"
	@echo "  make build   - Build binary"
	@echo "  make test    - Run tests"
	@echo "  make tidy    - Download & tidy dependencies"
	@echo "  make lint    - Run linter"
	@echo "  make clean   - Remove build artifacts"

run:
	go run $(MAIN_PATH)

build:
	go build -o bin/$(APP_NAME) $(MAIN_PATH)

test:
	go test ./... -cover

lint: 
	golangci-lint run

migrate-up:
	migrate -path migrations -database "$(DATABASE_URL)" up

migrate-down:
	migrate -path migrations -database "$(DATABASE_URL)" down

docker-up:
	docker compose up --build

docker-down:
	docker compose down