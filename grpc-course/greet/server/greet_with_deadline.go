package main

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	pb "grpc-course/greet/proto"
	"log"
	"time"
)

func (s *Server) GreetWithDeadline(ctx context.Context, req *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Println("Invoked GreetWithDeadline")
	for i := 0; i < 3; i++ {
		err := ctx.Err()
		if err == context.DeadlineExceeded {
			log.Printf("Context deadline exceeded")
			return nil, status.Error(codes.Canceled, "The client canceled the request")
		}
		if err != nil {
			log.Printf("Unknown error %s", err)
			return nil, err
		}

		time.Sleep(time.Second * 1)
	}

	return &pb.GreetResponse{Result: "Server: Hello"}, nil
}
