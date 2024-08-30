# ----------------------------
# Env Variables
# ----------------------------
PROJECT_NAME ?= "goffee-shop"
API_CONTAINER ?= api-gateway
DOCKER_COMPOSE_FILE ?= build/docker-compose.local.yaml

## run: starts containers to run api server
run: api-create

## api-create: starts api server
api-create:
	@echo Starting Api container
	docker-compose -f ${DOCKER_COMPOSE_FILE} -p=${PROJECT_NAME} up ${API_CONTAINER}
	@echo Api container started!

## api-go-generate: executes CLI command to generate go based config to handle graphql
api-go-generate:
	@echo Starting generate gqlgenerated file...
	sh -c 'go generate ./...'
	@echo Done!

## Ex: make protoc-gateway name=user
protoc-gateway:
	cd internal/gateway/$(name)/protobuf && protoc \
	--go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative \
	./*.proto

## down: removes all docker containers
down:
	@echo Removing containers...
	docker-compose down
	@echo Done!
