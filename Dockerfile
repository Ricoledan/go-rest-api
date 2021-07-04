FROM golang:1.16.5-alpine

WORKDIR /src

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o go-rest-api .

EXPOSE 9000

CMD ["./go-rest-api"]
