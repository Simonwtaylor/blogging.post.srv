package config

import (
	"fmt"
	"log"
	"os"
	"post/entities"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

const defaultPort = "8080"

type GormDatabase struct {
	db *gorm.DB
}

func NewGormDatabase() *GormDatabase {
	err := godotenv.Load()

	if err != nil {
		log.Panicf("Error loading .env file, %v", err)
	}

	db, err := gorm.Open("postgres", fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_NAME"), os.Getenv("DB_PASSWORD")))
	if err != nil {
		log.Panicf("unable to open postgres db %v", err)
	}

	db.AutoMigrate(&entities.User{}, &entities.Post{})

	return &GormDatabase{
		db: db,
	}
}
