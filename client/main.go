package main

import (
	"log"

	pb "github.com/neyaadeez/grpc-calc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var host string = "localhost:50051"

func main() {
	conn, err := grpc.NewClient(host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Unable to connect to server: %v\n", err)
	}

	defer conn.Close()

	c := pb.NewCalcServiceClient(conn)
	// doSum(c)
	doPrime(c)
}
