package utils

import (
	"log"
	"os"
	"strconv"
	"time"

	"fmt"

	"github.com/aman1117/backend/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Claims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

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

func GenerateToken(user *models.User) (string, time.Time, error) {
	now := time.Now()
	ttl, err := strconv.Atoi(GetFromEnv("TTL_ACCESS_TOKEN"))
	if err != nil {
		return "", time.Time{}, err
	}
	exp := now.Add(time.Duration(ttl) * time.Hour)

	claims := &Claims{
		UserID:   user.ID,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
			IssuedAt:  jwt.NewNumericDate(now),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString([]byte(GetFromEnv("JWT_SECRET")))
	return signed, exp, err
}
