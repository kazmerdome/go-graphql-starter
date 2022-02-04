package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

const (
	MODE_GLOBALENV = "GLOBALENV"
	// errors
	ERR_ALREADY_EXISTS = "secret is already created"
)

type Config interface {
	Get(key string) string
	Set(key string, value string) error
}

type config struct {
	Mode        string
	Environment string
	memCache    *map[string]string
}

func NewConfig(mode string) Config {
	memCache := make(map[string]string)

	// Set default env
	environment := os.Getenv(ENV_ENVIRONMENT)
	if environment == "" {
		environment = ENV_DEVELOPMENT
	}

	if mode == MODE_GLOBALENV && environment != ENV_PRODUCTION {
		godotenv.Load()
	}

	return &config{
		Mode:        mode,
		Environment: environment,
		memCache:    &memCache,
	}
}

func (r *config) Get(key string) string {
	c := *r.memCache
	if c[key] != "" {
		return c[key]
	}
	e := ""

	// If mode MODE_GLOBALENV
	e = os.Getenv(key)
	if r.Mode == MODE_GLOBALENV {
		r.Set(key, e)
	}

	return e
}

func (r *config) Set(key string, value string) error {
	c := *r.memCache

	if c[key] != "" {
		return errors.New(ERR_ALREADY_EXISTS)
	}

	c[key] = value
	r.memCache = &c

	return nil
}
