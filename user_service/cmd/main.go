package main

import (
	"log"

	"github.com/spf13/viper"

	"github.com/tiltedEnmu/puregrade-user/internal/repository"
	"github.com/tiltedEnmu/puregrade-user/internal/service"
)

func main() {
	viper.AddConfigPath("configs") // config directory
	viper.SetConfigName("config")  // config filename

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Config init error: ", err.Error())
	}

	db, err := repository.NewPostgresDB(
		repository.PGConfig{
			Host:     viper.GetString("pg.host"),
			Port:     viper.GetString("pg.port"),
			Username: viper.GetString("pg.username"),
			Password: viper.GetString("pg.password"),
			DBName:   viper.GetString("pg.dbname"),
			SSLMode:  viper.GetString("pg.sslmode"),
		},
	)

	if err != nil {
		log.Print("DB init error: %s", err.Error())
	}

	// init microservice structure
	repos := repository.NewRepository(db)
	service := service.NewService(repos)
}
