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
	pb.UnimplementedStudentServer
}

func (s *server) SearchStudent(ctx context.Context, r *pb.StudentRequest) (*pb.StudentResponse, error) {
	ids := Search(r.Keyword)
	data := pb.StudentResponse{
		Ids: ids,
	}
	return &data, nil
}

func (s *server) IndexStudent(ctx context.Context, r *pb.IndexStudentRequest) (*pb.IndexStudentResponse, error) {
	CreateIndex(r)
	data := pb.IndexStudentResponse{
		Message: r.Name + fmt.Sprint(r.Age),
	}
	return &data, nil
}

func init() {
	InitElastic()
}

func main() {
	list, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Print(err.Error())
		return
	}

	s := grpc.NewServer()
	pb.RegisterStudentServer(s, &server{})

	log.Print("listen at: 8081")
	s.Serve(list)
}
