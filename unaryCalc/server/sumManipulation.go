package main

import (
	"context"
	proto "github.com/jhowilbur/grpc-project/unaryCalc/proto/sum"
	"log"
)

type Sum struct {
	proto.UnimplementedSumServiceServer
}

func NewSum() *Sum {
	return &Sum{}
}

func (t *Sum) Sum(ctx context.Context, input *proto.SumRequest) (*proto.SumResponse, error) {
	log.Printf(`Sum numberOne %d, numberTwo %d`, input.GetNumberOne(), input.GetNumberTwo())

	response := &proto.SumResponse{
		Result: input.GetNumberOne() + input.GetNumberTwo(),
	}

	return response, nil
}
