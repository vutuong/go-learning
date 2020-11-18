package main

import (
	"log"
	"context"
	"net"
	"google.golang.org/grpc"
	"example.com/grpc-server/tuong"
)

type Server struct {}

func (s *Server) SayHello(ctx context.Context, message *tuong.Message) (*tuong.Message, error) {
	log.Printf("Received message body from client: %s", message.Body)
	return &tuong.Message{Body: "Hello From the Server!"}, nil
}

func main() {
	lis, err := net.Listen("tcp",":9000")
	if err != nil {
		log.Fatalf("Failed to listen on port 9000: %v", err)
	}

	// Creates a new gRPC server
	s := grpc.NewServer()
	tuong.RegisterTuongServiceServer(s, &Server{})
	if err:= s.Serve(lis); err !=nil {
		log.Fatalf("Failed to serve GRPC server over port 9000 : %v", err)
	}
	
}