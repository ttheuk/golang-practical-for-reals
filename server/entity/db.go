package DB

import (
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	host     = "127.0.0.1"
	port     = "5432"
	user     = "postgres"
	dbname   = "demo_db"
	password = "admin"
	sslmode  = "disable"
)

// Thay thế Model của gorm vì cần thêm tag
type MyModel struct {
	ID        uint       `json:"id,omitempty" gorm:"primary_key"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" sql:"index"`
}

func Migrate() {
	db, _ := ConnectDB()
	db.AutoMigrate(
		&Student{},
	)
}

func ConnectDB() (*gorm.DB, error) {
	// tạo kết nối đến postgreSQL
	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", host, port, user, dbname, password, sslmode)
	db, err := gorm.Open("postgres", connectionString)

	if err != nil {
		log.Print(err)
		return nil, err
	}

	return db, nil
}
