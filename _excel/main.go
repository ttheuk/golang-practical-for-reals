package main

import (
	"_excel/excel"
	"encoding/json"
	"log"
	pb "rpc"

	"github.com/streadway/amqp"
)

var excelService *excel.Service

func failOnError(err error, msg string) bool {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		return true
	}
	return false
}

func init() {
	excelRepo := excel.NewExcelRepository()
	excelService = excel.NewService(excelRepo)
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
