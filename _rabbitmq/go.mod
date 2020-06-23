module source/golang-practical-for-reals/_rabbitmq

go 1.13

replace rpc => ../_protobuf

require (
	github.com/360EntSecGroup-Skylar/excelize v1.4.1
	github.com/360EntSecGroup-Skylar/excelize/v2 v2.2.0 // indirect
	github.com/streadway/amqp v1.0.0
	google.golang.org/grpc v1.29.1
	rpc v0.0.0-00010101000000-000000000000
)
