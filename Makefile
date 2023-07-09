-include .env
export

build:
	docker-compose build

up:
	docker-compose up -d

down:
	docker-compose down

start:
	docker-compose start

stop:
	docker-compose stop

start-all: build up

client:
	cd clients && go run client.go

build-local-bin-client:
	go build -o bin/client clients/client.go
	chmod +x bin/client

build-local-bin-server:
	go build -o bin/server servers/server.go
	chmod +x bin/server

build-local-bin: build-local-bin-client build-local-bin-server
