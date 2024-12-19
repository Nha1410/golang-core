package pet

import (
	"github.com/gofiber/fiber/v2"
	"golang-docker-demo/app/server/generated"
)

type UseCase struct{}

func NewUseCase() UseCaseInterface {
	return &UseCase{}
}

func (u UseCase) FindPets(c *fiber.Ctx, params generated.FindPetsParams) error {
	return c.JSON(fiber.Map{
		"success": true,
		"message": "success",
	})
}
