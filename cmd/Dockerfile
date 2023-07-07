FROM golang:1.20.5-alpine3.18 AS builder
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY cmd/main.go .
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o ./bin/app