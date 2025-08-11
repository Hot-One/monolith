package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	ServiceName string
	Environment string
	HTTPPort    string
	HTTPScheme  string

	PostgresHost     string
	PostgresPort     string
	PostgresUser     string
	PostgresDatabase string
	PostgresPassword string
	PostgresMaxConns int

	RedisHost     string
	RedisPort     string
	RedisUsername string
	RedisPassword string
}

func Load() Config {
	if err := godotenv.Load(".env"); err != nil {
		if err := godotenv.Load("app/.env"); err != nil {
			fmt.Println("No /app/.env file found")
		}
		fmt.Println("No .env file found")
	}

	c := Config{}

	c.ServiceName = cast.ToString(getOrReturnDefault("SERVICE_NAME", ""))
	c.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", ""))
	c.HTTPPort = cast.ToString(getOrReturnDefault("HTTP_PORT", ""))
	c.HTTPScheme = cast.ToString(getOrReturnDefault("HTTP_SCHEME", ""))

	c.PostgresHost = cast.ToString(getOrReturnDefault("POSTGRES_HOST", ""))
	c.PostgresPort = cast.ToString(getOrReturnDefault("POSTGRES_PORT", ""))
	c.PostgresUser = cast.ToString(getOrReturnDefault("POSTGRES_USER", ""))
	c.PostgresDatabase = cast.ToString(getOrReturnDefault("POSTGRES_DATABASE", ""))
	c.PostgresPassword = cast.ToString(getOrReturnDefault("POSTGRES_PASSWORD", ""))
	c.PostgresMaxConns = cast.ToInt(getOrReturnDefault("POSTGRES_MAX_CONNS", 0))

	c.RedisHost = cast.ToString(getOrReturnDefault("REDIS_HOST", ""))
	c.RedisPort = cast.ToString(getOrReturnDefault("REDIS_PORT", ""))
	c.RedisUsername = cast.ToString(getOrReturnDefault("REDIS_USERNAME", ""))
	c.RedisPassword = cast.ToString(getOrReturnDefault("REDIS_PASSWORD", ""))

	return c
}

func getOrReturnDefault(key string, defaultValue any) any {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}
	return defaultValue
}
