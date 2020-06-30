module source/golang-practical-for-reals/app/cmd/server

go 1.13

replace handler => ./../../handler

replace repository/student => ../../repository/student

replace rpc => ../../../protobuf

replace entity => ../../entity

require (
	entity v0.0.0
	github.com/gin-gonic/gin v1.6.3
	github.com/jinzhu/gorm v1.9.14
	github.com/streadway/amqp v1.0.0
	google.golang.org/grpc v1.29.1
	handler v0.0.0
	repository/student v0.0.0
	rpc v0.0.0
)
