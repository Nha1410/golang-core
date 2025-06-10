package server

import (
	"golang-docker-demo/app/server/generated"
	"golang-docker-demo/config"
	"log"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
	fm "github.com/oapi-codegen/fiber-middleware"
)

type DevServer struct {
	fiber *fiber.App
	cfg   *config.Config
	api   *Api
}

// func NewDevServer(cfg *config.Config, api *server.Api, fiber *fiber.App) *DevServer {
func NewDevServer(fiber *fiber.App, cfg *config.Config, api *Api) *DevServer {
	return &DevServer{
		fiber: fiber,
		cfg:   cfg,
		api:   api,
	}
}

func (s *DevServer) Start() error {
	// log.Infof("Starting server on %s", s.cfg.Server.Address)
	s.Prepare()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	var serverShutdown sync.WaitGroup

	go func() {
		<-c
		log.Println("Gracefully shutting down...")
		serverShutdown.Add(1)
		defer serverShutdown.Done()
		// We add a timeout, after signal Interrupt we will wait for our
		// server end process or in worst case for 1 minute
		_ = s.fiber.ShutdownWithTimeout(60 * time.Second)
	}()

	if err := s.fiber.Listen(":" + s.cfg.HTTP.Port); err != nil {
		log.Fatalln(err)
	}

	// Waiting for start shutting down
	serverShutdown.Wait()
	log.Println("Running cleanup tasks...")

	return nil
}

func (s *DevServer) Prepare() *DevServer {
	router := s.fiber.Group(s.cfg.BaseUrl)

	if s.cfg.UseSwaggerSpec {
		spec, err := generated.GetSwagger()
		if err != nil {
			panic(err)
		}

		router.Use(fm.OapiRequestValidator(spec))

	}

	generated.RegisterHandlers(router, s.api)

	return s

}
