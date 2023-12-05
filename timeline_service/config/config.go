package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	// Config -.
	Config struct {
		App   `yaml:"app"`
		HTTP  `yaml:"http"`
		GRPC  `yaml:"grpc"`
		Kafka `yaml:"kafka"`
		Redis `yaml:"redis"`
	}

	// App -.
	App struct {
		Name    string `env-required:"true" yaml:"name"    env:"APP_NAME"`
		Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
	}

	// HTTP -.
	HTTP struct {
		Port string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
	}

	// GRPC -.
	GRPC struct {
		Port string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
	}

	// Redis -.
	Redis struct {
		Address  string `env-required:"true" yaml:"address" env:"REDIS_ADDRESS"`
		Password string `env-required:"true" yaml:"password" env:"REDIS_PASSWORD"`
	}

	Kafka struct {
		Addresses []string `env-required:"true" yaml:"addresses" env:"ADDRESSES"`
		Topics    Topics   `yaml:"topics"`
	}

	Topics struct {
		NewPosts string `env-required:"true" yaml:"newPosts" env:"NEW_POSTS"`
	}
)

// NewConfig returns app config.
func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig("./config/config.yml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
