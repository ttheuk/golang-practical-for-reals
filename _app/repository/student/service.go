package student

import (
	"entity"
)

type Service struct {
	repo StudentRepository
}

func NewService(repo StudentRepository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetAllById(ids []uint64) (*entity.ListStudent, error) {
	return s.repo.GetAllById(ids)
}

func (s *Service) Update(student *entity.Student) error {
	return s.repo.Update(student)
}

func (s *Service) Create(student *entity.Student) error {
	return s.repo.Create(student)
}

func (s *Service) Delete(id uint64) error {
	return s.repo.Delete(id)
}

func (s *Service) FindAll() (*entity.ListStudent, error) {
	return s.repo.FindAll()
}
