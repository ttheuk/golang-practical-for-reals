package student

import (
	pb "rpc"

	"google.golang.org/grpc"
)

type Service struct {
	repo *ExcelRepository
	conn *grpc.ClientConn
}

func NewService(repo *ExcelRepository, conn *grpc.ClientConn) *Service {
	return &Service{
		repo: repo,
		conn: conn,
	}
}

func (s *Service) ExportXLSX(r *pb.XlsxRequest) error {
	return s.repo.ExportXLSX(r, s.conn)
}
