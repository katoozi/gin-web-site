# note: call scripts from /scripts

help:
	@echo "Usage:"
	@echo "    make run                   run the server"
	@echo "    make build                 build the package"
	@echo "    make createsuperuser       create user with superuser permission"

build:
	go build cmd/gin-web-site/main.go

run:
	go run cmd/gin-web-site/main.go runserver

createsuperuser:
	go run cmd/gin-web-site/main.go createsuperuser

buildDockerCompose:
	docker-compose up --build

runDockerCompose:
	docker-compose up
