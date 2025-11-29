package services

import (
	"errors"

	"github.com/aman1117/backend/models"
	"github.com/aman1117/backend/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
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

func GetUserByIdentifier(identifier string) (*models.User, error) {
	db := utils.GetDB()
	user := &models.User{}
	result := db.Where("email = ? OR username = ?", identifier, identifier).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return user, nil
}
