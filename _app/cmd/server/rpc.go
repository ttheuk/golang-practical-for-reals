package main

import (
	"google.golang.org/grpc"
)

var (
	rpc        *grpc.ClientConn
	elasticAdd = "localhost:8081"
)

func ConnectElasticRPC() error {
	var err error
	rpc, err = grpc.Dial(elasticAdd, grpc.WithInsecure(), grpc.WithBlock())
	return err
}
