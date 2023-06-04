# protobuf-golang-todo
A Simple Todo app using Golang, Protocol Buffer, Wire, Postgres, Mockery, Grpc, and Docker

## Things to Install
1. [Docker Desktop](https://www.docker.com/products/docker-desktop)
    1. When using Docker Desktop,you can skip the rest of the steps
2. [Go](https://golang.org/doc/install)
3. [Wire](https://github.com/google/wire)
    1. Wire is used for dependency injection
4. [Mockery](https://github.com/vektra/mockery)
    1. Mockery is used for mocking interfaces
5. [Postgres](https://www.postgresql.org/download)
    1. Postgres is used as the database
6. [gRPC](https://grpc.io/docs/languages/go/quickstart/)
    1. gRPC is used for communication between services

## How to run

### Using Docker Compose
1. Modify .env accordingly
2. Build the app
```bash
docker-compose build
```
3. Run the app
```bash
docker-compose up
```
*Note: first time running this command will take a while to download the images,
but the next time will be faster, also the first time may fail because the app
is trying to connect to the database before it is ready, just run the command
again and it should work. In case it is still not working, modify the cmd in
.air.toml and add a sleep command before running the app (look at the .air2.toml
for example)*

### Using Go run
1. Setup Local Postgres database
    1. Create database `todo`
2. Modify .env file accordingly
3. Make sure postgres credentials match with your local postgres
4. Run the app
```bash
go run cmd/server/main.go
```
5. Tables will be created automatically
6. Run the client app
```bash
go run cmd/client/main.go
```
7. It should execute the commands in the cmd/client/main.go file and print the result

## Usefully Scripts
### Generate Protobuf
```bash
sh scripts/protoc.sh
```

### Generate Mocks
```bash
sh scripts/mockery.sh
```

### Generate Code Coverage Report
```bash
sh scripts/coverage.sh
```

### Generate Wire Dependency Injection
```bash
sh scripts/wire.sh
```

### Check all scripts
```bash
sh scripts/check_all.sh
```