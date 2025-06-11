//go:build wireinject
// +build wireinject

package server

import (
	"golang-docker-demo/app/modules/pet"
	"golang-docker-demo/config"
	"golang-docker-demo/db"

	"github.com/google/wire"
)

func Ready() *DevServer {
	wire.Build(
		NewFiber,         // *fiber.App
		config.NewConfig, // *config.Config
		// db.NewLocalDb,    // *gorm.DB
		db.NewLocalQuery, // *query.Query
		pet.NewRepository,
		pet.NewUseCase,
		NewApi,
		NewDevServer, // *DevServer (now includes *server.Api)
	)
	return &DevServer{}
}
