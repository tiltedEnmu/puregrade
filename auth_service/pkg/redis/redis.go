package redis

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

const (
	_defaultDatabase     = 0
	_defaultPassword     = ""
	_defaultConnAttempts = 3
	_defaultPingTimeout  = time.Second
	_defaultConnTimeout  = 3 * time.Second
	_defaultReadTimeout  = time.Second
	_defaultWriteTimeout = time.Second
)

type Redis struct {
	pingTimeout  time.Duration
	connTimeout  time.Duration
	connAttempts int

	opts   *redis.Options
	Client *redis.Client
}

func New(addr string, opts ...Option) (*Redis, error) {
	r := &Redis{
		pingTimeout:  _defaultPingTimeout,
		connAttempts: _defaultConnAttempts,
		connTimeout:  _defaultConnTimeout,
		opts: &redis.Options{
			Addr:         addr,
			Password:     _defaultPassword,
			DB:           _defaultDatabase,
			ReadTimeout:  _defaultReadTimeout,
			WriteTimeout: _defaultWriteTimeout,
		},
	}

	for _, opt := range opts {
		opt(r)
	}

	r.Client = redis.NewClient(r.opts)

	var err error
	for r.connAttempts > 0 {
		ctx, cancel := context.WithTimeout(context.Background(), r.pingTimeout)
		defer cancel()

		if err = r.Client.Ping(ctx).Err(); err == nil {
			break
		}

		log.Printf("redis is trying to connect, attempts left: %d", r.connAttempts)

		time.Sleep(r.connTimeout)

		r.connAttempts--
	}

	if err != nil {
		return nil, fmt.Errorf("redis connection error: %w", err)
	}

	return r, nil
}

func (r *Redis) Close() error {
	if r.Client != nil {
		return r.Client.Close()
	}

	return fmt.Errorf("redis client doesn't exists")
}
