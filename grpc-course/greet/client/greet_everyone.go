package main

import (
	"context"
	pb "grpc-course/greet/proto"
	"io"
	"log"
	"time"
)

/*
func doGreetEveryone(c pb.GreetServiceClient) {
	log.Println("Invoked doGreetEveryone")

	req := []*pb.GreetRequest{
		{FirstName: "Sharat"},
		{FirstName: "Manjunath"},
		{FirstName: "Naik"},
	}

	greetEveryoneClient, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("Error while getting greet everyone client")
		return
	}

	for _, req := range req {
		err := greetEveryoneClient.Send(req)
		if err != nil {
			log.Fatalf("Error while sending req to server")
			return
		}
		resp, err := greetEveryoneClient.Recv()
		if err != nil {
			log.Fatalf("Error while receiving resp from server")
			return
		}

		log.Printf("resp from server: %s", resp.Result)
	}

	err = greetEveryoneClient.CloseSend()
	if err != nil {
		log.Fatalf("Error while sending close req")
		return
	}
}
*/

func doGreetEveryone(c pb.GreetServiceClient) {
	log.Println("Invoked doGreetEveryone")

	reqs := []*pb.GreetRequest{
		{FirstName: "Sharat"},
		{FirstName: "Manjunath"},
		{FirstName: "Naik"},
	}

	greetEveryoneClient, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("Error while getting greet everyone client")
		return
	}

	waitc := make(chan struct{})

	go func() {
		for _, req := range reqs {
			err := greetEveryoneClient.Send(req)
			if err != nil {
				log.Fatalf("error while sending data to server")
			} else {
				log.Printf("sent %s", req)
			}
			time.Sleep(time.Second * 1)
		}
		err := greetEveryoneClient.CloseSend()
		if err != nil {
			log.Fatalf("error while closing stream")
		}
	}()

	go func() {
		for {
			resp, err := greetEveryoneClient.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("error while reading from stream")
			}
			log.Printf("resp from server: %s", resp.Result)
		}
		close(waitc)
	}()

	<-waitc
}
