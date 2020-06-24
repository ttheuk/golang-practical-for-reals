package student

import (
	"entity"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type StudentRepository struct {
	db *gorm.DB
}

func NewStudentRepository(connection *gorm.DB) *StudentRepository {
	return &StudentRepository{db: connection}
}

func (s *StudentRepository) GetAllById(ids []uint64) (*entity.ListStudent, error) {
	var list entity.ListStudent
	err := s.db.Where("id IN (?)", ids).Find(&list).Error
	return &list, err
}

func (s *StudentRepository) Create(student *entity.Student) error {
	return s.db.Create(student).Error
}

func (s *StudentRepository) Update(student *entity.Student) error {
	return s.db.Save(student).Error
}

func (s *StudentRepository) Delete(id uint64) error {
	return s.db.Where("id = ?", id).Delete(&entity.Student{}).Error
}

func (s *StudentRepository) ExportXLSX() (*entity.ListStudent, error) {
	var list entity.ListStudent
	if err := s.db.Find(&list).Error; err != nil {
		return nil, err
	}
	return &list, nil
}
