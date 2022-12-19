package main

import (
	"context"
	pb "grpc-course/greet/proto"
	"log"
	"time"
)

func doGreetWithDeadline(c pb.GreetServiceClient) {
	log.Println("doGreetWithDeadline was invoked")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()

	log.Printf("%+v", ctx)

	res, err := c.GreetWithDeadline(ctx, &pb.GreetRequest{FirstName: "Sharat"})
	if err != nil {
		log.Println("Client: ", err)
		return
	}

	log.Println(res.Result)
}
