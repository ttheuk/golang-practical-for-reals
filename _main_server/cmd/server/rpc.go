package main

import (
	"google.golang.org/grpc"
)

var (
	rpc     *grpc.ClientConn
	address = "localhost:8081"
)

func ConnectRPC() error {
	var err error
	rpc, err = grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	return err
}
