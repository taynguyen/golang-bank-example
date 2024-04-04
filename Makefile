current_dir = $(shell pwd)

help:
	@echo ''
	@echo 'Usage: make [TARGET] [EXTRA_ARGUMENTS]'
	@echo 'Targets:'
	@echo 'make dev: make dev for development work'
	@echo 'make build: make build container'
	@echo 'make production: docker production build'
	@echo 'clean: clean for all clear docker images'

init-tools:
	docker pull vektra/mockery

gen-mocks:
	docker run -v "$(current_dir)":/src -w /src vektra/mockery --all --inpackage

# TODO: Make test run via docker
.PHONY: test
test:
	go test -p 1 -coverprofile=c.out -failfast -timeout 5m ./...

dev:
	if [ ! -f .env ]; then cp .env.example .env; fi;
	docker-compose -f docker-compose-dev.yml down
	docker-compose -f docker-compose-dev.yml up

.PHONY: build
build:
	docker-compose -f docker-compose-dev.yml build
#	docker-compose -f docker-compose-prod.yml build

production:
	docker-compose -f docker-compose-prod.yml up -d --build

clean:
	docker-compose -f docker-compose-prod.yml down -v
	docker-compose -f docker-compose-dev.yml down -v
