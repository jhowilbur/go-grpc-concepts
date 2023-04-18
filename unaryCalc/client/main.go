package main

import (
	"context"
	proto "github.com/jhowilbur/grpc-project/unaryCalc/proto/sum"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
)

func main() {
	clientCert, err := credentials.NewClientTLSFromFile("unaryCalc/cert/server.crt", "")
	if err != nil {
		log.Fatalln("failed to create cert", err)
	}

	ctx := context.TODO()
	reminderConn, err := grpc.DialContext(ctx, "127.0.0.1:5000",
		grpc.WithTransportCredentials(clientCert))

	if err != nil {
		log.Fatalln("Failed to dial server: ", err)
	}

	// Create a client
	sumClient := proto.NewSumServiceClient(reminderConn)

	// Send a request
	response, err := sumClient.Sum(ctx, &proto.SumRequest{NumberOne: 10, NumberTwo: 3})
	if err != nil {
		log.Fatalln("Failed to send request: ", err)
	}
	log.Printf("Response from server: %v", response)
}
