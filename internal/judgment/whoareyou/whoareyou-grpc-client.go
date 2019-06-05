package whoareyou

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"time"

	pb "github.com/littletommytan/garen/internal/justice/hellogrpc"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:2233"
	defaultName = "world"
)

func grpcHello(c *gin.Context) {
	query := c.DefaultQuery("nihao", defaultName)
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	connection := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	name := query
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := connection.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	c.String(200, "Grpc greeting: %s", r.Message)
}
