package services

import (
	"time"

	"github.com/aman1117/backend/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

// RequestLoggerMiddleware logs all incoming HTTP requests with timing and trace_id
func RequestLoggerMiddleware(c *fiber.Ctx) error {
	start := time.Now()

	// Generate trace_id for this request
	traceID := uuid.New().String()[:8] // Short 8-char trace ID
	c.Locals("trace_id", traceID)

	// Process request
	err := c.Next()

	// Calculate duration
	duration := time.Since(start)

	// Get status code
	status := c.Response().StatusCode()

	// Get user_id if available
	userID, _ := c.Locals("user_id").(uint)

	// Base fields for all requests
	fields := []zap.Field{
		zap.String("trace_id", traceID),
		zap.String("method", c.Method()),
		zap.String("path", c.Path()),
		zap.Int("status", status),
		zap.Duration("duration", duration),
	}

	if userID > 0 {
		fields = append(fields, zap.Uint("user_id", userID))
	}

	// Determine log level based on status
	if status >= 500 {
		fields = append(fields, zap.String("ip", c.IP()), zap.String("user_agent", c.Get("User-Agent")))
		utils.Log.Error("HTTP Request", fields...)
	} else if status >= 400 {
		fields = append(fields, zap.String("ip", c.IP()))
		utils.Log.Warn("HTTP Request", fields...)
	} else {
		utils.Log.Info("HTTP Request", fields...)
	}

	return err
}
