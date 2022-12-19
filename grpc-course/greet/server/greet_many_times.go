package main

import (
	"fmt"
	pb "grpc-course/greet/proto"
	"log"
)

func (s *Server) GreetManyTimes(req *pb.GreetRequest, stream pb.GreetService_GreetManyTimesServer) error {
	log.Printf("Greet many invoked for req %v", req)
	for i := 0; i < 10; i++ {
		resp := fmt.Sprintf("%d: hello, %s", i, req.FirstName)
		err := stream.Send(&pb.GreetResponse{Result: resp})
		if err != nil {
			return err
		}
		//time.Sleep(time.Second * 2)
	}

	return nil
}
