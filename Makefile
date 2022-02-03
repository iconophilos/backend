.DEFAULT_GOAL := help

PROJECT_NAME := Iconophilos Backend
IMAGE_NAME := iconophilos-backend
IMAGE_REPO := europe-west6-docker.pkg.dev/iconophilos-340108/iconophilos
IMAGE_TAG := $$(git rev-parse --short HEAD)

.PHONY: help

help:
	@echo "------------------------------------------------------------------------"
	@echo "${PROJECT_NAME}"
	@echo "------------------------------------------------------------------------"
	@grep -E '^[a-zA-Z0-9_/%\-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: up
up: ## spin up db and app server on docker containers
	@docker compose -f docker-compose.yaml up db -d
	@docker compose -f docker-compose.yaml up app --build

.PHONY: db_connect
db-connect: ## connect to db via psql
	@psql postgres://user:password@localhost:5432/iconophilos_db?sslmode=disable

.PHONY: test-unit
test-unit: ## run unit tests
	@go test -v -cover -count=1 -short ./...

.PHONY: test
test: ## run unit tests and integration tests
	@go test -v -cover -count=1 ./...

.PHONY: lint
lint: ## run linter
	@docker run --rm -v $(pwd):/app -w /app golangci/golangci-lint:v1.44.0 golangci-lint run -v

.PHONY: build
build: ## build docker image
	@docker build -t ${IMAGE_NAME} .
	@docker tag ${IMAGE_NAME} ${IMAGE_REPO}/${IMAGE_NAME}:${IMAGE_TAG}

.PHONY: push
push: ## push docker image to artifact registry
	@docker push ${IMAGE_REPO}/${IMAGE_NAME}:${IMAGE_TAG}
