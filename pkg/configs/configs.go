package configs

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var envs map[string]string
var connectionString string

func init() {
	loadEnvs()
	connectionString = fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		GetEnv("DB_USER"), GetEnv("DB_PASSWORD"), GetEnv("DB_HOST"), GetEnv("DB_PORT"), GetEnv("DB_NAME"))
}

func GetEnv(name string) string {
	return envs[name]
}

func GetConnectionString() string {
	return connectionString
}

func IsProduction() bool {
	env := GetEnv("APP_ENV")
	return env == "prd" || env == "production" || env == "prod"
}

func IsDevelopment() bool {
	env := GetEnv("APP_ENV")
	return env == "dev" || env == "development" || env == "develop"
}

func IsLocal() bool {
	env := GetEnv("APP_ENV")
	return env == "local"
}

func loadEnvs() {
	err := godotenv.Load()
	if err != nil && IsLocal() {
		log.Fatal("Error loading .env file")
	}

	envs = map[string]string{
		"APP_ENV":  os.Getenv("APP_ENV"),
		"APP_PORT": os.Getenv("PORT"),

		"MEILISEARCH_HOST":    os.Getenv("MEILISEARCH_HOST"),
		"MEILISEARCH_API_KEY": os.Getenv("MEILISEARCH_API_KEY"),

		"JWT_SECRET": os.Getenv("JWT_SECRET"),

		"DB_HOST":     os.Getenv("DB_HOST"),
		"DB_PORT":     os.Getenv("DB_PORT"),
		"DB_NAME":     os.Getenv("DB_NAME"),
		"DB_USER":     os.Getenv("DB_USER"),
		"DB_PASSWORD": os.Getenv("DB_PASSWORD"),
		"DB_SSL_MODE": os.Getenv("DB_SSL_MODE"),

		"SMTP_HOST":     os.Getenv("SMTP_HOST"),
		"SMTP_PORT":     os.Getenv("SMTP_PORT"),
		"SMTP_FROM":     os.Getenv("SMTP_FROM"),
		"SMTP_USERNAME": os.Getenv("SMTP_USERNAME"),
		"SMTP_PASSWORD": os.Getenv("SMTP_PASSWORD"),
	}
}
