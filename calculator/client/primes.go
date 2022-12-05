package main

import (
	"context"
	"io"
	"log"

	pb "github.com/MochamadAkbar/go-grpc-example/calculator/proto"
)

func doPrimes(c pb.CalculatorServiceClient) {
	log.Println("doPrimes was invoked")

	req := &pb.PrimesRequest{
		Number: 210,
	}

	stream, err := c.Primes(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while calling doPrimes: %v\n", err)
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading the stream: %v\n", err)
		}

		log.Printf("Primes: %d\n", msg.Result)
	}
}
