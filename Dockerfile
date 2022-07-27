FROM golang:1.18.4-alpine

WORKDIR /app

COPY . .
RUN go mod download && go mod verify