# Variables
PROJECT = go-api-boilerplate
COMPOSE_PATH := ./docker-compose.yml
COMPOSE_COMMAND := docker-compose -f $(COMPOSE_PATH) -p $(PROJECT)
SERVICE_API := api
SERVIVE_DB := db

### Start selected environment
.PHONY: start
start:
	${COMPOSE_COMMAND} up -d

### Build and Start selected environment
.PHONY: restart
restart:
	${COMPOSE_COMMAND} kill \
 && ${COMPOSE_COMMAND} rm -f \
 && ${COMPOSE_COMMAND} up -d --build

### Shut down selected environment
.PHONY: stop
stop:
	${COMPOSE_COMMAND} kill \
 && ${COMPOSE_COMMAND} rm -f \

### Dell selected environment
.PHONY: down
down:
	${COMPOSE_COMMAND} down -v

 ## List containers
.PHONY: ps
ps:
	${COMPOSE_COMMAND} ps

### Show logs from selected environment
.PHONY: logs
logs:
	${COMPOSE_COMMAND} logs -f api

.PHONY: exec-api
exec-api:
	${COMPOSE_COMMAND} exec ${SERVICE_API} ash

.PHONY: exec-db
exec-db:
	${COMPOSE_COMMAND} exec ${SERVIVE_DB} bash
