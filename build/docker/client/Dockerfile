FROM golang:1.20-alpine

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

COPY go.mod go.sum ./
RUN go mod download
COPY . ./

EXPOSE 3000

CMD "go build -gcflags='all=-N -l' -o ./tmp/server ./cmd/server/main.go"