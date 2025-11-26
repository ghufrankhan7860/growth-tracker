package services

import "github.com/gofiber/fiber/v2"

type registerRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func RegisterHandler(c *fiber.Ctx) error {
	var body registerRequest
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	if body.Email == "" || body.Username == "" || body.Password == "" {
		return fiber.NewError(fiber.StatusBadRequest, "All fields are required")
	}

	if len(body.Password) < 8 {
		return fiber.NewError(fiber.StatusBadRequest, "Password must be at least 8 characters long")
	}

	if err := CreateUser(body.Email, body.Username, body.Password); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Could not create user (maybe email/username already used)")
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User created successfully.",
	})

}
