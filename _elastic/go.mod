module source/golang-practical-for-reals/_elastic

go 1.13

replace rpc => ../_protobuf

require (
	github.com/olivere/elastic v6.2.31+incompatible
	github.com/olivere/elastic/v7 v7.0.15
	google.golang.org/grpc v1.29.1
	gopkg.in/olivere/elastic.v6 v6.2.31
	rpc v0.0.0-00010101000000-000000000000
)
