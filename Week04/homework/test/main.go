package main

import (
	"context"
	"fmt"
	"log"
	"time"
	v1 "week04.homework.szs/api/article/v1"

	"google.golang.org/grpc"
)

const (
	address = "localhost:9090"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// grpc client
	c := v1.NewArticleClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// call rpc
	r, err := c.ContributeArticle(ctx, &v1.ArticleRequest{
		Id:    11,
		Title: "test",
	})
	if err != nil {
		fmt.Printf("save article failed: %v\n", err)
	}

	log.Printf("save aritcle id: %d", r.GetId())
}