/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package main implements a server for Greeter service.
package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	pb "grpc/protoc"
	"log"
	"net"

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
	log.Printf("Received: %v", req.SomeString)
	return &pb.HelloResponse{}, errors.New("")
}

func (s *server) MemoryRequirments(ctx context.Context, req *pb.RequirmentsRequest) (*pb.SupplyResponse, error) {
	log.Printf("Type: %v, Size: %v", req.Type, req.Size )
	return &pb.SupplyResponse{}, errors.New("")
}
