BINARY=api_server

run-local:
	go run ./internal/main.go

build-dc:
	docker-compose \
	-f docker/compose/docker-compose.dev.yml \
	--project-directory ./ \
	build

up:
	docker-compose \
	-f docker/compose/docker-compose.dev.yml \
	--project-directory ./ \
	up

build:
	go build -v -o ./${BINARY} ./internal/main.go