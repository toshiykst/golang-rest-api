# Build
FROM golang:1.15.2-alpine as builder

WORKDIR /go/src/github.com/toshiykst/golang-rest-api

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o golang-rest-api app/server.go

# Exec
FROM alpine:latest

COPY --from=builder /go/src/github.com/toshiykst/golang-rest-api/golang-rest-api golang-rest-api

EXPOSE 8080

CMD ["./golang-rest-api"]