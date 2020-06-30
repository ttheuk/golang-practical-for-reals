package main

import (
	"handler"

	"github.com/gin-gonic/gin"
)

func NewStudentHandler(e *gin.Engine) *gin.Engine {
	h := handler.NewStudentHandler(rabbitConn, elasticConn, localConn)

	e.GET("/students", h.GetStudents)
	e.GET("/students/xlsx", h.ExportXLSX)
	e.POST("/students", h.CreateStudent)
	// e.DELETE("/students", h.Delete)

	return e
}
