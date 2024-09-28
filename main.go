package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Đọc cổng từ biến môi trường, nếu không có thì dùng cổng mặc định là 3000
	port := os.Getenv("PORT")
	if port == "" {
		port = "8888"
	}

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Fiber!")
	})

	log.Fatal(app.Listen(fmt.Sprintf(":%s", port)))
}
