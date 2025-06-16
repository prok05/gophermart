ifneq ($(wildcard .env),)
include .env
export
else
$(warning WARNING: .env file not found! Using .env.example)
include .env.example
export
endif

BASE_STACK = docker compose -f docker-compose.yml

.PHONY: help

help: ## Display this help screen
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

----: ### ---

compose-up: ### Run docker compose (Postgres)
	$(BASE_STACK) up --build -d db && docker compose logs -f
.PHONY: compose-up-all

compose-up-full: ### Run docker compose (Postgres + App)
	$(BASE_STACK) up --build -d
.PHONY: compose-up-all

compose-down: ### Down docker compose
	$(BASE_STACK) down --remove-orphans
.PHONY: compose-down

--------: ### --------

swag-v1: ### swag init
	@swag init -g internal/controller/http/router.go && swag fmt
.PHONY: swag-v1

deps: ### deps tidy + verify
	go mod tidy && go mod verify
.PHONY: deps

run: deps swag-v1 ### deps + swag + migrations + run
	go mod download && \
	CGO_ENABLED=0 go run -tags migrate ./cmd/app
.PHONY: run

--------: ### --------

migrate-create:  ### create new migration
	migrate create -ext sql -dir migrations '$(word 2,$(MAKECMDGOALS))'
.PHONY: migrate-create

migrate-up: ### migration up
	migrate -path migrations -database '$(PG_URL)?sslmode=disable' up
.PHONY: migrate-up


migrate-dwn: ### migration down
	migrate -path migrations -database '$(PG_URL)?sslmode=disable' down
.PHONY: migrate-down


