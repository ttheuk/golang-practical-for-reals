package main

import (
	"context"
	"fmt"
	"log"
	"net"
	pb "rpc"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedExcelServer
}

func (s *server) GetAllStudent(ctx context.Context, e *pb.Empty) (*pb.AllStudentResponse, error) {
	fmt.Println("get all student")
	return nil, nil
}

func main() {
	list, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Print(err.Error())
		return
	}

	s := grpc.NewServer()
	pb.RegisterExcelServer(s, &server{})

	log.Print("[*] rpc listen at: 5000")
	s.Serve(list)
}
