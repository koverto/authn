.DEFAULT_GOAL := run

.PHONY: build
build: gen
	go build ./cmd/authn

.PHONY: docker
docker: build
	docker build . -t koverto/authn:latest

.PHONY: gen
gen:
	go generate ./api

.PHONY: run
run: gen
	go run ./cmd/authn
