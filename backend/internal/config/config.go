package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config - основная конфигурация приложения
type Config struct {
	Server ServerConfig
	DB     DBConfig
	LLM    LLMConfig
}

// ServerConfig - конфигурация HTTP сервера
type ServerConfig struct {
	Address string
}

// DBConfig - конфигурация базы данных
type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// LLMConfig - конфигурация LLM сервиса
type LLMConfig struct {
	URL       string
	Timeout   int // в секундах
	MaxTokens int
}

// Load загружает конфигурацию из переменных окружения
func Load() (*Config, error) {
	// Загружаем переменные из .env файла
	if err := godotenv.Load(); err != nil {
		// Продолжаем выполнение, даже если файл .env не найден
		fmt.Println("Warning: .env file not found or cannot be loaded")
	}

	port, _ := strconv.Atoi(getEnv("DB_PORT", "5432"))
	llmTimeout, _ := strconv.Atoi(getEnv("LLM_TIMEOUT", "30"))
	maxTokens, _ := strconv.Atoi(getEnv("LLM_MAX_TOKENS", "1024"))

	return &Config{
		Server: ServerConfig{
			Address: getEnv("SERVER_ADDRESS", ":8080"),
		},
		DB: DBConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     port,
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", "postgres"),
			DBName:   getEnv("DB_NAME", "tenderhack"),
			SSLMode:  getEnv("DB_SSLMODE", "disable"),
		},
		LLM: LLMConfig{
			URL:       getEnv("LLM_SERVICE_URL", "http://localhost:8000"),
			Timeout:   llmTimeout,
			MaxTokens: maxTokens,
		},
	}, nil
}

// getEnv получает переменную окружения или возвращает значение по умолчанию
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
