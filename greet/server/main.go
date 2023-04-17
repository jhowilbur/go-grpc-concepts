package main

import (
	"google.golang.org/grpc"
	"log"
	"net"

	pb "github.com/jhowilbur/grpc-project/greet/proto"
)

var addr string = "0.0.0.0:5000"

type Server struct {
	pb.GreetServiceServer
}

func main() {
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}
	log.Printf("Listen on %s\n", addr)

	server := grpc.NewServer()
	if err = server.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}

}
