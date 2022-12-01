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
	@cp -f gen/http/openapi3.json ./
	@echo -e "\nWait...⌛\n\nGOA GM will generate functions for you❤️\n"
	@goa example $(DIR)/design
	@clean
	@echo -e "\nEnjoy🐿️\n=> Documentation make api-doc"

doc:
	@echo "\n\tCopy paste "openapi3.json" (file in the root of the project) in this website\n"
	@echo "\t🔗 https://editor.swagger.io/\n"
	@open https://editor.swagger.io/

api-init:
	@echo -e "\n\t🔑\n"
	@go run $(GOROOT)/src/crypto/tls/generate_cert.go --host $(API_DOMAIN)

sql:
	@echo "\n\t🧠\n"
	@./sql/bin/sqlc generate
	@echo "\nQueries generated\n"

api-dev:
	@echo -e "\n\t💣\n"
	@rm -f cert.pem key.pem
	@go run $(GOROOT)/src/crypto/tls/generate_cert.go --host $(API_DOMAIN)
	docker-compose -p ${PROJECT} up --build --force-recreate --remove-orphans

.PHONY: api-init sql doc api-goa show-schema api-dev

