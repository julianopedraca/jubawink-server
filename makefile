.PHONY: start build build-prod clean-docker-images docker-ecr-push

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
	docker run -d --publish 8888:8080 golang-web-api
docker-ecr-push:
	docker tag golang-web-api:latest 028788912057.dkr.ecr.us-east-1.amazonaws.com/teste:latest
	docker push 028788912057.dkr.ecr.us-east-1.amazonaws.com/teste:latest
clean-docker-images:
	docker images -f "dangling=true" -q | xargs -r docker rmi
