package main

import (
	"context"
	pb "grpc-course/blog_without_pojo/proto"
	"log"
)

func createBlog(c pb.BlogServiceClient) string {
	res, err := c.CreateBlog(context.Background(), &pb.Blog{
		AuthorId: "Sharat",
		Title:    "My journey through Germany",
		Content:  "Ich bin Sharat",
	})

	if err != nil {
		log.Fatalf("Unexpected error: %v\n", err)
	}

	log.Printf("Blog has been created: %v\n", res.Id)
	return res.Id
}
