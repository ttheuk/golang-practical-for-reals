package main

import (
	"source/mix/server/handler"

	"github.com/gin-gonic/gin"
)

func init() {
	handler.MigrateDB()
}

func main() {
	r := gin.Default()

	api := r.Group("/api")
	api.GET("/students", handler.GetStudents)
	api.POST("/students", handler.CreateStudent)
	api.POST("/students/1000", handler.Create10000)

	r.Run()
}