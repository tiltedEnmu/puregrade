package mongo

import (
	"time"
)

// Option -.
type Option func(options *Mongo)

// ConnTimeout -.
func ConnTimeout(timeout time.Duration) Option {
	return func(m *Mongo) {
		m.connTimeout = timeout
	}
}

// ConnAttempts -.
func ConnAttempts(count int) Option {
	return func(m *Mongo) {
		m.connAttempts = count
	}
}

// Username -.
func Username(value string) Option {
	return func(m *Mongo) {
		m.username = value
	}
}

// Password -.
func Password(value string) Option {
	return func(m *Mongo) {
		m.password = value
	}
}
