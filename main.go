package main

import (
	"fmt"
	"os"

	fiber "github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	var PORT = os.Getenv("STCK_MG_PORT")
	os.Getenv("DATABASE_URL")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	fmt.Fprint(os.Stdout, "Server starting on port"+PORT)

	app.Listen(":" + PORT)
}
