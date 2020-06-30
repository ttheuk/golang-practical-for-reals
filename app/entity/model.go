package entity

import "time"

// Thay thế Model của gorm vì cần thêm tag
type MyModel struct {
	ID        uint       `json:"id,omitempty" gorm:"primary_key"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" sql:"index"`
}
