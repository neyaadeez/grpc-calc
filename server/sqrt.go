package main

import (
	"context"
	"fmt"
	"log"
	"math"

	pb "github.com/neyaadeez/grpc-calc/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) Sqrt(ctx context.Context, in *pb.SqrtRequest) (*pb.SqrtResponse, error) {
	log.Println("Sqrt function was invoked!")
	log.Printf("Request: %v", in)

	if in.Num1 < 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Received a negative number: %d\n", in.Num1),
		)
	}

	return &pb.SqrtResponse{
		Res: math.Sqrt(float64(in.Num1)),
	}, nil
}
