package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

var addr string = "localhost:5000"

func main() {
	connection, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect on: %v\n", err)
	}
	log.Printf("Connected on %s\n", addr)
	defer connection.Close()

	//..
}
