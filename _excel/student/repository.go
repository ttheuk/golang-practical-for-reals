package student

import pb "rpc"

type excelRepository interface {
	ExportXLSX(r *pb.XlsxRequest)
}
