package main

import (
	"log"
	"simple-microservice/api"
)

func main() {
	server, err := api.NewServer("5200")
	if err != nil {
		log.Fatalf("Error creating server: %v", err)
	}

	if err := server.Start(); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
