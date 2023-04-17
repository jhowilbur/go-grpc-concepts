package main

import (
	"context"
	pb "github.com/jhowilbur/grpc-project/greet/proto"
	"log"
)

func (server *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Greet funciton was invoked with %v\n", in)
	log.Printf("Greet context was invoked with %v\n", ctx)

	return &pb.GreetResponse{
		Result: "Hello " + in.FirstName,
	}, nil
}
