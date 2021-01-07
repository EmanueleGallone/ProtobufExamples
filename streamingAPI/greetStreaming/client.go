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

func biDirectionalStreaming(client greetStreamingpb.GreetServiceClient){
	fmt.Println("Starting Bi-directional streaming...")

	list := []greetStreamingpb.GreetEveryoneRequest{
		{Content: &greetStreamingpb.Greeting{
			FirstName: "Luke",
			LastName:  "Skywalker",
		}},
		{Content: &greetStreamingpb.Greeting{
			FirstName: "Alice",
			LastName:  "Bob",
		}},
	}

	stream,err :=client.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("error creating client stream: %v", err)
	}

	waitChannel := make(chan struct{})

	go func() { // goroutine to send a bunch of messages to server
		//sending bunch of requests
		for _, element := range list{
			_ = stream.Send(&element)
			time.Sleep(1000 * time.Millisecond)
		}
		_ = stream.CloseSend() //ending the stream
	}()

	go func() { // goroutine to receive the messages from server
		for true {
			res, err := stream.Recv() // receiving from server
			if err == io.EOF {
				break
			}
			if err != nil {
				close(waitChannel)
				log.Fatalf("Error receiving response from server: %v", err)
				return
			}

			fmt.Printf("Server response: %v\n", res.GetResult())
		}
		close(waitChannel) // closing waiting channel to end the function.
	}()

	<- waitChannel // this will block the function from ending (goroutine)

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

	//streamingLongGreet(client)

	biDirectionalStreaming(client)
}


