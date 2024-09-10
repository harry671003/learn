package main

import (
	"context"
	"io"
	"log"

	"github.com/harry671003/grpc-learn/chat"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial(":9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Error dialing %v", err)
	}
	defer conn.Close()

	client := chat.NewChatServiceClient(conn)
	ctx, _ := context.WithCancel(context.Background())

	stream, err := client.SayHello(ctx)
	if err != nil {
		log.Fatalf("Error dialing %v", err)
	}

	waitc := make(chan struct{})
	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				close(waitc)
				return
			}
			log.Printf("%v, %v", in, err)
		}
	}()

	for i := 0; i < 10; i++ {
		log.Println("Sending..")
		stream.Send(&chat.Message{Body: "Something"})
	}
	stream.CloseSend()

	<-waitc
}
