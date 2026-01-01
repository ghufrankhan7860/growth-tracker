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

func UpdateUsername(userID uint, newUsername string) error {
	db := utils.GetDB()
	result := db.Model(&models.User{}).Where("id = ?", userID).Update("username", newUsername)
	return result.Error
}

func UpdatePrivacy(userID uint, isPrivate bool) error {
	db := utils.GetDB()
	result := db.Model(&models.User{}).Where("id = ?", userID).Update("is_private", isPrivate)
	return result.Error
}

func GetUserPrivacy(userID uint) (bool, error) {
	db := utils.GetDB()
	var user models.User
	if err := db.Where("id = ?", userID).First(&user).Error; err != nil {
		return false, err
	}
	return user.IsPrivate, nil
}
