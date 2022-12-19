package main

import (
	"fmt"
	pb "grpc-course/greet/proto"
	"io"
	"log"
)

func (s *Server) GreetEveryone(stream pb.GreetService_GreetEveryoneServer) error {
	log.Printf("GreetEveryone invoked")

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error while reading from stream %s", err)
			return err
		}

		err = stream.Send(&pb.GreetResponse{Result: fmt.Sprintf("Hello from server, %s", req.FirstName)})
		if err != nil {
			log.Fatalf("Error while sending data to stream %s", err)
			return err
		}

	}
}
