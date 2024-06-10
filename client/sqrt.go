package main

import (
	"context"
	"log"

	pb "github.com/neyaadeez/grpc-calc/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doSqrt(c pb.CalcServiceClient, in int32) {
	log.Println("doSqrt function was invoked!")

	res, err := c.Sqrt(context.Background(), &pb.SqrtRequest{
		Num1: in,
	})

	if err != nil {
		e, ok := status.FromError(err)

		if ok {
			log.Printf("error from server: %v\n", e.Message())
			log.Printf("error code from server: %v\n", e.Code())

			if e.Code() == codes.InvalidArgument {
				log.Printf("Negative number error!")
				return
			}
		} else {
			log.Fatalf("A non-GRPC Error: %v\n", err)
		}
	}

	log.Printf("Response from the server: %f\n", res.Res)
}
