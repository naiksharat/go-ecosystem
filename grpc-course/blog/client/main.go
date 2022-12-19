package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "grpc-course/blog/proto"
	"log"
)

var addr string = "0.0.0.0:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatal("Failed to connect %v\n", err)
	}
	defer conn.Close()

	blogServiceClient := pb.NewBlogServiceClient(conn)

	bid := createBlog(blogServiceClient)

	readBlog(blogServiceClient, bid)

}
