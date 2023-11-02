package configs

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	DB_HOST     string
	DB_PORT     string
	DB_NAME     string
	DB_USER     string
	DB_PASSWORD string
)

func LoadConfigs() error {
	log.Println("[LoadConfigs] Loading env variables")

	err := godotenv.Load()
	if err != nil {
		log.Println("[LoadConfigs]", err.Error())
	}

	DB_HOST = os.Getenv("DB_HOST")
	DB_PORT = os.Getenv("DB_PORT")
	DB_NAME = os.Getenv("DB_NAME")
	DB_USER = os.Getenv("DB_USER")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")

	log.Println("[LoadConfigs] Env variables have loaded")
	return nil
}

func GetDSNString() string {
	log.Println("[GetDSNString] Getting dns string")

	dsn := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v", DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)

	log.Println("[GetDSNString] Returning dns string value")
	return dsn
}
