package config

import (
    "log"
    "os"
)

type Config struct {
    DatabaseURL string
    LogLevel    string
}

func LoadConfig() *Config {
    dbURL := os.Getenv("DATABASE_URL")
    if dbURL == "" {
        log.Fatal("DATABASE_URL environment variable is required")
    }

    logLevel := os.Getenv("LOG_LEVEL")
    if logLevel == "" {
        logLevel = "info" // default log level
    }

    return &Config{
        DatabaseURL: dbURL,
        LogLevel:    logLevel,
    }
}