ifndef $(GOROOT)
    GOROOT=$(shell go env GOROOT)
    export GOROOT
endif

include .env
export

api-ini:
	@echo -e "\n\t🔑\n"
	@go run $(GOROOT)/src/crypto/tls/generate_cert.go --host localhost

api-gen:
	@echo -e "\n\t🧠\n"
	@sqlc generate

show-schema:
	@echo -e "\n\t🐘\n"
	@go run ./cmd/migration/migration.go

backup-db:
	@echo -e "\n\t🐘\n"
	@/usr/local/Cellar/libpq/13.2/bin/pg_dump --format=p -h 127.0.0.1 -d ${POSTGRES_DB} -p ${POSTGRES_PORT} -U ${POSTGRES_USER} -f ./backup.sql

DIR=$(notdir $(shell pwd))
export DIR

api-dev:
	@echo "\n\t💣\n"
	docker-compose -p $(DIR) up --build --force-recreate --remove-orphans

.PHONY: api-ini api-gen show-schema api-dev backup-db