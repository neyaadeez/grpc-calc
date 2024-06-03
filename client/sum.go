package main

import (
	"context"
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
