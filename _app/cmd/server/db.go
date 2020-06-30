package main

import (
	"entity"
	"errors"
	"fmt"

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

func failOnError(err error, message string) bool {
	if err != nil {
		fmt.Println("[x] " + message)
		fmt.Println("[detail] " + err.Error())
		return true
	}
	return false
}

func Migrate() bool {
	err := db.DB().Ping()
	if failOnError(err, "no database connection") {
		return false
	}

	db.AutoMigrate(
		&entity.Student{},
	)
	return true
}

func ConnectDB() bool {
	// tạo kết nối đến postgreSQL
	var err error
	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", host, port, user, dbname, password, sslmode)
	db, err = gorm.Open("postgres", connectionString)
	return failOnError(err, "fail to connect to database")
}

func SetupDB() error {
	if !ConnectDB() || !Migrate() {
		return errors.New("[x] setup server failed")
	}
	fmt.Println("[*] connect DB: done")
	fmt.Println("[*] migrate DB: done")

	return nil
}
