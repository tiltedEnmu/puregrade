package mongo

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	_defaultConnAttempts = 3
	_defaultConnTimeout  = 1 * time.Second
)

type Mongo struct {
	connTimeout  time.Duration
	connAttempts int

	username string
	password string

	Client *mongo.Client
}

func New(addr string, opts ...Option) (*Mongo, error) {
	m := &Mongo{
		connTimeout:  _defaultConnTimeout,
		connAttempts: _defaultConnAttempts,
		Client:       nil,
	}

	for _, opt := range opts {
		opt(m)
	}

	uri := fmt.Sprintf("mongodb://127.0.0.1:48012/")
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	mongoOpts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	var err error
	for m.connAttempts > 0 {
		ctx, cancel := context.WithTimeout(context.Background(), m.connTimeout)
		defer cancel()
		log.Printf("mongo is trying to connect, attempts left: %d", m.connAttempts)

		var client *mongo.Client
		client, err = mongo.Connect(context.TODO(), mongoOpts)
		err = client.Ping(ctx, readpref.Primary())
		if err == nil {
			m.Client = client
			break
		}

		time.Sleep(m.connTimeout)

		m.connAttempts--
	}

	if err != nil {
		return nil, fmt.Errorf("mongo connection error: %s", err.Error())
	}

	return m, nil
}

func (m *Mongo) Close() error {
	if m.Client != nil {
		return m.Client.Disconnect(context.TODO())
	}

	return fmt.Errorf("mongo client doesn't exists")
}
