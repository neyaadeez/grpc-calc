package main

import (
	"context"
	"io"
	"log"

	pb "github.com/neyaadeez/grpc-calc/proto"
)

func doSum(c pb.CalcServiceClient) {
	log.Printf("doSum was invoked")
	res, err := c.AddNumbers(context.Background(), &pb.SumRequest{
		Num_1: 55,
		Num_2: 100,
	})

	if err != nil {
		log.Printf("Error while making request: %v\n", err)
	}
	log.Printf("Result of addition of two numbers returned by server is : %d ", res.Result)
}

func doPrime(c pb.CalcServiceClient) {
	log.Printf("doPrime was invoked")
	stream, err := c.PrimeNumbers(context.Background(), &pb.PrimeNRequest{
		Num: 120,
	})

	if err != nil {
		log.Fatalf("Error while making request to server: %v\n", err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error occured while reading stream: %v\n", err)
		}

		log.Printf("Data received from the server: %d", res.Result)
	}
}
