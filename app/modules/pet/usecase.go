package pet

import (
	"golang-docker-demo/app/server/generated"
	"golang-docker-demo/db/generated/model"
	"golang-docker-demo/db/generated/query"

	"github.com/gofiber/fiber/v2"
)

type UseCase struct{}

func NewUseCase(db *query.Query) UseCaseInterface {
	return &UseCase{}
}

func (u UseCase) FindPetsByStatus(c *fiber.Ctx, params generated.FindPetsByStatusParams) error {
	// return c.JSON(fiber.Map{
	// 	"success": true,
	// 	"message": "success",
	// })

	pets := []model.MPet{
		{
			ID:   "1",
			Name: "Cat",
		},
		{
			ID:   "2",
			Name: "Dog",
		},
	}

	res := make([]generated.Pet, len(pets))
	for i, v := range pets {
		res[i] = mPetToPet(v)
	}

	return c.JSON(res)
}

func mPetToPet(m model.MPet) (g generated.Pet) {
	return generated.Pet{
		Category:  nil,
		Id:        nil,
		Name:      m.Name,
		PhotoUrls: nil,
		Status:    nil,
		Tags:      nil,
	}
}
