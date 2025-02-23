package main

import (
	"fmt"
	"golang-docker-demo/app/modules/pet"
	"golang-docker-demo/app/server"
	"golang-docker-demo/app/server/generated"
	"golang-docker-demo/config"
	"golang-docker-demo/db"
	"golang-docker-demo/db/generated/query"
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
	config := config.NewConfig()
	log.Println(config)
	db := db.NewLocalDb(config.DB)
	fmt.Println(db)

	query := query.Use(db)
	api := server.NewApi(pet.NewUseCase(query))
	generated.RegisterHandlers(app, api)

	//app.Get("/", func(c *fiber.Ctx) error {
	//	return c.SendString("Hello, Fiber1!")
	//})

	log.Printf("Staring Server")
	log.Fatal(app.Listen(fmt.Sprintf(":%s", port)))
}
