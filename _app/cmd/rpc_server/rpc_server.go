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
	list := FindAllStudent()

	data := pb.AllStudentResponse{}
	for _, obj := range *list {
		student := pb.AllStudentResponse_StudentStruct{
			Id:        int64(obj.ID),
			Name:      obj.Name,
			Age:       int32(obj.Age),
			CreatedAt: obj.CreatedAt.UnixNano(),
			UpdatedAt: obj.UpdatedAt.UnixNano(),
			// DeletedAt: obj.DeletedAt.UnixNano(),
		}
		data.Students = append(data.Students, &student)
	}

	return &data, nil
}

func main() {
	if err := ConnectDB(); err != nil {
		fmt.Println("[x] fail to connect db")
		fmt.Println("[detail] " + err.Error())
		return
	}

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
