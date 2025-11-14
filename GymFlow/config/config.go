package config

import (
	"log"
	"os"
	"strconv"
)

type Config struct {
	DB_DSN        string
	JWTSecret     string
	JWTExpireMins int
	Port          int
}

func Load() *Config {
	portStr := getEnv("PORT", "8080")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatalf("invalid PORT: %v", err)
	}

	expStr := getEnv("JWT_EXPIRE_MINS", "60")
	exp, err := strconv.Atoi(expStr)
	if err != nil {
		log.Fatalf("invalid JWT_EXPIRE_MINS: %v", err)
	}

	return &Config{
    DB_DSN:        getEnv("DB_DSN", "postgres://postgres:123456@localhost:5432/gymflow?sslmode=disable"),
    JWTSecret:     getEnv("JWT_SECRET", "supersecret"),
    JWTExpireMins: exp,
    Port:          port,
}

}

func getEnv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}