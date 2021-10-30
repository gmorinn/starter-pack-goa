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
	@echo -e "\nEnjoy🐿️\n=> Documentation make api-doc"

api-doc:
	@cd documentation && live-server

api-ini:
	@echo -e "\n\t🔑\n"
	@go run $(GOROOT)/src/crypto/tls/generate_cert.go --host localhost

api-gen:
	@echo -e "\n\t🧠\n"
	@sqlc generate


backup-db:
	@echo -e "\n\t🐘\n"
	@/usr/local/Cellar/libpq/13.2/bin/pg_dump --format=p -h 127.0.0.1 -d ${POSTGRES_DB} -p ${POSTGRES_PORT} -U ${POSTGRES_USER} -f ./backup.sql

api-dev:
	@echo "\n\t💣\n"
	docker-compose -p $(DIR) up --build --force-recreate --remove-orphans

.PHONY: api-ini api-gen show-schema api-dev backup-db