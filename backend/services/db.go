package services

import (
	"github.com/aman1117/backend/utils"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(email, username, password string) error {
	db := utils.GetDB()
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	result := db.Exec("INSERT INTO users (email, username, password_hash) VALUES ($1, $2, $3)", email, username, string(hash))
	return result.Error
}
