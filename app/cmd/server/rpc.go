package main

import (
	"google.golang.org/grpc"
)

var (
	elasticConn       *grpc.ClientConn
	localConn         *grpc.ClientConn
	elasticRPCAddress = "localhost:8081"
	localRPCAdress    = "localhost:5000"
)

func ConnectRPC() error {
	var err error
	elasticConn, err = grpc.Dial(elasticRPCAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return err
	}

	var err2 error
	localConn, err2 = grpc.Dial(localRPCAdress, grpc.WithInsecure(), grpc.WithBlock())
	if err2 != nil {
		return err2
	}

	return nil
}
