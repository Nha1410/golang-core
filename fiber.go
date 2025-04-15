package main

import (
	"encoding/json"
	"errors"
	"golang-docker-demo/app/server/generated"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/samber/lo"
)

func NewFiber() *fiber.App {
	settings := fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
		ErrorHandler: func(c *fiber.Ctx, err error) error {

			// Status code defaults to 500
			code := fiber.StatusInternalServerError

			// Retrieve the custom status code if it's a *fiber.Error
			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}

			res := &generated.ApiResponse{
				Code:    lo.ToPtr(int32(code)),
				Type:    lo.ToPtr("error"),
				Message: lo.ToPtr(err.Error()),
			}

			if err := c.Status(code).JSON(res); err != nil {
				// In case the JSON failed
				return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
			}

			// Return from handler
			return nil
		},
	}
	app := fiber.New(settings)
	app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
	}))

	return app

}
