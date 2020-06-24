module source/golang-practical-for-reals/_excel

go 1.13

replace rpc => ../_protobuf

replace _excel/excel => ./excel

replace entity => ./../_main_server/entity

require (
	_excel/excel v0.0.0
	entity v0.0.0
	github.com/360EntSecGroup-Skylar/excelize v1.4.1
	github.com/jinzhu/gorm v1.9.14
	github.com/streadway/amqp v1.0.0
	google.golang.org/grpc v1.29.1
	rpc v0.0.0
)
