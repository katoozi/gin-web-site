# note: call scripts from /scripts
all: help
help:
	@echo "Usage:"
	@echo "    make build                 build docker stack file"
	@echo "    make up                    run the services"
	@echo "    make down                  stop the services"
	@echo "    make logs                  show all services logs"
	@echo "    make log                   show special service logs => make service=web-server log"
	@echo "    make deploy                deploy the stack   => make stack_name=postgres stack_file=postgres/docker-stack.yml deploy"
	@echo "    make rebuild               rebuild service(s) => make services='postgres web-server redis' rebuild"
	@echo "    make initial-network       create project network in docker"

build:
	docker-compose --compatibility -f postgres/docker-compose.yml -f redis-persistence/docker-compose.yml -f project/docker-compose.yml -f reverse-proxy/docker-compose.yml config > docker-compose.yml

up:
	docker-compose --compatibility up --build -d

logs:
	docker-compose logs -f

log:
	docker-compose logs -f ${service}

rebuild:
	docker-compose --compatibility up -d --no-deps --build ${services}

down:
	docker-compose --compatibility down

deploy:
	docker stack deploy -c ${stack_file} ${stack_name}

initial-network:
	docker network create project
