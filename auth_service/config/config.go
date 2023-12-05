package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	// Config -.
	Config struct {
		App              `yaml:"app"`
		HTTP             `yaml:"http"`
		GRPC             `yaml:"grpc"`
		Redis            `yaml:"redis"`
		ExtServicesAddrs `yaml:"externalServiceAddresses"`
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

	ExtServicesAddrs struct {
		UserServiceAddr string `yaml:"userService" env:"USER_SERVICE_ADDR"`
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
