package repository

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoConfig struct {
	Host     string
	Port     string
	Username string
	Password string
}

func NewMongoDB(cfg MongoConfig) (*mongo.Client, error) {
	uri := fmt.Sprintf("mongodb://%v:%v@%v:%v/", cfg.Username, cfg.Password, cfg.Host, cfg.Port)

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	return mongo.Connect(context.TODO(), opts)
}
