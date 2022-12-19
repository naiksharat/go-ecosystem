package main

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	pb "grpc-course/blog_without_pojo/proto"
	"log"
)

/*
	CreateBlog(context.Context, *Blog) (*BlogId, error)
	ReadBlog(context.Context, *BlogId) (*Blog, error)
	UpdateBlog(context.Context, *Blog) (*emptypb.Empty, error)
	DeleteBlog(context.Context, *BlogId) (*emptypb.Empty, error)
	ListBlogs(*emptypb.Empty, BlogService_ListBlogsServer) error
*/

func (s *Server) CreateBlog(ctx context.Context, in *pb.Blog) (*pb.BlogId, error) {
	log.Println("Create rpc invoked")
	res, err := collection.InsertOne(ctx, in)

	if err != nil {
		log.Println(err)
	}
	return &pb.BlogId{Id: res.InsertedID.(primitive.ObjectID).Hex()}, nil
}
