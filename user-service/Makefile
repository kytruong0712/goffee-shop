# ----------------------------
# Env Variables
# ----------------------------
CONTAINER_NAME ?= "goffee-shop"
DATABASE_CONTAINER ?= database
API_CONTAINER ?= user-service
DOCKER_COMPOSE_FILE ?= docker-compose.local.yaml

## run: starts containers to run api server
run: api-create

## setup: executes pre-defined steps to setup api server
setup: pg-create pg-migrate

## api-create: starts api server
api-create:
	@echo Starting Api container
	docker-compose -f ${DOCKER_COMPOSE_FILE} -p=${CONTAINER_NAME} up ${API_CONTAINER}
	@echo Api container started!

## api-generate-models: executes CLI command to generate new database models
api-gen-models:
	@echo Starting generate db model...
	sh -c 'sqlboiler psql'
	@echo Done!

## api-go-generate: executes CLI command to generate go based config to handle graphql
api-go-generate:
	@echo Starting generate gqlgenerated file...
	sh -c 'go generate ./...'
	@echo Done!

## pg-create: starts postgres container
pg-create:
	@echo Starting Postgres database container
	docker-compose -f ${DOCKER_COMPOSE_FILE} -p=${CONTAINER_NAME} up -d ${DATABASE_CONTAINER}
	@echo Database container started!

## Create a DB migration files e.g `make new-migration-file name=init`
new-migration-file:
	docker-compose -f ${DOCKER_COMPOSE_FILE} -p=${CONTAINER_NAME} --profile tools run --rm migrate create -ext sql -dir /migrations $(name)

## pg-migrate: executes latest migration files
pg-migrate:
	@echo Running migration
	docker-compose -f ${DOCKER_COMPOSE_FILE} -p=${CONTAINER_NAME} --profile tools run --rm migrate up
	@echo Migration done!

pg-redo: pg-drop pg-migrate

protoc:
	cd internal/handler/grpcserver/proto && protoc --go_out=../protogen --go_opt=paths=source_relative \
	--go-grpc_out=../protogen --go-grpc_opt=paths=source_relative \
	./**/*.proto

## pg-drop: reset db to blank
pg-drop:
	@echo Dropping database...
	docker-compose -f ${DOCKER_COMPOSE_FILE} -p=${CONTAINER_NAME} --profile tools run --rm migrate drop
	@echo Done!

## down: removes all docker containers
down:
	@echo Removing containers...
	docker-compose -f -f ${DOCKER_COMPOSE_FILE} down
	@echo Done!
