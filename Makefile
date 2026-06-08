TEMPL := templ
GO := go
DOCKER_IMAGE := gootmplhtmx
DOCKER_CONTAINER := gootmplhtmx-container

.PHONY: generate run build tidy docker-build docker-run docker-stop

# Your original commands
generate:
	$(TEMPL) generate

run: generate
	$(GO) run main.go

build: generate
	$(GO) build .

tidy:
	$(GO) mod tidy

# New Docker commands
docker-build:
	docker build -t $(DOCKER_IMAGE) .

docker-run: docker-build
	docker run -p 4000:4000 -e APP_PORT=4000 -e DB_HOST=host.docker.internal -e DB_PORT=5432 -e DB_USER=postgres -e DB_PASSWORD=0909 -e DB_NAME=serverside -e DB_SSLMODE=disable --name $(DOCKER_CONTAINER) $(DOCKER_IMAGE)

docker-stop:
	docker stop $(DOCKER_CONTAINER) || true
	docker rm $(DOCKER_CONTAINER) || true

docker-dev: docker-stop docker-run

# Run with PostgreSQL on your PC (no Docker PostgreSQL)
local-with-docker: docker-build
	docker run -p 4000:4000 -e APP_PORT=4000 -e DB_HOST=host.docker.internal -e DB_PORT=5432 -e DB_USER=postgres -e DB_PASSWORD=0909 -e DB_NAME=serverside -e DB_SSLMODE=disable $(DOCKER_IMAGE)