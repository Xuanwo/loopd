package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/Xuanwo/loopd/service/grpc"
)

func main() {
	lis, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterLoopdServer(s, &pb.Server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
