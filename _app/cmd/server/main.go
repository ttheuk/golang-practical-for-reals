package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func init() {
	if err := SetupDB(); err != nil {
		log.Print(err)
		return
	}
	fmt.Println("=> connect to database: successful")

	if err := ConnectElasticRPC(); err != nil {
		log.Print(err)
		return
	}
	fmt.Println("=> connect to elastic server: successful")
}

func main() {
	r := gin.Default()
	r = NewStudentHandler(r)

	r.Run()
}
