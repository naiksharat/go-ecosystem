package main

import (
	"context"
	"google.golang.org/grpc/status"
	pb "grpc-course/calculator/proto"
	"log"
)

func getSqrt(c pb.CalcServiceClient, n int32) {
	log.Println("Client: invoked for n :", n)
	resp, err := c.Sqrt(context.Background(), &pb.CalcRequest{Number: n})
	if err != nil {
		//log.Println("Client:", err.Error())
		e, ok := status.FromError(err)
		if ok {
			log.Printf("Client: Error code %s, Message %s", e.Code(), e.Message())
			return
		} else {
			log.Printf("Client: Non grpc error")
		}

	}

	log.Println("Client:", resp.Sqrt)
}
