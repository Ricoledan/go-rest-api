# go-rest-api

## commands

run application
`go run main.go`

build executable (windows)
`go build -o go-rest-api.exe main.go`

### docker

build docker image
`docker build --tag go-rest-api .`
run docker container
`docker run -d -p 9000:9000 go-rest-api`
