package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "example.com/grpc/calculatorpb" // Import generated proto package
	"google.golang.org/grpc"
)

// Server struct implements CalculatorServiceServer
type server struct {
	pb.UnimplementedCalculatorServiceServer
}

// Add method implementation
func (s *server) Add(ctx context.Context, req *pb.AddRequest) (*pb.AddResponse, error) {
	fmt.Printf("Received request: %v\n", req) // log request
	result := req.Num1 + req.Num2
	return &pb.AddResponse{Result: result}, nil
}

func main() {
	// Listen on port 50051
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Create new gRPC server
	grpcServer := grpc.NewServer()

	// Register CalculatorService with server
	pb.RegisterCalculatorServiceServer(grpcServer, &server{})

	log.Println("Server is running on port 50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
