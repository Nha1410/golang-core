package pet

import (
	"context"
	"golang-docker-demo/app/server/generated"
	"golang-docker-demo/db/generated/model"
	"golang-docker-demo/db/generated/query"

	"github.com/gofiber/fiber/v2"
)

type UseCaseInterface interface {
	FindPetsByStatus(c *fiber.Ctx, params generated.FindPetsByStatusParams) error
	AddPet(c *fiber.Ctx) error
}

type RepositoryInterface interface {
	FindPetByStatus(ctx context.Context, tx *query.Query, params generated.FindPetsByStatusParams) (pets []*model.MPet, err error)
	AddPet(ctx context.Context, tx *query.Query, body *model.MPet) (pet *model.MPet, err error)
}
