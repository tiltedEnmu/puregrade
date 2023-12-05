package main

import (
	"log"

	"github.com/puregrade/puregrade-auth/config"
	"github.com/puregrade/puregrade-auth/internal/app"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	app.Run(cfg)
}
