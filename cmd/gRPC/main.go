package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"inventory-manager/api/pb"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterInventoryServiceServer(s, &server{items: make(map[string]*pb.Item)})
	reflection.Register(s)
	log.Println("Server is running on port :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
