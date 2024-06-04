package main

import (
	"io"
	"log"

	pb "github.com/neyaadeez/grpc-calc/proto"
)

func (s *Server) AvgNumbers(stream pb.CalcService_AvgNumbersServer) error {
	log.Println("avgNumbers was invoked")

	var res []int32

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			var finalres int32
			for _, num := range res {
				finalres += num
			}
			finalres = finalres / int32(len(res))
			return stream.SendAndClose(&pb.AvgResponse{
				Res: finalres,
			})
		}

		if err != nil {
			log.Fatalf("error while reading the data from client: %v\n", err)
		}

		res = append(res, req.Num)
	}
}
