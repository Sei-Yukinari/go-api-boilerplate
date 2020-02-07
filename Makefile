# Variables
COMPOSE_COMMAND := docker-compose -f ./docker/docker-compose.yml

### Start selected environment
start:
	${COMPOSE_COMMAND} up -d

### Build and Start selected environment
restart:
	${COMPOSE_COMMAND} up -d --build

### Shut down selected environment
stop:
	${COMPOSE_COMMAND}  down -v

### Show logs from selected environment
logs:
	${COMPOSE_COMMAND} logs -f
