# Golang GRPC Socket.io proxy


### Here are some requirements that help us create and use these files:
### First: it's a good idea to use golang:alpine image for testing our application.
### Second: inside your running container that based on golang:alpine image you will need to run the command:
- apk add protoc
- apk add git
- go get github.com/golang/protobuf/protoc-gen-go
- go get golang.org/x/net/context
- go get google.golang.org/grpc

These commands helps you use protoc and some modules that will need for running our program.
### Third: when you get your events.proto file you will need to create a directory with the name events, where you save events.pb.go file.
Also, for creating events.pb.go file you will need to use the next command:
protoc --go_out=plugins=grpc:events events.proto

