# -----
# 🐘 MIGRATIONS

POSTGRES_DB_URL ?= "postgres://user:password@localhost:5432/db"
PATH_TO_DB_MIGRATIONS = "./services/payment/adapters/postgres/migrations"

## Creates a database migration file
# Example: "make db-create-migration MIGRATION_NAME=add_new_table"
db-create-migration:
	goose -dir=$(PATH_TO_DB_MIGRATIONS) postgres $(POSTGRES_DB_URL) create $(MIGRATION_NAME) sql

## Runs all the migration files against the database
# Example: "make db-migrate"
db-migrate:
	goose -dir=$(PATH_TO_DB_MIGRATIONS) postgres $(POSTGRES_DB_URL) up

## Removes the last migration file from the database
# Example: "make db-rollback-last-migration"
db-rollback-last-migration:
	goose -dir=$(PATH_TO_DB_MIGRATIONS) postgres $(POSTGRES_DB_URL) down

# -----
# 📡 GRPC PROTOBUFS
SERVICE ?= "fraud"
PROTO_PATH = "api/proto/$(SERVICE)/$(SERVICE).proto"

## Generates protobufs for the service
# Example: "make gen-protobufs SERVICE=fraud"
gen-protobufs:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative $(PROTO_PATH)

# -----
# ⚡️ TESTS

## Runs all the tests
# Example: "make test"
test:
	go test ./...

## Runs tests for a particular microservice
# Example: "make test-microservice SERVICE=payment"
SERVICE ?= "payment"
test-microservice:
	go test ./services/$(SERVICE)/...
