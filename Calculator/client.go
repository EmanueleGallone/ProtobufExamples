package main

import (
	"ProtobufExamples/Calculator/calculatorpb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
)

func UnaryRPCCall(client calculatorpb.CalculatorServiceClient) {

	//fmt.Printf("Created client: %v", client)
	req := calculatorpb.CalculatorRequest{
		Content: &calculatorpb.Number{
			FirstNumber:  1,
			SecondNumber: 2,
		},
		Operation: calculatorpb.Operation_sum,
	}

	res, err := client.Calculate(context.Background(), &req)
	if err != nil{
		log.Fatalf("error in getting response: %v", err)
	}
	fmt.Println(fmt.Sprintf("result: %d",res.GetResult()))
}

func streamingRPC(client calculatorpb.CalculatorServiceClient) error{
	list := []calculatorpb.AverageRequest{
		{Number: 10},
		{Number: 15},
		{Number: 20},
		{Number: 25},
	}
	stream, err := client.Average(context.Background())

	if err != nil {
		log.Fatalf("error opening stream: %v", err)
	}

	for _, num := range list {
		_ = stream.Send(&num)
	}

	res, _ := stream.CloseAndRecv()
	fmt.Printf("Average: %v\n", res.GetResult())

	return nil
}
func main() {
	fmt.Print("Starting client \n")

	conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("error dialing: %v ", err)
	}

	defer conn.Close() //deferring here means that the connection will be closed when function is exiting

	client := calculatorpb.NewCalculatorServiceClient(conn)

	//UnaryRPCCall(client)

	streamingRPC(client)

}
