package api

import (
	"backend/internal/api/handlers"
	"backend/internal/api/middleware"
	"backend/internal/services/chat"

	"github.com/gin-gonic/gin"
)

// SetupRouter настраивает маршруты API
func SetupRouter(chatService *chat.Service) *gin.Engine {
	router := gin.Default()

	// Middleware
	router.Use(middleware.CORSMiddleware())

	// API группа
	api := router.Group("/api")
	{
		// Чат
		api.POST("/chat", handlers.HandleSendMessage(chatService))
		api.GET("/chat/:id", handlers.HandleGetChat(chatService))
		api.GET("/chat/:id/messages", handlers.HandleGetChatMessages(chatService))

		// История чатов
		api.GET("/chats", handlers.HandleListChats(chatService))

		// Обратная связь
		api.POST("/feedback", handlers.HandleSaveFeedback(chatService))
	}

	return router
}
