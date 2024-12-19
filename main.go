package main

import (
	"fmt"
	"golang-docker-demo/app/modules/pet"
	"golang-docker-demo/app/server"
	"golang-docker-demo/config"
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
	configConfig := config.NewConfig()
	api := server.NewApi(pet.NewUseCase())
	db := db.
		generated.RegisterHandlers(app, api)

	//app.Get("/", func(c *fiber.Ctx) error {
	//	return c.SendString("Hello, Fiber1!")
	//})

	log.Fatal(app.Listen(fmt.Sprintf(":%s", port)))
}
