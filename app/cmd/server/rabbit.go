package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

var (
	rabbitConn *amqp.Connection
)

func failOnError(err error, message string) bool {
	if err != nil {
		fmt.Println("[x] " + message)
		fmt.Println("[detail] " + err.Error())
		return true
	}
	return false
}

func ConnectRabbit() error {
	var err error
	rabbitConn, err = amqp.Dial("amqp://guest:guest@localhost:5672/")
	if failOnError(err, "fail to connect to rabbit") {
		return err
	}
	return nil
}
