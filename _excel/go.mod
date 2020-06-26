module source/golang-practical-for-reals/_excel

go 1.13

replace rpc => ../_protobuf

replace _excel/student => ./student

require (
	_excel/student v0.0.0
	github.com/streadway/amqp v1.0.0
	google.golang.org/grpc v1.30.0
	rpc v0.0.0
)
