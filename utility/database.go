package utility

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init() {
	LoadEnv()
}

var DB *gorm.DB

func ConnectToDB() {
	var err error

	dsn := os.Getenv("DB")

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err == nil {
		log.Println("Connected to DB...")
	} else {
		log.Fatal("Connection to DB lost")
	}
}
