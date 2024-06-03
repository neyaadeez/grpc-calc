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

func (s *Server) PrimeNumbers(in *pb.PrimeNRequest, steam pb.CalcService_PrimeNumbersServer) error {
	log.Printf("server prime numbers was invoked with input: %v", in)

	var val int32 = in.Num
	var k int32 = 2
	for val != 0 {
		if val%k == 0 {
			steam.Send(&pb.PrimeNResponse{
				Result: k,
			})
			val = val / k
		} else {
			k = k + 1
		}
	}

	return nil
}
