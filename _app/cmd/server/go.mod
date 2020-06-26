module source/golang-practical-for-reals/_app/cmd/server

go 1.13

replace handler => ./../../handler

replace entity => ../../entity

replace repository/student => ../../repository/student

replace rpc => ../../../_protobuf

require (
	entity v0.0.0
	github.com/gin-gonic/gin v1.6.3
	github.com/jinzhu/gorm v1.9.14
	google.golang.org/grpc v1.29.1
	handler v0.0.0
	repository/student v0.0.0
	rpc v0.0.0
)
