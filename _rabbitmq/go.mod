module source/golang-practical-for-reals/_rabbitmq

go 1.13

replace rpc => ../_protobuf

require (
	github.com/360EntSecGroup-Skylar/excelize v1.4.1
	github.com/streadway/amqp v1.0.0
	rpc v0.0.0-00010101000000-000000000000
)
