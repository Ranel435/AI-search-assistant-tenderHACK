// @title        Tender Chat API
// @version      1.0
// @description  API для чат-бота с контекстным поиском
// @BasePath     /api
package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"backend/internal/api"
	"backend/internal/config"
	"backend/internal/db/postgres"
	"backend/internal/services/chat"
	"backend/internal/services/llm"
	"backend/internal/services/search"
)

func main() {
	// Загрузка конфигурации
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Подключение к БД
	db, err := postgres.NewDB(cfg.DB)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Инициализация сервисов
	llmService := llm.NewService(cfg.LLM)
	searchService := search.NewService(db)
	chatService := chat.NewService(db, llmService, searchService)

	// Инициализация API
	router := api.SetupRouter(chatService)

	// Настройка HTTP сервера
	server := &http.Server{
		Addr:    cfg.Server.Address,
		Handler: router,
	}

	// Канал для грацеозной остановки
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Запуск сервера в отдельной горутине
	go func() {
		log.Printf("Server starting on %s", cfg.Server.Address)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error starting server: %v", err)
		}
	}()

	// Ожидание сигнала остановки
	<-quit
	log.Println("Shutting down server...")

	// Грацеозная остановка
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited properly")
}
