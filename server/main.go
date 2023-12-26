package main

import (
	"context"
	"flag"
	"fmt"
	pb "grpc/protoc"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 9000, "The server port")
)

type server struct {
	pb.UnimplementedExampleServer
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterExampleServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
func (s *server) ServerReply(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	var err error
	log.Printf("Received: %v", req.SomeString)
	// time.Sleep(5 * time.Second)
	return &pb.HelloResponse{Reply: fmt.Sprintf("Hello %v", req.SomeString)}, err
}

func (s *server) MemoryRequirments(ctx context.Context, req *pb.RequirmentsRequest) (*pb.SupplyResponse, error) {
	var err error
	log.Printf("Type: %v, Size: %v", req.Type, req.Size)
	availableResources := 10
	if availableResources > int(req.Size) {
		return &pb.SupplyResponse{SizeAvailable: 10}, err
	}
	return &pb.SupplyResponse{}, err
}
