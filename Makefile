# Variables
PROJECT = go-api-boilerplate
COMPOSE_COMMAND := docker-compose -f ./docker-compose.yml

### Start selected environment
start:
	${COMPOSE_COMMAND} -p $(PROJECT) up -d

### Build and Start selected environment
restart:
	${COMPOSE_COMMAND} -p $(PROJECT) kill \
 && ${COMPOSE_COMMAND} -p $(PROJECT) rm -f \
 && ${COMPOSE_COMMAND} -p $(PROJECT) up -d --build

### Shut down selected environment
stop:
	${COMPOSE_COMMAND} -p $(PROJECT) kill \
 && ${COMPOSE_COMMAND} -p $(PROJECT) rm -f \

### Dell selected environment
down:
	${COMPOSE_COMMAND} -p $(PROJECT) down

ps: ## List containers : ## make ps
	${COMPOSE_COMMAND} -p $(PROJECT) ps

### Show logs from selected environment
logs:
	${COMPOSE_COMMAND} -p $(PROJECT) logs -f
