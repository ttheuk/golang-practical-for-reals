package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func init() {
	if err := ConnectRPC(); err != nil {
		log.Print(err)
		return
	}
	fmt.Println("=> connect to rpc servers: successful")

	if err := ConnectRabbit(); err != nil {
		log.Print(err)
		return
	}
	fmt.Println("=> connect to rabbitmq server: successful")
}

func main() {
	r := gin.Default()
	r = NewStudentHandler(r)

	r.Run()
}
