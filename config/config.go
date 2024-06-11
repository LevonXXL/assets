package config

import (
	"errors"
	"os"
	"strconv"
	"time"
)

type Config struct {
	AppPort              string
	PgDBDSN              string
	SessionDurationHours *time.Duration
	UseHTTPS             bool
}

func NewConfig() (*Config, error) {
	pgDBDSN := os.Getenv("PG_DB_DSN")
	if pgDBDSN == "" {
		return nil, errors.New("PG_DB_DSN is not configured")
	}

	sessionDurationHours, err := strconv.Atoi(getEnv("SESSION_DURATION_HOURS", "0"))
	if err != nil {
		return nil, err
	}

	cfg := &Config{
		UseHTTPS: getEnv("USE_HTTPS", "false") == "true",
		AppPort:  getEnv("APP_PORT", "8080"),
		PgDBDSN:  pgDBDSN,
	}

	if sessionDurationHours > 0 {
		sessionDuration := time.Duration(sessionDurationHours) * time.Hour
		cfg.SessionDurationHours = &sessionDuration
	}

	return cfg, nil
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
