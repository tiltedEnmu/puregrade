package app

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/segmentio/kafka-go"
	"github.com/tiltedEnmu/puregrade_post/config"
	"github.com/tiltedEnmu/puregrade_post/internal/repository"
	"github.com/tiltedEnmu/puregrade_post/internal/service"
	"github.com/tiltedEnmu/puregrade_post/internal/transport/http"
	"github.com/tiltedEnmu/puregrade_post/pkg/httpserver"
	"github.com/tiltedEnmu/puregrade_post/pkg/mongo"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	// MongoDB
	m, err := mongo.New(
		cfg.Mongo.Address,
		mongo.Username(cfg.Mongo.Username),
		mongo.Password(cfg.Mongo.Password),
	)
	if err != nil {
		log.Fatalf("app - Run - mongo.New: %s", err)
	}

	// Kafka producer
	w := &kafka.Writer{
		Addr:     kafka.TCP(cfg.Kafka.Addresses...),
		Topic:    cfg.Topics.NewPosts,
		Balancer: &kafka.LeastBytes{},
	}

	// Initializing the main application structures
	repos := &repository.Repository{
		Post:     repository.NewMongoPost(m.Client),
		Notifier: repository.NewKafkaNotifications(w),
	}
	services := service.Service{
		Post: service.NewPostService(repos),
	}
	handler := rest.NewHandler(services)
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

	err = m.Close()
	if err != nil {
		log.Printf("app - Run - mongo.Close: %s", err)
	}
}
