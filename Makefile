BINARY=vd

run-local:
	go run ./internal/main.go

build:
	docker-compose \
	-f docker/compose/docker-compose.dev.yml \
	--project-directory ./ \
	build

up:
	docker-compose \
	-f docker/compose/docker-compose.dev.yml \
	--project-directory ./ \
	up