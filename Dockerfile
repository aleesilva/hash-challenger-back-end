FROM golang:1.17

WORKDIR /app

ENV GO11MODULE=on


COPY go.mod .

COPY go.sum .

RUN apt-get update &&  apt-get install build-essential -y

RUN go get github.com/githubnemo/CompileDaemon

RUN go mod download

COPY . .

ENTRYPOINT CompileDaemon --build="go build -o ./bin/checkout_api" --command=./bin/checkout_api
