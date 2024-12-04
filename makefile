.PHONY: start build

default start:
	swag init -g ./cmd/jubawink/main.go -o ./docs && docker compose up

build:
	docker compose up --build