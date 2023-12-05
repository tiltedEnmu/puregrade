package main

import (
	"log"

	"github.com/tiltedEnmu/puregrade_timeline/config"
	"github.com/tiltedEnmu/puregrade_timeline/internal/app"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	app.Run(cfg)
}
