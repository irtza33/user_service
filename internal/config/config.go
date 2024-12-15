package config

import (
    "os"
)

type Config struct {
    Database DatabaseConfig
    Server   ServerConfig
    LogLevel string
}

type DatabaseConfig struct {
    Host     string
    Port     string
    User     string
    Password string
    DBName   string
}

func LoadConfig() *Config {
    return &Config{
        Database: DatabaseConfig{
            Host:     getEnv("DB_HOST", "localhost"),
            Port:     getEnv("DB_PORT", "5432"),
            User:     getEnv("DB_USER", "postgres"),
            Password: getEnv("DB_PASSWORD", "postgres"),
            DBName:   getEnv("DB_NAME", "userdb"),
        },
        Server: ServerConfig{
            Address: getEnv("SERVER_ADDRESS", ":50051"),
        },
        LogLevel: getEnv("LOG_LEVEL", "info"),
    }
}

func getEnv(key, defaultValue string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return defaultValue
}