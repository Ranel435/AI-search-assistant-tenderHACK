package handlers

import (
	"net/http"
	"strconv"

	"backend/internal/services/chat"

	"github.com/gin-gonic/gin"
)

// HandleSendMessage обрабатывает отправку сообщения
func HandleSendMessage(chatService *chat.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req chat.ChatRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		if req.Message == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Message is required"})
			return
		}

		// Временная заглушка для демо
		if req.UserID == "" {
			req.UserID = "demo_user"
		}

		response, err := chatService.SendMessage(c.Request.Context(), req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, response)
	}
}

// HandleGetChat получает информацию о чате
func HandleGetChat(chatService *chat.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		chatID, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid chat ID"})
			return
		}

		messages, err := chatService.GetChatHistory(c.Request.Context(), chatID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"messages": messages})
	}
}

// HandleGetChatMessages получает сообщения чата
func HandleGetChatMessages(chatService *chat.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		chatID, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid chat ID"})
			return
		}

		messages, err := chatService.GetChatHistory(c.Request.Context(), chatID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"messages": messages})
	}
}

// HandleListChats получает список чатов пользователя
func HandleListChats(chatService *chat.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Query("user_id")
		if userID == "" {
			userID = "demo_user" // Временно для демо
		}

		limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
		offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

		chats, err := chatService.ListUserChats(c.Request.Context(), userID, limit, offset)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"chats": chats})
	}
}

// HandleSaveFeedback сохраняет обратную связь пользователя
func HandleSaveFeedback(chatService *chat.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req chat.FeedbackRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		if req.MessageID == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Message ID is required"})
			return
		}

		if req.Rating < 1 || req.Rating > 5 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Rating must be between 1 and 5"})
			return
		}

		if err := chatService.SaveFeedback(c.Request.Context(), req); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"success": true})
	}
}
