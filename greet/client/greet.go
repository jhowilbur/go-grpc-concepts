package main

import (
	"context"
	pb "github.com/jhowilbur/grpc-project/greet/proto"
	"log"
)

func doGreet(connection pb.GreetServiceClient) {
	log.Println("doGreet was invoked")

	result, err := connection.Greet(context.TODO(), &pb.GreetRequest{
		FirstName: "Wilbur",
	})
	if err != nil {
		log.Fatalf("Could not greet: %v\n", err)
	}

	log.Printf("Greeting: %s\n", result.Result)
}
