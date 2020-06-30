module repository/app/student

go 1.13

replace entity => ../../entity

require (
	entity v0.0.0
	github.com/jinzhu/gorm v1.9.14
)
