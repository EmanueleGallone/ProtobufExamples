package main

import (
	"ProtobufExamples/streamingAPI/greetStreaming/greetStreamingpb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"time"
)

func streamingLongGreet(client greetStreamingpb.GreetServiceClient){

	list := []greetStreamingpb.LongGreetRequest{
		{Content: &greetStreamingpb.Greeting{
			FirstName: "Luke",
			LastName:  "Skywalker",
		}},
		{Content: &greetStreamingpb.Greeting{
			FirstName: "Lucy",
			LastName:  "Madeleine",
		}},
	}

	stream ,err := client.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("error starting service: %v", err)
	}

	for _, request := range list {
		_ = stream.Send(&request)
		time.Sleep(1000 * time.Millisecond)
	}

	response, _ := stream.CloseAndRecv()
	fmt.Println(fmt.Sprintf("received response: %v", response))
}

func streamingAPI(client greetStreamingpb.GreetServiceClient){
	req := greetStreamingpb.GreetManyRequest{Content: &greetStreamingpb.Greeting{
		FirstName: "Qwerty",
		LastName:  "Uiop",
	}}

	responseStream, err := client.GreetMany(context.Background(), &req)
	if err != nil{
		log.Fatalf("error sending greet many: %v", err)
	}

	//Handling the stream
	for {
		msg, err := responseStream.Recv() //receiving messages through stream
		if err == io.EOF { //receiving EOF means stream has ended sending
			fmt.Println("stream ended")
			break
		}

		if msg != nil {
			fmt.Printf("Received: %v \n", msg.GetResult())
		}
	}
}

func main() {
	fmt.Print("Starting client \n")

	conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("error dialing: %v ", err)
	}

	defer conn.Close()

	client := greetStreamingpb.NewGreetServiceClient(conn)

	//streamingAPI(client)

	streamingLongGreet(client)
}


