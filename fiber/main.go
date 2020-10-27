package main

import (
	"fmt"
	"github.com/gofiber/fiber"
)

func handleHome(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
  }

func main() {
	app := fiber.New()

	app.Get("/", handleHome)
	

	fmt.Println("Listening on port 3000...")
	app.Listen(":3000")
}