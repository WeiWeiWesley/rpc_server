FROM golang:1.12-alpine

COPY . /go/src/rpc_server
WORKDIR /go/src/rpc_server

RUN go build -o app main.go

ENTRYPOINT [ "./app" ]