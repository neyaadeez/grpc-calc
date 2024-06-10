package main

import (
	"io"
	"log"

	pb "github.com/neyaadeez/grpc-calc/proto"
)

func (s *Server) StreamNumbers(stream pb.CalcService_StreamNumbersServer) error {
	log.Println("streamNumbers was invoked (bi-direction streaming)")

	var temp int32 = -1

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("error while reading the data from client: %v\n", err)
		}

		if temp < req.Num {
			temp = req.Num
			err = stream.Send(&pb.StreamNResponse{
				Res: temp,
			})

			if err != nil {
				log.Fatalf("error while sending data: %v\n", err)
			}
		}
	}
}
