package student

import (
	pb "rpc"

	"google.golang.org/grpc"
)

type excelRepository interface {
	ExportXLSX(r *pb.XlsxRequest, conn *grpc.ClientConn)
}
