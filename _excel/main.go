package main

import (
	"_excel/student"
	"encoding/json"
	"fmt"
	"log"
	pb "rpc"

	"github.com/streadway/amqp"
	"google.golang.org/grpc"
)

var (
	excelService *student.Service
	rpc          *grpc.ClientConn
	dataAdd      = "localhost:5000"
)

func ConnectRPC() error {
	var err error
	rpc, err = grpc.Dial(dataAdd, grpc.WithInsecure(), grpc.WithBlock())
	return err
}

func failOnError(err error, msg string) bool {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		return true
	}
	return false
}

func init() {
	if err := ConnectRPC(); err != nil {
		fmt.Println("[x] (_excel/main.go:36) fail to connect to database RPC server /n[detail] " + err.Error())
		return
	}

	excelRepo := student.NewExcelRepository()
	excelService = student.NewService(excelRepo, rpc)
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare("logs", "fanout", true, false, false, false, nil)
	failOnError(err, "Failed to declare an exchange")

	q, err := ch.QueueDeclare("", false, false, true, false, nil)
	failOnError(err, "Failed to declare a queue")

	err = ch.QueueBind(q.Name, "", "logs", false, nil)
	failOnError(err, "Failed to bind a queue")

	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	failOnError(err, "Failed to register a consumer")

	go func() {
		for d := range msgs {
			r := pb.XlsxRequest{}
			err := json.Unmarshal(d.Body, &r)
			if failOnError(err, "[x] Unmashal fail") {
				return
			}

			err = excelService.ExportXLSX(&r)
			failOnError(err, "[x] Fail to export file at: "+r.Path+"/"+r.FileName)
		}
	}()

	forever := make(chan bool)
	log.Printf("[*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
