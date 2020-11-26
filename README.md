# Golang GRPC Socket.io proxy

## Installing pre-requisites to Alpine Linux (like golang:alpine in Docker)
```shell
apk add protoc
apk add git
go get github.com/golang/protobuf/protoc-gen-go
go get golang.org/x/net/context
go get google.golang.org/grpc
go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

## Compiling the protocols file
```shell
protoc --go_out=plugins=grpc:events events.proto
```
