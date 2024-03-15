package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	Cfg *Config
)

type Config struct {
	PORT                   string
	DBHost                 string
	DBUserName             string
	DBPassword             string
	DBName                 string
	DBPort                 string
	DBssl                  string
	DBTimezone             string
	REDIS_URL              string
	REDIS_PORT             string
	REDIS_PASSWORD         string
	REDIS_REPLICATION_MODE string
	SECRET_KEY             string
	VERSION                string
	API_KEY                string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return nil, err
	}

	config := &Config{
		PORT:                   os.Getenv("PORT"),
		DBHost:                 os.Getenv("DB_HOST"),
		DBUserName:             os.Getenv("DB_USERNAME"),
		DBPassword:             os.Getenv("DB_PASSWORD"),
		DBPort:                 os.Getenv("DB_PORT"),
		DBName:                 os.Getenv("DB_NAME"),
		DBssl:                  os.Getenv("DB_SSL"),
		DBTimezone:             os.Getenv("DB_TIMEZONE"),
		REDIS_URL:              os.Getenv("REDIS_URL"),
		REDIS_PORT:             os.Getenv("REDIS_PORT"),
		REDIS_PASSWORD:         os.Getenv("REDIS_PASSWORD"),
		REDIS_REPLICATION_MODE: os.Getenv("REDIS_REPLICATION_MODE"),
		SECRET_KEY:             os.Getenv("SECRET_KEY"),
		VERSION:                os.Getenv("VERSION"),
		API_KEY:                os.Getenv("API_KEY"),
	}

	Cfg = config

	return config, nil
}
