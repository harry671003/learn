package chat

import (
	"io"
	"log"
)

type Server struct {
	ChatServiceServer
}

func (s *Server) SayHello(stream ChatService_SayHelloServer) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			log.Println("Stopping")
			return nil
		}
		if err != nil {
			return err
		}
		log.Printf("Received message from client: %s", in.Body)
		stream.Send(&Message{Body: "Hello from the server"})
	}
}
