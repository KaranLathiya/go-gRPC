package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	pb "grpc/protoc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewExampleClient(conn)

	HelloResponse, err := c.ServerReply(context.TODO(), &pb.HelloRequest{SomeString: "Karan"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
		return
	}
	fmt.Println(HelloResponse)
	RequirementsResponse, err := c.MemoryRequirments(context.TODO(), &pb.RequirmentsRequest{Type: "RAM", Size: 8})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
		return
	}
	fmt.Println(RequirementsResponse)
	
}
