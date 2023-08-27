#####################
#### DEVELOPMENT ####
#####################

.PHONY: setup
setup:
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate
	go install github.com/matryer/moq
	go install github.com/cosmtrek/air
	go install github.com/golangci/golangci-lint/cmd/golangci-lint
	go install golang.org/x/tools/cmd/goimports

.PHONY: docker-up
docker-up:
	@docker compose up -d

.PHONY: docker-down
docker-down:
	@docker compose down --remove-orphans

.PHONY: goimports
goimports:
	@goimports -w $(shell find . -name '*.go' | xargs)

####################
#### MIGRATIONS ####
####################
PG_ADDR ?= "postgresql://moki:moki@localhost:5432/moki?sslmode=disable"

.PHONY: create-migration
create-migration:
	migrate create -seq -dir gateways/postgres/migrations -ext sql $(name)

.PHONY: migrate
migrate:
	@go run cmd/migration/main.go -database $(PG_ADDR)

###################
####   TESTS   ####
###################

.PHONY: test
test:
	@go test -v -race -vet=all -count=1 -coverprofile=coverage.out ./...

.PHONY: lint
lint:
	@golangci-lint run ./...

.PHONY: lint/ci
lint/ci:
	@golangci-lint --out-format github-actions run ./...
