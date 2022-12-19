package main

import (
	"context"
	"fmt"
	pb "grpc-course/greet/proto"
	"log"
)

func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Server invoked: %+v", in)
	return &pb.GreetResponse{
		Result: fmt.Sprintf("Hi %s (from server)", in.FirstName),
	}, nil
}
