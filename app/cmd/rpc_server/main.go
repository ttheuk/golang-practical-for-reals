package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"repository/student"
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
		}
		data.Students = append(data.Students, &student)
	}

	return &data, nil
}

func (s *server) Create(ctx context.Context, r *pb.StudentStruct) (*pb.StudentStruct, error) {
	data, err := Create(r)
	if failOnError(err, "fail to create student") {
		return nil, err
	}

	r.Id = int64(data.ID)
	r.CreatedAt = data.CreatedAt.UnixNano()
	r.CreatedAt = data.UpdatedAt.UnixNano()
	return r, nil
}

func (s *server) GetAllById(ctx context.Context, r *pb.StudentResponse) (*pb.AllStudentResponse, error) {
	list, err := GetAllById(r)
	if failOnError(err, "fail to get student by list of ids") {
		return nil, err
	}

	data := pb.AllStudentResponse{}
	for _, obj := range *list {
		s := pb.AllStudentResponse_StudentStruct{
			Id:        int64(obj.ID),
			Name:      obj.Name,
			Age:       int32(obj.Age),
			CreatedAt: obj.CreatedAt.UnixNano(),
			UpdatedAt: obj.UpdatedAt.UnixNano(),
		}
		data.Students = append(data.Students, &s)
	}

	return &data, nil
}

func init() {
	if err := SetupDB(); err != nil {
		log.Print(err)
		return
	}
	fmt.Println("=> connect to database: successful")
}

func main() {
	repo := student.NewStudentRepository(db)
	studentService = student.NewService(*repo)

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
