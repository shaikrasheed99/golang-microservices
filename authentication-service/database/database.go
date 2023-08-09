package database

import (
	"log"

	"github.com/shaikrasheed99/authentication-service/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase() (*gorm.DB, error) {
	log.Println("[InitDatabase] Initiating database")
	dsn := configs.GetDSNString()

	log.Println("[InitDatabase] Opening the database connection")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("[InitDatabase]", err.Error())
		return nil, err
	}

	log.Println("[InitDatabase] Database connection has established")
	return db, nil
}
