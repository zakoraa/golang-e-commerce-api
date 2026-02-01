APP_NAME := e-commerce-api
MAIN_PATH := cmd/api/main.go

.PHONY: help run build test tidy lint clean migrate-up migrate-down docker-up docker-down logs

help:
	@echo "Available commands:"
	@echo "  make run          - Run application inside Docker"
	@echo "  make build        - Build binary inside Docker"
	@echo "  make test         - Run tests inside Docker"
	@echo "  make tidy         - Download & tidy dependencies"
	@echo "  make lint         - Run golangci-lint"
	@echo "  make migrate-up   - Apply database migrations"
	@echo "  make migrate-down - Rollback database migrations"
	@echo "  make docker-up    - Start all Docker services"
	@echo "  make docker-down  - Stop all Docker services"
	@echo "  make logs         - Tail api logs"

run:
	docker compose run --rm api go run $(MAIN_PATH)

build:
	docker compose run --rm api go build -o bin/$(APP_NAME) $(MAIN_PATH)

test:
	docker compose run --rm api go test ./... -cover

tidy:
	docker compose run --rm api go mod tidy

lint:
	docker compose run --rm api golangci-lint run

clean:
	rm -rf bin

migrate-up:
	docker compose run --rm api migrate \
		-path migrations \
		-database "$$DATABASE_URL" up

migrate-down:
	docker compose run --rm api migrate \
		-path migrations \
		-database "$$DATABASE_URL" down

docker-up:
	docker compose up --build -d

docker-down:
	docker compose down

logs:
	docker compose logs -f api
