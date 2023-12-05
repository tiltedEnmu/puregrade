package app

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	kafkaGo "github.com/segmentio/kafka-go"
	"github.com/tiltedEnmu/puregrade_timeline/config"
	"github.com/tiltedEnmu/puregrade_timeline/internal/repository"
	"github.com/tiltedEnmu/puregrade_timeline/internal/service"
	"github.com/tiltedEnmu/puregrade_timeline/internal/transport/http"
	"github.com/tiltedEnmu/puregrade_timeline/internal/transport/kafka"
	"github.com/tiltedEnmu/puregrade_timeline/pkg/httpserver"
	"github.com/tiltedEnmu/puregrade_timeline/pkg/redis"
)

// Run creates objects via constructors.
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
	repos := repository.NewRedisTimeline(r.Client)
	services := service.NewService(repos)
	handler := rest.NewHandler(services)
	handler.InitRoutes()

	consumer := kafka.NewConsumer(
		kafkaGo.NewReader(
			kafkaGo.ReaderConfig{
				Brokers:  cfg.Addresses,
				Topic:    cfg.Kafka.Topics.NewPosts,
				GroupID:  "timeline-service-group",
				MaxBytes: 10e6, // 10MB
			},
		),
		services,
	)
	consumer.Run()

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

	err = consumer.Close()
	if err != nil {
		log.Printf("app - Run - kafka.Close: %s", err)
	}

	err = r.Close()
	if err != nil {
		log.Printf("app - Run - mongo.Close: %s", err)
	}
}
