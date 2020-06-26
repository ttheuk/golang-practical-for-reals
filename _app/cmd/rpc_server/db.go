package main

import (
	"fmt"
	"log"
	"repository/student"

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

var (
	db             *gorm.DB
	studentService *student.Service
)

func failOnError(err error, message string) bool {
	if err != nil {
		fmt.Println("[x] " + message)
		fmt.Println("[detail] " + err.Error())
		return true
	}
	return false
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
