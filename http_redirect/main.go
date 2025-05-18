package main

import (
	"context"
	"log"
	"net/http"

	pb "grpc-url-shortener/proto"

	"github.com/gin-gonic/gin"

	"google.golang.org/grpc"
)

func main() {
	// Connect to gRPC server
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	client := pb.NewURLShortenerClient(conn)

	r := gin.Default()

	r.GET("/short/:code", func(c *gin.Context) {
		code := c.Param("code")

		// Call gRPC Resolve
		res, err := client.Resolve(context.Background(), &pb.ResolveRequest{
			ShortCode: code,
		})
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Short code not found"})
			return
		}

		c.Redirect(http.StatusFound, res.OriginalUrl)
	})

	log.Println("HTTP redirect server running at http://localhost:8080")
	r.Run(":8080")
}
