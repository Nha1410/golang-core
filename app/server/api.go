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

func (a Api) DeletePet(c *fiber.Ctx, petId int64) error {
	//TODO implement me
	panic("implement me")
}

func (a Api) AddPet(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (a Api) UpdatePet(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (a Api) FindPetsByTags(c *fiber.Ctx, params generated.FindPetsByTagsParams) error {
	//TODO implement me
	panic("implement me")
}

func (a Api) GetPetById(c *fiber.Ctx, petId int64) error {
	//TODO implement me
	panic("implement me")
}

func (a Api) UpdatePetWithForm(c *fiber.Ctx, petId int64) error {
	//TODO implement me
	panic("implement me")
}

func (a Api) UploadFile(c *fiber.Ctx, petId int64) error {
	//TODO implement me
	panic("implement me")
}

func (a Api) GetInventory(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (a Api) PlaceOrder(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (a Api) DeleteOrder(c *fiber.Ctx, orderId int64) error {
	//TODO implement me
	panic("implement me")
}

func (a Api) GetOrderById(c *fiber.Ctx, orderId int64) error {
	//TODO implement me
	panic("implement me")
}

func (a Api) CreateUser(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (a Api) CreateUsersWithArrayInput(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (a Api) CreateUsersWithListInput(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (a Api) LoginUser(c *fiber.Ctx, params generated.LoginUserParams) error {
	//TODO implement me
	panic("implement me")
}

func (a Api) LogoutUser(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (a Api) DeleteUser(c *fiber.Ctx, username string) error {
	//TODO implement me
	panic("implement me")
}

func (a Api) GetUserByName(c *fiber.Ctx, username string) error {
	//TODO implement me
	panic("implement me")
}

func (a Api) UpdateUser(c *fiber.Ctx, username string) error {
	//TODO implement me
	panic("implement me")
}

func NewApi(pet pet.UseCaseInterface) *Api {
	return &Api{pet: pet}
}

func (a Api) FindPetsByStatus(c *fiber.Ctx, params generated.FindPetsByStatusParams) error {
	return a.pet.FindPetsByStatus(c, params)
}
