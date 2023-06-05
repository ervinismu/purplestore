help: ## You are here! showing all command documenentation.
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

migrate-create: ## Create migration file
	docker compose --profile tools run --rm migrate create -ext sql -dir /migrations $(name)

migrate-up:  ## Run migrations UP
	docker compose --profile tools run --rm migrate up

migrate-down:  ## Run migrations DOWN
	docker compose --profile tools run --rm migrate down

shell-db: ## Enter to database console docker
	docker compose exec db psql -U postgres -d postgres

run: ## Run http server
	go run cmd/main.go

environment: ## Init app environment
	docker compose up -d

lint: ## Running linter
	docker compose --profile tools run --rm lint golangci-lint run

test: ## Running testing
	go test ./... -v -coverprofile=output.html

test-cover: ## Open test coverage
	go tool cover -html=output.html
