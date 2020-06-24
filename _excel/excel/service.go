package excel

import pb "rpc"

type Service struct {
	repo *ExcelRepository
}

func NewService(repo *ExcelRepository) *Service {
	return &Service{repo: repo}
}

func (s *Service) ExportXLSX(r *pb.XlsxRequest) error {
	return s.repo.ExportXLSX(r)
}
