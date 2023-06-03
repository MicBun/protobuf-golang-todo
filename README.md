# protobuf-golang-todo
A Simple Todo app using Protocol Buffer, Wire, Postgres, Mockery, and Docker

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


## How to run

### Using Docker Compose
1. Modify .env file to match .env.dev
2. Alternatively, you can copy .env.dev to .env
```bash
cp .env.dev .env
```
3. Run the app
```bash
docker-compose up
```
*Note: first time running this command will take a while to download the images,
but the next time will be faster, also the first time may fail because the app
is trying to connect to the database before it is ready, just run the command
again and it should work.*

### Using Go run
1. Setup Local Postgres database
    1. Create database `todo`
2. Modify .env file to match .env.local
3. Alternatively, you can copy .env.local to .env
```bash
cp .env.local .env
```
4. Make sure postgres credentials match with your local postgres
5. Run the app
```bash
go run cmd/server/main.go
```
6. Tables will be created automatically

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