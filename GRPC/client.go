package main

import (
	"context"
	"log"
	"time"

	pb "example.com/grpc/calculatorpb"
	"google.golang.org/grpc"
)

func main() {
	// Connect to server
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	// Create client
	c := pb.NewCalculatorServiceClient(conn)

	// Set timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Call Add RPC
	res, err := c.Add(ctx, &pb.AddRequest{Num1: 10, Num2: 20})
	if err != nil {
		log.Fatalf("error while calling Add RPC: %v", err)
	}

	log.Printf("Result from server: %d", res.Result)
}
