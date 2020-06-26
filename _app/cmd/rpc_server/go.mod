module source/golang-practical-for-reals/_app/rpc_server

go 1.13

replace rpc => ../../../_protobuf

require (
	google.golang.org/grpc v1.30.0
	rpc v0.0.0-00010101000000-000000000000
)
