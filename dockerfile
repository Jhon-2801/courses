FROM golang:latest AS builder
RUN apt-get update
ENV GO111MODULE=on \
		CGO_ENABLED=0 \
		COOS=linux \
		GOARCH=amd64
WORKDIR /app
COPY go.mod .
RUN go mod download
COPY . .
RUN go build main.go

FROM scratch
COPY --from=builder /app .
ENTRYPOINT ["./main"]