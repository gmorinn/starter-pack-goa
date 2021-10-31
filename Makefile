ifndef $(GOROOT)
    GOROOT=$(shell go env GOROOT)
    export GOROOT
endif

include .env
export

DIR=$(notdir $(shell pwd))
export DIR

api-goa:
	@echo -e "\n\t🔥GOA GM X SQUIRREL🔥\n\n\tLoading...⌛\n"
	@goa gen $(DIR)/design
	@cp -f gen/http/openapi.json ./documentation
	@mv gen/http/openapi.json ./
	@echo -e "\nEnjoy🐿️\n=> Documentation make api-doc"

api-doc:
	@cd documentation && live-server

api-init:
	@echo -e "\n\t🔑\n"
	@go run $(GOROOT)/src/crypto/tls/generate_cert.go --host localhost

api-gen:
	@echo -e "\n\t🧠\n"
	@sqlc generate


backup-db:
	@echo -e "\n\t🐘\n"
	@pg_dump --inserts -h  ${POSTGRES_HOST} -d ${POSTGRES_DB} -p ${POSTGRES_PORT} -U ${POSTGRES_USER} -f ./backup.sql

api-dev:
	@echo "\n\t💣\n"
	docker-compose -p $(DIR) up --build --force-recreate --remove-orphans

.PHONY: api-init api-gen api-doc api-goa show-schema api-dev backup-db