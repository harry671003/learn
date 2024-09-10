package main

import (
	"log"
	"net"

	"github.com/harry671003/grpc-learn/chat"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")

	if err != nil {
		log.Fatalf("Failed to listen on port 9000: %v", err)
	} else {
		log.Println("Running")
	}

	srv := grpc.NewServer()
	s := &chat.Server{}
	chat.RegisterChatServiceServer(srv, s)

	if err := srv.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
