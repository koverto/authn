.DEFAULT_GOAL := run

.PHONY: build
build: gen
	go build ./cmd/credentials

.PHONY: docker
docker: build
	docker build . -t koverto/credentials:latest

.PHONY: gen
gen:
	go generate ./api

.PHONY: run
run: gen
	go run ./cmd/credentials
