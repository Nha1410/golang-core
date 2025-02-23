package pet

import (
	"golang-docker-demo/app/server/generated"

	"github.com/gofiber/fiber/v2"
)

type UseCaseInterface interface {
	FindPetsByStatus(c *fiber.Ctx, params generated.FindPetsByStatusParams) error
}

type RepositoryInterface interface{}
