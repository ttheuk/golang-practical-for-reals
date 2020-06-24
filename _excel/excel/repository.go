package excel

import pb "rpc"

type excelRepository interface {
	ExportXLSX(r *pb.XlsxRequest)
}
