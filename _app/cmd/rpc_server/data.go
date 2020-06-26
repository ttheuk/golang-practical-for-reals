package main

import (
	"entity"
	"repository/student"
)

func FindAllStudent() *entity.ListStudent {
	repo := student.NewStudentRepository(db)
	studentService = student.NewService(*repo)

	list, err := studentService.FindAll()
	if failOnError(err, "fail on get all student") {
		return nil
	}

	return list
}
