module source/my-clean-arch/handler

go 1.13

replace rpc => ../_protobuf

replace entity => ./../entity

replace repository/student => ./../repository/student

require (
	entity v0.0.0
	github.com/gin-gonic/gin v1.6.3
	github.com/jinzhu/gorm v1.9.12 // indirect
	google.golang.org/grpc v1.29.1
	repository/student v0.0.0
	rpc v0.0.0-00010101000000-000000000000
)
