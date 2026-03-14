package config

import (
	"fmt"
	"os"
)

type Config struct {
	AppPort     string
	MySQLDSN    string
	KafkaBroker string
}

func Load() Config {
	host := getEnv("MYSQL_HOST", "localhost")
	port := getEnv("MYSQL_PORT", "3306")
	user := getEnv("MYSQL_USER", "root")
	password := getEnv("MYSQL_PASSWORD", "root")
	db := getEnv("MYSQL_DB", "tasks_db")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		user,
		password,
		host,
		port,
		db,
	)

	return Config{
		AppPort:     getEnv("APP_PORT", "8080"),
		MySQLDSN:    dsn,
		KafkaBroker: getEnv("KAFKA_BROKER", "localhost:9092"),
	}
}

func getEnv(key string, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
