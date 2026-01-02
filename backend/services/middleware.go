package services

import (
	"time"

	"github.com/aman1117/backend/utils"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

// RequestLoggerMiddleware logs all incoming HTTP requests with timing
func RequestLoggerMiddleware(c *fiber.Ctx) error {
	start := time.Now()

	// Process request
	err := c.Next()

	// Calculate duration
	duration := time.Since(start)

	// Get status code
	status := c.Response().StatusCode()

	// Determine log level based on status
	if status >= 500 {
		utils.Log.Error("HTTP Request",
			zap.String("method", c.Method()),
			zap.String("path", c.Path()),
			zap.Int("status", status),
			zap.Duration("duration", duration),
			zap.String("ip", c.IP()),
			zap.String("user_agent", c.Get("User-Agent")),
		)
	} else if status >= 400 {
		utils.Log.Warn("HTTP Request",
			zap.String("method", c.Method()),
			zap.String("path", c.Path()),
			zap.Int("status", status),
			zap.Duration("duration", duration),
			zap.String("ip", c.IP()),
		)
	} else {
		utils.Log.Info("HTTP Request",
			zap.String("method", c.Method()),
			zap.String("path", c.Path()),
			zap.Int("status", status),
			zap.Duration("duration", duration),
		)
	}

	return err
}
