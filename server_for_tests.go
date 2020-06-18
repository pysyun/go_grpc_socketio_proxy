package main
import (
	"log"
	"net"
	"google.golang.org/grpc"
	"./events"
)

func main() {

	s := events.EventsServer{}

	grpcServer := grpc.NewServer()
	// RegisterStreamServer - ./events/events.pb.go
	events.RegisterStreamServer(grpcServer, &s)

	l, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Println("Listening on tcp://localhost:8081")
	grpcServer.Serve(l)
}
