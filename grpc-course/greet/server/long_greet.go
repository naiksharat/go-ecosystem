package main

import (
	pb "grpc-course/greet/proto"
	"io"
	"log"
)

func (s *Server) LongGreet(stream pb.GreetService_LongGreetServer) error {
	log.Println("Invoked long greet")
	res := "Hello!"
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			stream.SendAndClose(&pb.GreetResponse{Result: res})
			return nil
		} else {
			log.Printf("received %s", req.FirstName)
		}
		if err != nil {
			log.Fatalf("Error while receiveing stream data from client %v\n", err)
			return err
		}
		res += req.FirstName + ","
	}

}
