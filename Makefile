.DEFAULT_GOAL := help

apigen:
	cd backend && \
	oapi-codegen --config=./configs/oapi.yaml ../api/api.yaml

up:
	docker compose -f build/docker-compose.yml up --build

help:
	@echo "make apigen	- generate API endpoints for backend"
	@echo "make up		- build and start application in docker"
