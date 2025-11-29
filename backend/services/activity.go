package services

import (
	"github.com/aman1117/backend/models"
	"github.com/aman1117/backend/utils"
	"github.com/gofiber/fiber/v2"
)

type ActivityRequest struct {
	Username string              `json:"username"`
	Activity models.ActivityName `json:"activity"`
	Hours    float32             `json:"hours"`
}

func UpdateActivityHandler(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Activity updated successfully.",
	})
}

func CreateActivityHandler(c *fiber.Ctx) error {
	var body ActivityRequest
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   "Invalid request body",
		})
	}

	if body.Username == "" || body.Activity == "" || body.Hours <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   "Invalid request body",
		})
	}

	if body.Username != c.Locals("username").(string) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   "you are not authorized to create activity for this username",
		})
	}

	if !models.ActivityName(body.Activity).IsValid() {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   "Invalid activity name",
		})
	}

	db := utils.GetDB()
	// find all the records for this person for current day.
	todayActivities := []models.Activity{}
	result := db.Where("user_id = ? AND date(created_at) = date('now')", c.Locals("user_id").(uint)).Find(&todayActivities)
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   "Failed to find activities",
		})
	}
	// if already present, update the duration_hours with the same activityName or create a new record
	if len(todayActivities) > 0 {
		for _, activity := range todayActivities {
			if activity.Name == models.ActivityName(body.Activity) {
				activity.DurationHours = body.Hours
				result = db.Save(&activity)
				if result.Error != nil {
					return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
						"success": false,
						"error":   "Failed to update activity",
					})
				}
				break
			}
		}
		if len(todayActivities) == 0 {
			activity := models.Activity{
				UserID:        c.Locals("user_id").(uint),
				Name:          models.ActivityName(body.Activity),
				DurationHours: body.Hours,
			}
			result = db.Create(&activity)
		}
		if result.Error != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"error":   "Failed to create activity",
			})
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Activity created successfully",
	})
}
