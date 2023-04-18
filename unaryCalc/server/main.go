package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"

	proto "github.com/jhowilbur/grpc-project/unaryCalc/proto/sum"
)

func main() {

	socketListen, err := net.Listen("tcp", "localhost:5000")
	if err != nil {
		log.Fatalf("Error starting tcp listener on port 5000: %v\n", err)
	}
	log.Println("Running at port 5000")

	// TLS
	serverCert, err := credentials.NewServerTLSFromFile("unaryCalc/cert/server.crt", "unaryCalc/cert/server.key")
	if err != nil {
		log.Fatalln("failed to create cert", err)
	}
	server := grpc.NewServer(grpc.Creds(serverCert)) // create gRPC server

	sumInstance := NewSum() // new instance for Sum Server

	//reflection.Register(server)                         // register API -> https://github.com/grpc/grpc/blob/master/doc/server-reflection.md
	proto.RegisterSumServiceServer(server, sumInstance) // register it to the gRPC server

	// start serving
	//server.Serve(socketListen)
	if err = server.Serve(socketListen); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
