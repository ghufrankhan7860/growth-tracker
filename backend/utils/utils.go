package utils

import (
	"log"
	"os"

	"fmt"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDB() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("err loading: %v", err)
	}
	username := GetFromEnv("DB_USERNAME")
	password := GetFromEnv("DB_PASSWORD")
	host := GetFromEnv("DB_HOST")
	port := GetFromEnv("DB_PORT")
	dbname := GetFromEnv("DB_NAME")
	psql := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=require&channel_binding=require", username, password, host, port, dbname)
	db, err := gorm.Open(postgres.Open(psql), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func GetFromEnv(key string) string {
	return os.Getenv(key)
}
