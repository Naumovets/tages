package config

import (
	"log"
	"os"
	"strconv"
)

func getEnv(key string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	log.Fatal()
	return ""
}

type StorageConfig struct {
	Location  string
	BatchSize int
}

type DBConfig struct {
	DB_NAME  string
	USER     string
	PASSWORD string
	DB_HOST  string
	DB_PORT  string
}

type Config struct {
	Storage StorageConfig
	DB      DBConfig
}

// New returns a new Config struct
func New() *Config {
	return &Config{
		Storage: StorageConfig{
			Location: getEnv("STORAGE_LOCATION"),
			BatchSize: func() int {
				value, ok := os.LookupEnv("STORAGE_BATCH_SIZE")
				if !ok {
					log.Fatal("STORAGE_BATCH_SIZE is not set")
				}
				i, err := strconv.Atoi(value)
				if err != nil {
					log.Fatal(err)
				}
				return i
			}(),
		},
		DB: DBConfig{
			DB_NAME:  getEnv("DB_NAME"),
			USER:     getEnv("DB_USER"),
			PASSWORD: getEnv("DB_PASSWORD"),
			DB_HOST:  getEnv("DB_HOST"),
			DB_PORT:  getEnv("DB_PORT"),
		},
	}
}
