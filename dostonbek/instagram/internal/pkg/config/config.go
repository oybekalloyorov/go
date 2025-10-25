package config

import (
    "log"
    "os"
    "strconv"

    "github.com/joho/godotenv"
)

type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func LoadConfig() *DBConfig {
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️  .env fayl topilmadi, default qiymatlar ishlatiladi")
	}
	portStr := getEnv("PORT", "5432")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		port = 5432
	}
	config := &DBConfig{
		Host:     getEnv("HOST", "localhost"),
		Port:     port,
		User:     getEnv("USER", "instagram"),
		Password: getEnv("PASSWORD", "oybek"),
		DBName:   getEnv("DBNAME", "instagram"),
		SSLMode:  getEnv("SSLMODE", "disable"),
	}
	return config
}