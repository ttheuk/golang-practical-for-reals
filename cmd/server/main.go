package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func init() {
	if err := SetupServer(); err != nil {
		log.Print(err)
	}
}

func main() {
	r := gin.Default()

	r = NewStudentHandler(r)

	r.Run()
}
