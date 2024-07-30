# ----------------------------
# Env Variables
# ----------------------------
CONTAINER_NAME ?= "goffee-shop"
API_CONTAINER ?= server
DOCKER_COMPOSE_FILE ?= docker-compose.local.yaml

## run: starts containers to run api server
run: api-create

## api-create: starts api server
api-create:
	@echo Starting Api container
	docker-compose -f ${DOCKER_COMPOSE_FILE} -p=${CONTAINER_NAME} up ${API_CONTAINER}
	@echo Api container started!

## api-go-generate: executes CLI command to generate go based config to handle graphql
api-go-generate:
	@echo Starting generate gqlgenerated file...
	sh -c 'go generate ./...'
	@echo Done!

## down: removes all docker containers
down:
	@echo Removing containers...
	docker-compose down
	@echo Done!