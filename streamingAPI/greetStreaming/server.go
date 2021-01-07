package main

import (
	"ProtobufExamples/streamingAPI/greetStreaming/greetStreamingpb"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"time"
)

type server struct {}

func (*server) GreetMany(req *greetStreamingpb.GreetManyRequest, stream greetStreamingpb.GreetService_GreetManyServer) error {
	firstName := req.GetContent().GetFirstName()

	for i := 0; i < 10; i++ {
		content := fmt.Sprintf("Hello %v, id: %d", firstName, i)
		msg := greetStreamingpb.GreetManyResponse{Result: content}
		if err := stream.Send(&msg); err != nil {
			log.Fatalf("Error sending msg: %v", err)
		}

		time.Sleep(1000 * time.Millisecond)
	}

	return nil
}

func (*server) LongGreet(stream greetStreamingpb.GreetService_LongGreetServer) error{
	response := "Hello "
	for {
		msg, err := stream.Recv()
		response += msg.GetContent().GetFirstName() + " "
		if err == io.EOF {
			fmt.Sprintln("Reached end")
			_ = stream.SendAndClose(&greetStreamingpb.LongGreetResponse{Result: response})
			break
		}
	}
	return nil
}

/*
Used in bi-directional streaming
 */
func (*server) GreetEveryone(stream greetStreamingpb.GreetService_GreetEveryoneServer) error {
	fmt.Println("Greet Everyone invoked")
	var result string

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("error receiving from stream: %v", err)
			return err
		}

		result = "Hello " + req.GetContent().GetFirstName() + ";"
		res := greetStreamingpb.GreetEveryoneResponse{
			Result: result,
		}

		if err := stream.Send(&res); err != nil{
			log.Fatalf("error sending response: %v", err)
			return err
		}

	} //end of for true

	return nil
}

func main() {
	fmt.Println("Starting server")

	listener, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	greetStreamingpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
