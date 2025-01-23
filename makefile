.PHONY: start build build-prod clean-docker-images docker-ecr-push

default: start
start:
	swag init -g ./cmd/jubawink/main.go -o ./docs && docker compose up

build:
	docker compose up --build

build-debug:
	docker-compose -f docker-compose.debug.yml up --build

debug:
	docker-compose -f docker-compose.debug.yml up

build-prod:
	@if docker images -q 028788912057.dkr.ecr.us-east-1.amazonaws.com/jubawink-api; then \
		echo "Removing existing 028788912057.dkr.ecr.us-east-1.amazonaws.com/jubawink-api image..."; \
		docker rmi -f 028788912057.dkr.ecr.us-east-1.amazonaws.com/jubawink-api || true; \
	fi
	aws ecr get-login-password --region us-east-1 | docker login --username AWS --password-stdin 028788912057.dkr.ecr.us-east-1.amazonaws.com && \
	docker buildx build -t jubawink-api -f Dockerfile.prod . && \
	docker tag jubawink-api:latest 028788912057.dkr.ecr.us-east-1.amazonaws.com/jubawink-api:latest && \
	docker push 028788912057.dkr.ecr.us-east-1.amazonaws.com/jubawink-api:latest
clean-docker-images:
	docker images -f "dangling=true" -q | xargs -r docker rmi
