package pet

import (
	"github.com/gofiber/fiber/v2"
	"golang-docker-demo/app/server/generated"
)

type UseCaseInterface interface {
	FindPets(c *fiber.Ctx, params generated.FindPetsParams) error
}

type RepositoryInterface interface{}
