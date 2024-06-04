package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/neyaadeez/grpc-calc/proto"
)

func doStream(c pb.CalcServiceClient) {
	log.Println("doStream was invoked")

	reqs := []*pb.StreamNRequest{
		{Num: 1},
		{Num: 5},
		{Num: 3},
		{Num: 6},
		{Num: 2},
		{Num: 20},
	}

	stream, err := c.StreamNumbers(context.Background())

	if err != nil {
		log.Fatalf("unable to send request to server: %v\n", err)
	}

	waitc := make(chan struct{})

	go func() {
		for _, req := range reqs {
			log.Printf("sending stream of request to server: %v\n", req)
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalf("Error occured while reading stream: %v\n", err)
			}

			log.Printf("Data received from the server: %d", res.Res)
		}
		close(waitc)
	}()

	<-waitc
}
