#!make
include .env

APP_NAME := golang-e-commerce-api
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
	docker run --rm \
	  --network golang-e-commerce-api_default \
	  -v $(PWD)/migrations:/migrations \
	  migrate/migrate \
	  -path /migrations \
	  -database "$(DB_URL)" \
	  up

migrate-down:
	docker run --rm \
	  --network golang-e-commerce-api_default \
	  -v $(PWD)/migrations:/migrations \
	  migrate/migrate \
	  -path /migrations \
	  -database "$(DB_URL)" \
	  down

dev:
	docker compose up --build

# dev:
# 	docker compose build --no-cache api
prod:
	docker compose -f docker-compose.prod.yml up --build

docker-down:
	docker compose down

docker-up:
	docker compose up -d

psql:
	docker compose exec db psql -U postgres -d e_commerce_db

logs:
	docker compose logs -f api
