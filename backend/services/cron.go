package services

import (
	"context"
	"fmt"
	"time"

	"github.com/aman1117/backend/models"
	"github.com/aman1117/backend/utils"
)

func RunDailyJob(ctx context.Context) error {
	db := utils.GetDB().WithContext(ctx)

	var users []models.User
	if err := db.Find(&users).Error; err != nil {
		return err
	}

	// Compute today's date in IST once
	loc, _ := time.LoadLocation("Asia/Kolkata")
	nowIST := time.Now().In(loc)
	todayIST := time.Date(
		nowIST.Year(),
		nowIST.Month(),
		nowIST.Day(),
		0, 0, 0, 0,
		loc,
	)
	fmt.Println(todayIST)
	for _, user := range users {
		if err := AddStreak(user.ID, todayIST, true); err != nil {
			return err
		}
	}

	return nil
}

func CronJob(ctx context.Context) error {
	return RunDailyJob(ctx)
}
