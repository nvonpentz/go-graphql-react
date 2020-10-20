SHELL := /bin/bash

build: server client

server:
	cd cmd && CGO_ENABLED=0 go build -o ../bin/server && cd ..

client:
	npm run build --prefix frontend

run-server:
	source .env.sh && go run cmd/main.go

run-client:
	npm run start --prefix frontend

docker: lint
	docker build . -t go-graphql-react

docker-up:
	docker-compose up

graphql:
	gqlgen generate

db:
	go run github.com/volatiletech/sqlboiler/v4 psql --wipe

db-reset:
	dropdb postgres -U postgres && createdb postgres -U postgres

models: db graphql

lint:
	go fmt ./...

test:
	source .env.sh && go test -v ./...
