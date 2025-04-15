package pet

import (
	"golang-docker-demo/app/server/generated"
	"golang-docker-demo/db/generated/model"
	"golang-docker-demo/db/generated/query"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/samber/lo"
)

type UseCase struct {
	query *query.Query
	repo  RepositoryInterface
}

func NewUseCase(query *query.Query, repo RepositoryInterface) UseCaseInterface {
	return &UseCase{query: query, repo: repo}
}

func (u UseCase) FindPetsByStatus(c *fiber.Ctx, params generated.FindPetsByStatusParams) error {
	// return c.JSON(fiber.Map{
	// 	"success": true,
	// 	"message": "success",
	// })

	pets, err := u.repo.FindPetByStatus(c.UserContext(), u.query, params)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	res := make([]generated.Pet, len(pets))
	for i, v := range pets {
		res[i] = mPetToPet(v)
	}

	return c.JSON(res)
}

func mPetToPet(m *model.MPet) (g generated.Pet) {
	return generated.Pet{
		Category:  nil,
		Id:        nil,
		Name:      m.Name,
		PhotoUrls: nil,
		Status:    (*generated.PetStatus)(&m.Status),
		Tags:      nil,
	}
}

// create function insert new pet by send model to repository
func (u UseCase) AddPet(c *fiber.Ctx) error {
	pet := new(generated.Pet)
	if err := c.BodyParser(pet); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	mPet := &model.MPet{
		ID:        lo.ToPtr(uuid.New().String()),
		CreatedBy: uuid.New().String(),
		UpdatedBy: uuid.New().String(),
		Name:      pet.Name,
		Status:    string(*pet.Status),
	}

	res, err := u.repo.AddPet(c.UserContext(), u.query, mPet)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"success": true,
		"message": "success",
		"data":    res,
	})
}
