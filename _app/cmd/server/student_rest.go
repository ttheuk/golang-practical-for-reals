package main

import (
	"handler"
	"repository/student"

	"github.com/gin-gonic/gin"
)

func NewStudentHandler(e *gin.Engine) *gin.Engine {
	repo := student.NewStudentRepository(db)
	h := handler.NewStudentHandler(*repo, rpc)

	e.GET("/students", h.GetStudents)
	e.GET("/students/xlsx", h.ExportXLSX)
	e.POST("/students", h.CreateStudent)
	// e.DELETE("/students", h.Delete)

	return e
}
