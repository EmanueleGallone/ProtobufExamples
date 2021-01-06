package main

import (
	"ProtobufExamples/greetpb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error){
	firstName := req.GetContent().GetFirstName()
	lastName := req.GetContent().GetLastName()

	resp := greetpb.GreetResponse{Result: fmt.Sprintf("Greetings from server %v %v", firstName, lastName)}
	return &resp, nil
}

func main() {
	fmt.Print("Starting server")

	listener, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
