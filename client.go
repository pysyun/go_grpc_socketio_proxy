package main

import (
    "time"
    "golang.org/x/net/context"
    "google.golang.org/grpc"
    "./events"
    "log"
)

func main() {
    conn, err := grpc.Dial("34.239.11.167:84", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("failed to connect: %s", err)
    }
    defer conn.Close()

    // NewStreamClient - ./events/events.pb.go
    client := events.NewStreamClient(conn)
    stream, err := client.Connect(context.Background())

    waitc := make(chan struct{})
    // Message - ./events/events.pb.go
    msg := &events.Message{Event: "publish", Data: string("{\"channel\":\"))\",\"message\":\"::))\"}")}

    go func() {
        for i := 0; i < 3; i++ {
            log.Println("Sleeping...")
            time.Sleep(1 * time.Second)
            log.Println("Sending msg Event '" + msg.Event + "' and Data '" + msg.Data + "'")
            stream.Send(msg)

            log.Println("And then...")
            getMSG, err := stream.Recv()
            if err != nil {
                log.Fatalf("failed: %s", err)
            }
            log.Println("We get '" + getMSG.Event + "' and '" + getMSG.Data + "'")
        }
        time.Sleep(1 * time.Second)
        close(waitc)
    }()
    <-waitc
    stream.CloseSend()
}
