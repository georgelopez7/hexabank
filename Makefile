# -----
# 🐘 MIGRATIONS

POSTGRES_DB_URL ?= "postgres://user:password@localhost:5432/db"
PATH_TO_DB_MIGRATIONS = "./services/payment/migrations"

## Creates a database migration file
# Example: "make db-create-migration MIGRATION_NAME=add_puzzles_table"
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