module source/golang-practical-for-reals/_excel/student

replace rpc => ./../../_protobuf

go 1.13

require (
	github.com/360EntSecGroup-Skylar/excelize v1.4.1
	github.com/jinzhu/gorm v1.9.14
	google.golang.org/grpc v1.29.1
	rpc v0.0.0
)
