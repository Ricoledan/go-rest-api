FROM golang:1.16.5-alpine3.14 AS build
WORKDIR /app

ADD go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -o . .
