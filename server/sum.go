package main

import (
	"context"
	"log"

	pb "github.com/neyaadeez/grpc-calc/proto"
)

func (s *Server) AddNumbers(conx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Server invoked with the request: %v", in)
	return &pb.SumResponse{
		Result: in.Num_1 + in.Num_2,
	}, nil
}
