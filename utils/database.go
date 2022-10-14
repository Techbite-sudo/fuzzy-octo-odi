package utils

import (
	"log"
	"odibet/models"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	log.Print("Initialising Database...")
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		DB, err = gorm.Open(sqlite.Open("odi.db"), &gorm.Config{})
	} else {
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	}

	if err != nil {
		log.Panic(err)
	}

	log.Println("Successfully Connected")

	setupModels(
		&models.User{},
	)
}

func setupModels(models ...interface{}) {
	err := DB.AutoMigrate(models...)
	if err != nil {
		panic(err)
	}
}
