package chat

import (
	"context"
	"fmt"

	"backend/internal/db/models"
	"backend/internal/db/postgres"
	"backend/internal/services/llm"
	"backend/internal/services/search"
)

// Service предоставляет функциональность для работы с чатами
type Service struct {
	chatRepo      postgres.ChatRepository
	llmService    *llm.Service
	searchService *search.Service
}

// NewService создаёт новый сервис для чатов
func NewService(db *postgres.DB, llmService *llm.Service, searchService *search.Service) *Service {
	return &Service{
		chatRepo:      postgres.NewChatRepo(db),
		llmService:    llmService,
		searchService: searchService,
	}
}

// ChatRequest представляет запрос на отправку сообщения
type ChatRequest struct {
	UserID  string `json:"user_id"`
	ChatID  int64  `json:"chat_id,omitempty"`
	Message string `json:"message"`
}

// ChatResponse представляет ответ на запрос чата
type ChatResponse struct {
	ChatID  int64    `json:"chat_id"`
	Message string   `json:"message"`
	Sources []string `json:"sources,omitempty"`
}

// SendMessage обрабатывает сообщение пользователя и генерирует ответ
func (s *Service) SendMessage(ctx context.Context, req ChatRequest) (*ChatResponse, error) {
	// Если чат не указан, создаём новый
	var chat *models.Chat
	var err error

	if req.ChatID == 0 {
		// Создаём заголовок на основе первого сообщения
		title := req.Message
		if len(title) > 50 {
			title = title[:50] + "..."
		}

		chat, err = s.chatRepo.CreateChat(ctx, req.UserID, title)
		if err != nil {
			return nil, fmt.Errorf("failed to create chat: %w", err)
		}
	} else {
		// Проверяем существующий чат
		chat, err = s.chatRepo.GetChat(ctx, req.ChatID)
		if err != nil {
			return nil, fmt.Errorf("failed to get chat: %w", err)
		}
		if chat == nil {
			return nil, fmt.Errorf("chat not found")
		}
	}

	// Сохраняем сообщение пользователя
	_, err = s.chatRepo.CreateMessage(ctx, chat.ID, "user", req.Message)
	if err != nil {
		return nil, fmt.Errorf("failed to save user message: %w", err)
	}

	// Ищем релевантную информацию по запросу
	docs, err := s.searchService.Search(ctx, req.Message)
	if err != nil {
		return nil, fmt.Errorf("failed to search for documents: %w", err)
	}

	// Готовим контекст для LLM
	context := s.searchService.PrepareContext(docs)

	// Получаем источники
	sources := s.searchService.GetSourceReferences(docs)

	// Генерируем ответ с помощью LLM
	answer, err := s.llmService.GenerateAnswer(ctx, req.Message, context)
	if err != nil {
		// Если произошла ошибка, отправляем заготовленный ответ
		answer = "Извините, я не смог обработать ваш запрос. Пожалуйста, попробуйте переформулировать вопрос или обратитесь в службу поддержки."
	}

	// Сохраняем ответ ассистента
	_, err = s.chatRepo.CreateMessage(ctx, chat.ID, "assistant", answer)
	if err != nil {
		return nil, fmt.Errorf("failed to save assistant message: %w", err)
	}

	return &ChatResponse{
		ChatID:  chat.ID,
		Message: answer,
		Sources: sources,
	}, nil
}

// GetChatHistory возвращает историю сообщений чата
func (s *Service) GetChatHistory(ctx context.Context, chatID int64) ([]models.Message, error) {
	return s.chatRepo.GetMessages(ctx, chatID)
}

// ListUserChats возвращает список чатов пользователя
func (s *Service) ListUserChats(ctx context.Context, userID string, limit, offset int) ([]models.ChatSummary, error) {
	return s.chatRepo.GetChatSummaries(ctx, userID, limit, offset)
}

// FeedbackRequest представляет запрос на отправку обратной связи
type FeedbackRequest struct {
	MessageID int64  `json:"message_id"`
	Rating    int    `json:"rating"`
	Comment   string `json:"comment,omitempty"`
}

// SaveFeedback сохраняет обратную связь пользователя
func (s *Service) SaveFeedback(ctx context.Context, req FeedbackRequest) error {
	_, err := s.chatRepo.SaveFeedback(ctx, req.MessageID, req.Rating, req.Comment)
	if err != nil {
		return fmt.Errorf("failed to save feedback: %w", err)
	}
	return nil
}
