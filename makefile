.PHONY: start build build-prod clean-docker-images

default: start
start:
	swag init -g ./cmd/jubawink/main.go -o ./docs && docker compose up

build:
	docker compose up --build
	
build-prod:
	@if docker images -q golang-web-api; then \
		echo "Removing existing golang-web-api image..."; \
		docker rmi -f golang-web-api:latest || true; \
	fi
	docker buildx build -t golang-web-api:latest -f Dockerfile.prod .
clean-docker-images:
	docker images -f "dangling=true" -q | xargs -r docker rmi
