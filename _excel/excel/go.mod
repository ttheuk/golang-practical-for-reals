module _excel/excel

replace rpc => ./../../_protobuf

replace entity => ./../../_main_server/entity

go 1.13

require (
	entity v0.0.0
	github.com/360EntSecGroup-Skylar/excelize v1.4.1
	github.com/jinzhu/gorm v1.9.14
	rpc v0.0.0
)
