package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func HelloTest (c *fiber.Ctx) error {
	return c.SendString("Hello, Hakimi World!")
}