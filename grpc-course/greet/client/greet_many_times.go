package main

import (
	"context"
	pb "grpc-course/greet/proto"
	"io"
	"log"
)

func doGreetManyTimes(c pb.GreetServiceClient) {
	log.Printf("doGreetManyTimes invoked")
	greetManyTimesClient, err := c.GreetManyTimes(context.Background(), &pb.GreetRequest{FirstName: "Sharat"})
	if err != nil {
		log.Fatalf("Error while calling greet many rpc %v\n", err)
	}

	for {
		response, err := greetManyTimesClient.Recv()
		//time.Sleep(time.Second * 5)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while receiving stream data from server %v", err)
		}
		log.Printf("Greet many times %s", response.Result)
	}
}
