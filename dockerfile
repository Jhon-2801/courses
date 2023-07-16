FROM golang:latest AS builder

RUN apt-get update

WORKDIR /app
COPY go.mod .
RUN go mod download
COPY . .
RUN go get -d -v ./...
RUN go build -o app -v

ENTRYPOINT ["./app"]