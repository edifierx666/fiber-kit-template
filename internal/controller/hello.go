package controller

import "github.com/gofiber/fiber/v2"

func Hello(c *fiber.Ctx) error {
  return c.SendString("hello fiber kit template")
}
