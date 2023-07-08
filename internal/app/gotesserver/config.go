package gotesserver

import (
    "fmt"
    "github.com/caarlos0/env/v6"
    "github.com/joho/godotenv"
    "log"
    "os"
)

// Common Config
type Config struct {
    BindAddr          string `env:"APP_BIND_ADDR"`
    LogLevel          string `env:"APP_LOG_LEVEL"`
    DatabaseURL       string `env:"DATABASE_URL"`
    DatabaseTestURL   string `env:"DATABASE_TEST_URL"`
    TokenExpiredHours int    `env:"APP_TOKEN_EXPIRED_HOURS"`
    TokenSecret       string `env:"APP_TOKEN_SECRET"`
}

// Application Config
func NewConfig() *Config {
    fmt.Println(os.Args)
    err := godotenv.Load("./configs/.env")

    if err != nil {
        log.Fatal("Error loading .env file")
    }

    err = godotenv.Overload("./configs/.env.local")
    if err != nil {
        log.Fatal("Error overloading .env file")
    }

    config := &Config{}
    err = env.Parse(config)
    if err != nil {
        log.Fatal("Error to parse envs files")
    }

    return config
}

// Config for tests
func NewTestConfig() *Config {
    config := &Config{
        BindAddr:          ":8025",
        TokenSecret:       "secret",
        TokenExpiredHours: 72,
    }

    return config
}
