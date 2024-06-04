package main

import (
	"context"
	"log"
	"time"

	pb "github.com/neyaadeez/grpc-calc/proto"
)

func doAvg(c pb.CalcServiceClient) {
	log.Println("doAvg was invoked")

	reqs := []*pb.AvgNRequest{
		{Num: 10},
		{Num: 50},
		{Num: 80},
		{Num: 90},
		{Num: 100},
	}

	stream, err := c.AvgNumbers(context.Background())

	if err != nil {
		log.Fatalf("unable to send request to server: %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("sending stream of request to server: %v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Printf("error while receiving response from the server: %v\n", err)
	}

	log.Printf("received data from the server: %d\n", res.Res)
}
