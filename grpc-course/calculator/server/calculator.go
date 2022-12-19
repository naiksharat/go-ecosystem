package main

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	pb "grpc-course/calculator/proto"
	"log"
	"math"
)

func (s *Server) Sqrt(ctx context.Context, in *pb.CalcRequest) (*pb.CalcResponse, error) {
	log.Printf("Server: sqrt of %d", in.Number)

	if in.Number < 0 {
		return nil, status.Error(codes.InvalidArgument, "Negative numbers not allowed")
	}

	return &pb.CalcResponse{Sqrt: float32(math.Sqrt(float64(in.Number)))}, nil
}
