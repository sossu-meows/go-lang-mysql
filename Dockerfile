FROM golang:1.20.5 AS builder
WORKDIR /usr/src/app
COPY . .
RUN go build -v -o /usr/local/bin/app ./...