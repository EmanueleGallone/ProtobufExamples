package main

import (
	"ProtobufExamples/greetpb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
)

func UnaryRPCCall(client greetpb.GreetServiceClient) {

	//fmt.Printf("Created client: %v", client)
	req := greetpb.GreetRequest{Content: &greetpb.Greeting{
		FirstName: "qwerty",
		LastName:  "azerty",
	},
	}

	res, err := client.Greet(context.Background(), &req)
	if err != nil{
		log.Fatalf("error in getting response: %v", err)
	}
	fmt.Println(res.GetResult())
}

func main() {
	fmt.Print("Starting client \n")

	conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("error dialing: %v ", err)
	}

	defer conn.Close()

	client := greetpb.NewGreetServiceClient(conn)

	UnaryRPCCall(client)
}


