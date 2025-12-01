package services

import (
	"time"

	"github.com/aman1117/backend/models"
	"github.com/aman1117/backend/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type GetStreakRequest struct {
	Username string `json:"username"`
	Date     string `json:"date"`
}

type StreakDTO struct {
	ID      uint   `json:"id"`
	Current int    `json:"current"`
	Longest int    `json:"longest"`
	Date    string `json:"date"`
}

func AddStreak(userID uint, date time.Time) error {
	today := time.Now().Truncate(24 * time.Hour)
	if date.Before(today) {
		return nil
	}
	db := utils.GetDB()
	streak := models.Streak{}
	result := db.Where("user_id = ?", userID).Order("activity_date DESC").First(&streak)

	if result.Error != nil {
		if result.Error != gorm.ErrRecordNotFound {
			return result.Error
		}
	}
	if result.RowsAffected > 0 {
		streakToInsert := models.Streak{}
		if streak.ActivityDate == date.Add(-24*time.Hour) {
			streakToInsert.Current = streak.Current + 1
		} else if streak.ActivityDate == date {
			if streakToInsert.Current == 1 {
				return nil
			}
			streakToInsert.Current = 1
		}
		streakToInsert.ActivityDate = date
		streakToInsert.Longest = max(streak.Longest, streakToInsert.Current)
		streakToInsert.UserID = userID
		if streak.ActivityDate == date {
			if err := db.Save(&streakToInsert).Error; err != nil {
				return err
			}
		} else {
			if err := db.Create(&streakToInsert).Error; err != nil {
				return err
			}
		}
	} else {
		streak := models.Streak{
			UserID:       userID,
			Current:      1,
			Longest:      1,
			ActivityDate: date,
		}
		if err := db.Create(&streak).Error; err != nil {
			return err
		}
	}
	return nil
}

func GetStreakHandler(c *fiber.Ctx) error {
	var body GetStreakRequest
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   "Invalid request body",
		})
	}

	const layout = "2006-01-02"
	date, err := time.Parse(layout, body.Date)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   "Invalid date format, use YYYY-MM-DD",
		})
	}
	date = date.Truncate(24 * time.Hour)
	db := utils.GetDB()
	streak := models.Streak{}
	user := models.User{}
	result := db.Where("username = ?", body.Username).First(&user)
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   "Failed to find user",
		})
	}
	result = db.Where("user_id = ? AND activity_date = ?", user.ID, date).Last(&streak)
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   "Failed to find streak",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    ToStreakDTOs(streak),
	})
}

func ToStreakDTOs(in models.Streak) StreakDTO {
	return StreakDTO{
		ID:      in.ID,
		Current: in.Current,
		Longest: in.Longest,
		Date:    in.ActivityDate.Format("2006-01-02"),
	}
}
