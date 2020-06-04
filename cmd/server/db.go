package main

import (
	"entity"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
)

const (
	host     = "127.0.0.1"
	port     = "5432"
	user     = "postgres"
	dbname   = "demo_db"
	password = "admin"
	sslmode  = "disable"
)

var db *gorm.DB

func Migrate() error {
	if err := db.DB().Ping(); err != nil {
		return err
	}

	db.AutoMigrate(
		&entity.Student{},
	)

	return nil
}

func ConnectDB() error {
	// tạo kết nối đến postgreSQL
	var err error
	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", host, port, user, dbname, password, sslmode)
	db, err = gorm.Open("postgres", connectionString)

	if err != nil {
		log.Print(err)
		return err
	}

	return nil
}

func SetupServer() error {
	err := ConnectDB()
	// defer db.Close()

	if err != nil {
		return err
	}

	err = Migrate()
	if err != nil {
		return err
	}

	err = ConnectRPC()
	if err != nil {
		return err
	}

	return nil
}
