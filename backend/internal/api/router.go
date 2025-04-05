package api

import (
	"backend/internal/api/handlers"
	"backend/internal/api/middleware"
	"backend/internal/services/chat"

	_ "backend/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Tender Chat API
// @version 1.0
// @description API для чат-бота с поиском в базе знаний
// @BasePath /api

func SetupRouter(chatService *chat.Service) *gin.Engine {
	router := gin.Default()

	// Middleware
	router.Use(middleware.CORSMiddleware())

	// Swagger
	router.GET("/swagger/*any", ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "DISABLE_SWAGGER"))

	// API группа
	api := router.Group("/api")
	{
		// Чат
		api.POST("/chat", handlers.HandleSendMessage(chatService))
		api.GET("/chat/:id", handlers.HandleGetChat(chatService))

		// История чатов
		api.GET("/chats", handlers.HandleListChats(chatService))

		// Обратная связь
		api.POST("/feedback", handlers.HandleSaveFeedback(chatService))
	}

	return router
}
