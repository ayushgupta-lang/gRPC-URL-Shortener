package main

import (
	"context"
	"log"
	"time"

	pb "grpc-url-shortener/proto"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewURLShortenerClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Shorten URL
	originalURL := "https://google.com"
	shortResp, err := client.Shorten(ctx, &pb.ShortenRequest{OriginalUrl: originalURL})
	if err != nil {
		log.Fatalf("could not shorten: %v", err)
	}
	log.Printf("Short code: %s", shortResp.ShortCode)

	// Resolve URL
	resolveResp, err := client.Resolve(ctx, &pb.ResolveRequest{ShortCode: shortResp.ShortCode})
	if err != nil {
		log.Fatalf("could not resolve: %v", err)
	}
	log.Printf("Original URL: %s", resolveResp.OriginalUrl)
}
