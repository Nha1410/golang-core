package server

import (
	"golang-docker-demo/app/server/generated"

	"github.com/gofiber/fiber/v2"
)

var _ generated.ServerInterface = (*Api)(nil)

type Api struct {
}

func NewApi() *Api {
	return &Api{}
}

func (a Api) FindPets(c *fiber.Ctx, params generated.FindPetsParams) error {
	//TODO implement me
	return c.Status(200).JSON(
		"123",
	)
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
