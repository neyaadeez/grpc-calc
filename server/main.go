package main

import (
	"log"
	"net"

	pb "github.com/neyaadeez/grpc-calc/proto"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.CalcServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("unable to setup network connection: %v\n", err)
	}

	log.Printf("listening on address: %s\n", addr)

	s := grpc.NewServer()
	pb.RegisterCalcServiceServer(s, &Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to setup server: %v\n", err)
	}
}
