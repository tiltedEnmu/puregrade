package app

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/puregrade/puregrade-auth/config"
	"github.com/puregrade/puregrade-auth/internal/repository"
	"github.com/puregrade/puregrade-auth/internal/service"
	"github.com/puregrade/puregrade-auth/internal/transport/rest"
	"github.com/puregrade/puregrade-auth/pkg/httpserver"
	"github.com/puregrade/puregrade-auth/pkg/redis"
)

func Run(cfg *config.Config) {
	// Redis
	r, err := redis.New(
		cfg.Redis.Address,
		redis.ConnTimeout(time.Second),
	)
	if err != nil {
		return
	}

	// Initializing the main application structures
	repos := repository.NewRepository(r.Client)
	services := service.NewService(repos)
	handler := rest.NewHandler(services, cfg.ExtServicesAddrs.UserServiceAddr)
	handler.InitRoutes()

	// HTTP Server
	httpServer := httpserver.New(
		handler.Routes,
		httpserver.Port(cfg.HTTP.Port),
	)

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Printf("app - Run - signal: %s", s.String())
	case err := <-httpServer.Notify():
		log.Printf("app - Run - httpServer.Notify: %s", err)
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		log.Printf("app - Run - httpServer.Shutdown: %s", err)
	}

	err = r.Close()
	if err != nil {
		log.Printf("app - Run - mongo.Close: %s", err)
	}
}
