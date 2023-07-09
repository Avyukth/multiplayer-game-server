package main

import (
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
)




func main(){
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
	os.Exit(0)
}
