.PHONY: build
build:
	go build -v ./cmd/notes

img-build:
	docker rmi notes_image; docker build -t notes_image .

start:
	docker compose down && docker compose up --build -d

stop:
	docker compose down

init:
	chmod ugo+x tools/initilize.sh && tools/./initilize.sh

init-test:
	chmod ugo+x tools/initilize-test.sh && tools/./initilize-test.sh

run-pg:
	docker run --name postgres --env POSTGRES_PASSWORD=123456 --volume postgres-volume:/var/lib/postgresql/data --publish 5432:5432 --detach postgres

del-pg:
	docker stop postgres; docker rm postgres

migrate:
	migrate -path migrations -database "postgres://postgres:123456@localhost:5432/notes_dev?sslmode=disable" -verbose up

migrate-test:
	migrate -path migrations -database "postgres://postgres:123456@localhost:5432/notes_test?sslmode=disable" -verbose up

.PHONY: test
test:
	go test -v -race -timeout 30s ./...

.DEFAULT_GOAL := build