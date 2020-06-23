package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
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
			switch string(d.Body) {
			case "export_xlsx":
				{
					fmt.Println("ex")
					ExportXLSX(d.Body)
				}
			}
		}
	}()

	forever := make(chan bool)
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func ExportXLSX(body []byte) error {
	// var value string
	// err := json.Unmarshal(body, value)
	// failOnError(err, "fail to unmarshal message")
	// // Tạo client để gọi phương thức từ RPC

	// client := pb.NewStudentClient(h.conn)
	// ctx := context.Background()

	// xlsxRequest := pb.XlsxRequest{
	// 	Students: []*pb.XlsxRequest_Student{&s1, &s2, &s3},
	// 	Path:     path,
	// 	FileName: fileName,
	// }

	// // Gọi hàm tìm kiếm student từ RPC
	// xlsxResponse, err := client.ExportXLSX(ctx, &xlsxRequest)
	// failOnError(c, err)
	return nil
}
