package main

import (
	"ProtobufExamples/Calculator/calculatorpb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
)

type server struct {}

func (*server) Calculate(ctx context.Context, req *calculatorpb.CalculatorRequest) (*calculatorpb.CalculatorResponse, error){
	operation := req.GetOperation()
	firstNumber := req.GetContent().GetFirstNumber()
	secondNumber := req.GetContent().GetSecondNumber()

	var result int32 = 0

	switch operation {
		case calculatorpb.Operation_sum:
			result = firstNumber + secondNumber

		case calculatorpb.Operation_subtraction:
			result = firstNumber - secondNumber

		case calculatorpb.Operation_multiplication:
			result = firstNumber * secondNumber

		case calculatorpb.Operation_division:
			result = firstNumber / secondNumber

		default:
			log.Fatalln("Error in  parsing operation:")
	}

	response := calculatorpb.CalculatorResponse{Result: float32(result)}
	return &response, nil
}

func (*server) Average(stream calculatorpb.CalculatorService_AverageServer) error{
	fmt.Println("Received Average Request")

	var array []int32
	var average float32

	for {
		msg, err := stream.Recv()
		array = append(array, msg.GetNumber())

		if err == io.EOF {
			for _, num := range array {
				average += float32(num)
			}
			average = average / (float32(len(array)))

			response := &calculatorpb.CalculatorResponse{Result: average}
			_ = stream.SendAndClose(response)
			break
		}
	}
	return nil
}

func main() {
	fmt.Println("Starting server")

	listener, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
