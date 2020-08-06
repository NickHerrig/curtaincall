SHELL := /bin/bash

.DEFAULT_GOAL := help

.PHONY: start-api
start-api:
	go run ./cmd/api/dev/


.PHONY: start-frontend
start-frontend:
	cd frontend/ && yarn install && yarn serve


.PHONY: build-api
build-api:
	go build -o curtaincall ./cmd/api/prod


.PHONY: build-frontend
build-frontend:
	cd frontend/ && yarn install && yarn build


.PHONY: dist
dist: build-api build-frontend
	mkdir -p dist/
	cp -r setup.sql dist/setup.sql
	cp -r curtaincall dist/curtaincall
	cp -r frontend/dist/ dist/frontend/

.PHONY: clean
clean: 
	rm -rf curtaincall frontend/dist dist/

.PHONY: help
help:
	@$(MAKE) -pRrq -f $(lastword $(MAKEFILE_LIST)) : 2>/dev/null | awk -v RS= -F: '/^# File/,/^# Finished Make data base/ {if ($$1 !~ "^[#.]") {print $$1}}' | sort | egrep -v -e '^[^[:alnum:]]' -e '^$@$$'

	
