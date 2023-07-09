# Build stage
FROM golang:1.20-alpine AS builder

# Install the protobuf compiler and git
RUN apk add --no-cache protobuf git

WORKDIR /app
COPY . .


RUN go mod download

# Set up Go path
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPATH=/root/go \
    PATH="$PATH:/root/go/bin"

# Install the Go protobuf plugin
RUN GOBIN=/root/go/bin go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN GOBIN=/root/go/bin go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Generate Go code from all .proto files in the 'proto' directory
RUN mkdir -p api
RUN for f in proto/*.proto; do protoc --go_out=./api --go_opt=paths=source_relative --go-grpc_out=./api --go-grpc_opt=paths=source_relative $f; done

RUN go build -o server ./servers/server.go

# Final stage
FROM gcr.io/distroless/static-debian11

WORKDIR /app

COPY --from=builder /app/server /server
COPY .env /.env
EXPOSE $PORT
CMD ["/server"]
