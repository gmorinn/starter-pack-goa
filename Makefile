ifndef $(GOROOT)
    GOROOT=$(shell go env GOROOT)
    export GOROOT
endif

include .env
export

DIR=$(notdir $(shell pwd))
export DIR

api-goa:
	@echo -e "\n\t🔥GOA GM\n\n\tLoading...⌛\n"
	@goa gen $(DIR)/design
	@cp -f gen/http/openapi.json ./cmd/documentation/
	@cp -f gen/http/openapi3.json ./
	@echo -e "\nWait...⌛\n\nGOA GM will generate functions for you❤️\n"
	@goa example $(DIR)/design
	@go build -o cl cmd/clean/clean.go cmd/clean/cleanHTTP.go cmd/clean/cleanMAIN.go cmd/clean/cleanFolderAPI.go && ./cl
	@rm cl
	@echo -e "\nEnjoy🐿️\n=> Documentation make api-doc"

api-doc:
	@cd cmd/documentation && live-server

api-init:
	@echo -e "\n\t🔑\n"
	@go run $(GOROOT)/src/crypto/tls/generate_cert.go --host localhost

api-gen:
	@echo -e "\n\t🧠\n"
	@sqlc generate

backup-db:
	@echo -e "\n\t🐘\n"
	@pg_dump --inserts -h  ${POSTGRES_HOST} -d ${POSTGRES_DB} -p ${POSTGRES_PORT} -U ${POSTGRES_USER} -f ./backup.sql

createdb:
	docker exec -it postgres createdb --username=${POSTGRES_USER} --owner=${POSTGRES_USER} ${POSTGRES_DB}

dropdb:
	docker exec -it postgres dropdb ${POSTGRES_DB}

api-test:
	@go test -coverprofile=coverage.out ./internal/
	@	go tool cover -html=coverage.out


startpostgres:
	docker run --name postgres -p ${POSTGRES_PORT}:${POSTGRES_PORT} -e POSTGRES_USER=${POSTGRES_USER} -e POSTGRES_PASSWORD=${POSTGRES_PASSWORD} -d postgres:12-alpine

api-dev:
	@echo -e "\n\t💣\n"
	docker-compose -p $(DIR) up --build --force-recreate --remove-orphans

.PHONY: api-init api-gen api-doc api-goa show-schema api-dev backup-db createdb dropdb startpostgres api-test

