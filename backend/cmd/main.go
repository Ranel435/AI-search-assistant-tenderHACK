package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"your-project/handlers"
	"your-project/services"
)

func main() {
	router := gin.Default()

	// Инициализация сервисов
	llmService := services.NewLLMService()
	dbService := services.NewDBService()
	vectorDBService := services.NewVectorDBService()

	// Регистрация маршрутов
	router.POST("/api/chat", handlers.ChatHandler(llmService, dbService, vectorDBService))
	router.GET("/api/history", handlers.GetHistoryHandler(dbService))
	router.POST("/api/feedback", handlers.FeedbackHandler(dbService))

	// Запуск сервера
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to run server:", err)
	}
}
