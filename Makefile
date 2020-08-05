SHELL := /bin/bash

.DEFAULT_GOAL := help

.PHONY: start-api
start-api:
	go run ./cmd/api/dev/


.PHONY: start-frontend
start-frontend:
	cd frontend/ && yarn install && yarn serve

.PHONY: help
help:
	@$(MAKE) -pRrq -f $(lastword $(MAKEFILE_LIST)) : 2>/dev/null | awk -v RS= -F: '/^# File/,/^# Finished Make data base/ {if ($$1 !~ "^[#.]") {print $$1}}' | sort | egrep -v -e '^[^[:alnum:]]' -e '^$@$$'

	
