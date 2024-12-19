package server

import (
	"golang-docker-demo/app/modules/pet"
	"golang-docker-demo/app/server/generated"

	"github.com/gofiber/fiber/v2"
)

var _ generated.ServerInterface = (*Api)(nil)

type Api struct {
	pet pet.UseCaseInterface
}

func NewApi(pet pet.UseCaseInterface) *Api {
	return &Api{pet: pet}
}

func (a Api) FindPets(c *fiber.Ctx, params generated.FindPetsParams) error {
	return a.pet.FindPets(c, params)
}

func (a Api) AddPet(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (a Api) DeletePet(c *fiber.Ctx, id int64) error {
	//TODO implement me
	panic("implement me")
}

func (a Api) FindPetByID(c *fiber.Ctx, id int64) error {
	//TODO implement me
	panic("implement me")
}
