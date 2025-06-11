package main

import (
	"golang-docker-demo/app/server"
	"log"
)

func Start() {
	log.Println("Starting API server...")
	sever := server.Ready()

	if err := sever.Start(); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
