package redis

import (
	"time"
)

// Option -.
type Option func(options *Redis)

// DB -.
func DB(db int) Option {
	return func(r *Redis) {
		r.opts.DB = db
	}
}

// Password -.
func Password(password string) Option {
	return func(r *Redis) {
		r.opts.Password = password
	}
}

// WriteTimeout -.
func WriteTimeout(timeout time.Duration) Option {
	return func(r *Redis) {
		r.opts.WriteTimeout = timeout
	}
}

// ReadTimeout -.
func ReadTimeout(timeout time.Duration) Option {
	return func(r *Redis) {
		r.opts.ReadTimeout = timeout
	}
}

// PingTimeout -.
func PingTimeout(timeout time.Duration) Option {
	return func(r *Redis) {
		r.pingTimeout = timeout
	}
}

// ConnTimeout -.
func ConnTimeout(timeout time.Duration) Option {
	return func(r *Redis) {
		r.connTimeout = timeout
	}
}

// ConnAttempts -.
func ConnAttempts(count int) Option {
	return func(r *Redis) {
		r.connAttempts = count
	}
}
