package handlers

import (
	"net/http"
	"strconv"

	"backend/internal/services/chat"

	"github.com/gin-gonic/gin"
)

// HandleSendMessage godoc
// @Summary Отправка сообщения
// @Description Отправляет сообщение и получает ответ от ассистента
// @Tags chat
// @Accept json
// @Produce json
// @Param request body chat.ChatRequest true "Сообщение пользователя"
// @Success 200 {object} chat.ChatResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /chat [post]
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

// HandleGetChat godoc
// @Summary Получение информации о чате
// @Description Получает информацию о чате по его ID
// @Tags chat
// @Accept json
// @Produce json
// @Param id path int true "ID чата"
// @Success 200 {object} map[string][]models.Message
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /chat/{id} [get]
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

// HandleGetChatMessages godoc
// @Summary Получение сообщений чата
// @Description Получает историю сообщений чата по его ID
// @Tags chat
// @Accept json
// @Produce json
// @Param id path int true "ID чата"
// @Success 200 {object} map[string][]models.Message
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /chat/{id}/messages [get]
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

// HandleListChats godoc
// @Summary Список чатов пользователя
// @Description Получает список чатов пользователя с пагинацией
// @Tags chat
// @Accept json
// @Produce json
// @Param user_id query string false "ID пользователя (если не указан, используется demo_user)"
// @Param limit query int false "Лимит результатов (по умолчанию 10)"
// @Param offset query int false "Смещение (по умолчанию 0)"
// @Success 200 {object} map[string][]models.ChatSummary
// @Failure 500 {object} map[string]string
// @Router /chats [get]
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

// HandleSaveFeedback godoc
// @Summary Сохранение обратной связи
// @Description Сохраняет оценку и комментарий пользователя по сообщению
// @Tags feedback
// @Accept json
// @Produce json
// @Param request body chat.FeedbackRequest true "Данные обратной связи"
// @Success 200 {object} map[string]bool
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /feedback [post]
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
