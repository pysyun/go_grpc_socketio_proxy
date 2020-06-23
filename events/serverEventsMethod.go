package events

import (
        "io"
        "log"
)

type EventsServer struct {
}

// Stream_ConnectServer - ./events/events.pb.go
func (s *EventsServer) Connect(stream Stream_ConnectServer) error {
    log.Println("Started stream")
    for {
        in, err := stream.Recv()
        log.Println("Received value")
        if err == io.EOF {
                return nil
        }
        if err != nil {
                return err
        }
        log.Println("Received Event '" + in.Event + "' and received Data '" + in.Data + "'")

        if err := stream.Send(in); err != nil {
            return err
        }
        log.Println("Sended Event '" + in.Event + "' and sended Data '" + in.Data + "'")
    }
}
