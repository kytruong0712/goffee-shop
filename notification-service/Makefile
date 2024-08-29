# ----------------------------
# Env Variables
# ----------------------------
PROJECT_NAME ?= "goffee-shop"
API_CONTAINER ?= notification-service
DOCKER_COMPOSE_FILE ?= build/docker-compose.local.yaml

## run: starts containers to run api server
run: api-create

## api-create: starts api server
api-create:
	@echo Starting Api container
	docker-compose -f ${DOCKER_COMPOSE_FILE} -p=${PROJECT_NAME} up ${API_CONTAINER}
	@echo Api container started!

protoc:
	cd internal/handler/grpcserver/proto && protoc --go_out=../protogen --go_opt=paths=source_relative \
	--go-grpc_out=../protogen --go-grpc_opt=paths=source_relative \
	./**/*.proto

protoc-handler:
	cd internal/handler/grpc/protobuf && protoc \
	--go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative \
	./*.proto
