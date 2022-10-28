package config

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	user     = "postgres"
	password = "presiden123"
	host     = "localhost"
	port     = "5432"
	dbName   = "finaltask"
	db       *gorm.DB
	err      error
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbName, port)
	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal("Error starting database:", err)
	}

	// db.Debug().AutoMigrate(&model.User{}, &model.Comment{}, &model.Photo{}, &model.SocialMedia{})
}

func GetDB() *gorm.DB {
	return db
}
