NAME := bookshelf_go

PHONY: local-build
local-build:
	docker-compose -f ./local/docker-compose.yaml -p ${NAME} build

PHONY: local-up
local-up:
	docker-compose -f ./local/docker-compose.yaml -p ${NAME} up -d

PHONY: local-down
local-down:
	docker-compose -f ./local/docker-compose.yaml -p ${NAME} down

PHONY: local-clean
local-clean:
	docker-compose -f ./local/docker-compose.yaml -p ${NAME} down -v
