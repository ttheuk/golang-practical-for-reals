package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func init() {
	if err := SetupServer(); err != nil {
		log.Print(err)
		return
	}
	fmt.Println("=> initiate server: successful")
}

func main() {
	r := gin.Default()

	r = NewStudentHandler(r)

	r.Run()
}
