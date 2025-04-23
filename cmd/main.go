package main

import (
	"log"

	"github.com/lipaysamart/go-webhook-exercise/internal/bootstrap"
)

func main() {
	server := bootstrap.NewBootStrap()
	if err := server.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
