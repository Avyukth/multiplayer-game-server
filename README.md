# lila-app Service

The `lila-app` service provides game statistics based on the geographic area code. This service is defined in the [Protocol Buffers (proto3)](https://developers.google.com/protocol-buffers) language and utilizes gRPC for service communication.

## Building and Running the Service

# GameStats Microservice Project

This is a GRPC-based microservice project that provides game stats based on the given area code.

## Environment Setup

Make sure to create an `.env` file at the root of your project with the necessary environment variables. A `.env.local` file is also used for local builds. Refer to `.env.example` file for an example setup.

## Building and Running with Docker

The following commands are available for building and running the project with Docker:

- Build the Docker images:  
  `make build`
- Start the Docker containers:  
  `make up`
- Stop the Docker containers:  
  `make down`
- Start the containers if they're stopped:  
  `make start`
- Stop the containers if they're running:  
  `make stop`
- Build and start the containers:  
  `make start-all`

## Running the GRPC Client

To run the GRPC client, use the command:
`make client`

## Building and Running Locally

The following commands are available for local builds and runs:

- Build the client binary:  
  `make build-local-bin-client`
- Build the server binary:  
  `make build-local-bin-server`
- Build both client and server binaries:  
  `make build-local-bin`
- Run the server:
  `make up && ./bin/server -e=.env.local`
- Run the client on different terminal other than the server:
  `make up && ./bin/client -e=.env.local`

## Client Run Configuration

```bash
cd clients && go run client.go
```
