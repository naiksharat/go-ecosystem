package main

import (
	"context"
	pb "grpc-course/greet/proto"
	"log"
)

func doLongGreet(c pb.GreetServiceClient) {
	log.Printf("Invoked long greet")
	longGreetClient, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("Error while getting long greet client %v\n", err)
		return
	}

	reqs := []*pb.GreetRequest{
		{FirstName: "Sharat"},
		{FirstName: "Manjunath"},
		{FirstName: "Naik"},
	}

	for _, req := range reqs {
		longGreetClient.Send(req)
		log.Printf("Sent %s", req)
	}

	resp, err := longGreetClient.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while receeving data from server at the very end %v", err)
	}

	log.Printf("Resp from the server %s", resp.Result)
}
