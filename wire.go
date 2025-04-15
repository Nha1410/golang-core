//go:build wireinject
// +build wireinject

package main

import (
	"golang-docker-demo/config"

	"github.com/google/wire"
)

func Ready() *DevServer {
	wire.Build(
		NewFiber,         // *fiber.App
		config.NewConfig, // *config.Config
		// db.NewLocalDb,    // (*gorm.DB), depends on *config.Config

		// query.Use,         // *query.Query
		// pet.NewRepository, // pet.Repository
		// pet.NewUseCase,    // pet.UseCaseInterface
		// server.NewApi,     // *server.Api
		NewDevServer, // *DevServer
	)
	return &DevServer{}
}
