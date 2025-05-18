package main

import (
	"context"
	"log"
	"math/rand"
	"net"
	"sync"
	"time"

	pb "grpc-url-shortener/proto"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedURLShortenerServer
	mapping map[string]string
	mu      sync.RWMutex
}

func (s *server) Shorten(ctx context.Context, req *pb.ShortenRequest) (*pb.ShortenResponse, error) {
	shortCode := generateShortCode()
	s.mu.Lock()
	s.mapping[shortCode] = req.OriginalUrl
	s.mu.Unlock()
	return &pb.ShortenResponse{ShortCode: shortCode}, nil
}

func (s *server) Resolve(ctx context.Context, req *pb.ResolveRequest) (*pb.ResolveResponse, error) {
	s.mu.RLock()
	originalURL := s.mapping[req.ShortCode]
	s.mu.RUnlock()
	if originalURL == "" {
		return nil, grpc.Errorf(404, "Short code not found")
	}
	return &pb.ResolveResponse{OriginalUrl: originalURL}, nil
}

func generateShortCode() string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.Seed(time.Now().UnixNano())
	code := make([]byte, 6)
	for i := range code {
		code[i] = letters[rand.Intn(len(letters))]
	}
	return string(code)
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterURLShortenerServer(s, &server{mapping: make(map[string]string)})

	log.Println("gRPC server listening on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
