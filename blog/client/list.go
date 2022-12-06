package main

import (
	"context"
	"io"
	"log"

	pb "github.com/MochamadAkbar/go-grpc-example/blog/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

func ListBlogs(c pb.BlogServiceClient) {
	log.Println("--- listBlogs was invoked ---")

	stream, err := c.ListBlogs(context.Background(), &emptypb.Empty{})

	if err != nil {
		log.Fatalf("Error whole calling ListBlog: %v\n", err)
	}

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Something happened: %v\n", err)
		}

		log.Println(res)
	}
}
