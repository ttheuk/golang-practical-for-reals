package main

import (
	"entity"
	pb "rpc"
)

func FindAllStudent() *entity.ListStudent {
	list, err := studentService.FindAll()
	if err != nil {
		return nil
	}
	return list
}

func Create(r *pb.StudentStruct) (*entity.Student, error) {
	student := entity.Student{
		Name: r.Name,
		Age:  int(r.Age),
	}
	err := studentService.Create(&student)
	if err != nil {
		return nil, err
	}
	return &student, err
}

func GetAllById(r *pb.StudentResponse) (*entity.ListStudent, error) {
	list, err := studentService.GetAllById(r.Ids)
	if err != nil {
		return nil, err
	}
	return list, nil
}
