package main

import (
	"google.golang.org/grpc"
	"log"
	"net"

	pb "grpc-course/calculator/proto"
)

type Server struct {
	pb.CalcServiceServer
}

var addr string = "0.0.0.0:50051"

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen on %v\n", err)
	}
	defer lis.Close()

	log.Printf("listening on %s\n", addr)

	s := grpc.NewServer()
	pb.RegisterCalcServiceServer(s, &Server{})
	defer s.Stop()

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
