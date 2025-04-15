package main

import (
	"golang-docker-demo/config"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

type DevServer struct {
	fiber *fiber.App
	cfg   *config.Config
	// api   *server.Api
}

// func NewDevServer(cfg *config.Config, api *server.Api, fiber *fiber.App) *DevServer {
func NewDevServer(fiber *fiber.App, cfg *config.Config) *DevServer {
	return &DevServer{
		fiber: fiber,
		cfg:   cfg,
		// api:   api,
	}
}

func main() {
	// Lấy PORT từ env
	port := os.Getenv("PORT")
	if port == "" {
		port = "8888"
	}

	// app, err := InitializeServer()
	// if err != nil {
	// 	log.Fatalf("failed to initialize server: %v", err)
	// }

	log.Printf("Starting server on port %s", port)
	// log.Fatal(app.Listen(fmt.Sprintf(":%s", port)))
}
