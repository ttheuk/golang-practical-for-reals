package handler

import (
	"google.golang.org/grpc"
)

var (
	address = "localhost:8081"
)

func ConnectRPC() (*grpc.ClientConn, error) {
	return grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
}
