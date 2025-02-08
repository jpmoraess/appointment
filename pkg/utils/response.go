package utils

import "github.com/gofiber/fiber/v2"

func SendSuccessResponse(c *fiber.Ctx, statusCode int, data interface{}) error {
	return c.Status(statusCode).JSON(fiber.Map{
		"success": true,
		"data":    data,
	})
}

func SendErrorResponse(c *fiber.Ctx, statusCode int, err error) error {
	return c.Status(statusCode).JSON(fiber.Map{
		"success": false,
		"message": err.Error(),
	})
}
