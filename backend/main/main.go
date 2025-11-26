package main

import (
	"fmt"
	"log"

	"github.com/aman1117/backend/models"
	"github.com/aman1117/backend/utils"
)

func main() {

	db := utils.GetDB()
	if err := db.Error; err != nil {
		log.Fatalf("DB connection failed: %v", err)
	}

	if err := db.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("AutoMigrate failed: %v", err)
	}
	fmt.Println("DB Migrations Successful")
}
