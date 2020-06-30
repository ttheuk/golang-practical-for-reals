module source/golang-practical-for-reals/app/rpc_server

go 1.13

replace rpc => ../../../protobuf

replace repository/student => ../../repository/student

replace entity => ../../entity

require (
	entity v0.0.0
	github.com/jinzhu/gorm v1.9.14
	google.golang.org/grpc v1.30.0
	repository/student v0.0.0
	rpc v0.0.0-00010101000000-000000000000
)
