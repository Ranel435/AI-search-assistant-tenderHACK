package postgres

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"backend/internal/db/models"
)

// ChatRepository интерфейс для работы с чатами
type ChatRepository interface {
	CreateChat(ctx context.Context, userID, title string) (*models.Chat, error)
	GetChat(ctx context.Context, id int64) (*models.Chat, error)
	ListChats(ctx context.Context, userID string, limit, offset int) ([]models.Chat, error)

	CreateMessage(ctx context.Context, chatID int64, role, content string) (*models.Message, error)
	GetMessages(ctx context.Context, chatID int64) ([]models.Message, error)

	SaveFeedback(ctx context.Context, messageID int64, rating int, comment string) (*models.Feedback, error)
	GetChatSummaries(ctx context.Context, userID string, limit, offset int) ([]models.ChatSummary, error)
}

// ChatRepo имплементация ChatRepository
type ChatRepo struct {
	db *DB
}

// NewChatRepo создаёт новый репозиторий для работы с чатами
func NewChatRepo(db *DB) ChatRepository {
	return &ChatRepo{db: db}
}

// CreateChat создаёт новый чат
func (r *ChatRepo) CreateChat(ctx context.Context, userID, title string) (*models.Chat, error) {
	chat := &models.Chat{
		UserID:    userID,
		Title:     title,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	query := `
		INSERT INTO chats (user_id, title, created_at, updated_at)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`

	err := r.db.DB.QueryRowxContext(
		ctx, query, chat.UserID, chat.Title, chat.CreatedAt, chat.UpdatedAt,
	).Scan(&chat.ID)

	if err != nil {
		return nil, err
	}

	return chat, nil
}

// GetChat получает чат по ID
func (r *ChatRepo) GetChat(ctx context.Context, id int64) (*models.Chat, error) {
	query := `SELECT id, user_id, title, created_at, updated_at FROM chats WHERE id = $1`

	var chat models.Chat
	err := r.db.DB.GetContext(ctx, &chat, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // Чат не найден
		}
		return nil, err
	}

	return &chat, nil
}

// ListChats получает список чатов пользователя
func (r *ChatRepo) ListChats(ctx context.Context, userID string, limit, offset int) ([]models.Chat, error) {
	query := `
		SELECT id, user_id, title, created_at, updated_at
		FROM chats
		WHERE user_id = $1
		ORDER BY updated_at DESC
		LIMIT $2 OFFSET $3
	`

	var chats []models.Chat
	err := r.db.DB.SelectContext(ctx, &chats, query, userID, limit, offset)
	if err != nil {
		return nil, err
	}

	return chats, nil
}

// CreateMessage добавляет сообщение в чат
func (r *ChatRepo) CreateMessage(ctx context.Context, chatID int64, role, content string) (*models.Message, error) {
	message := &models.Message{
		ChatID:    chatID,
		Role:      role,
		Content:   content,
		CreatedAt: time.Now(),
	}

	query := `
		INSERT INTO messages (chat_id, role, content, created_at)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`

	err := r.db.DB.QueryRowxContext(
		ctx, query, message.ChatID, message.Role, message.Content, message.CreatedAt,
	).Scan(&message.ID)

	if err != nil {
		return nil, err
	}

	// Обновляем время последнего сообщения в чате
	updateQuery := `UPDATE chats SET updated_at = $1 WHERE id = $2`
	_, err = r.db.DB.ExecContext(ctx, updateQuery, message.CreatedAt, chatID)
	if err != nil {
		return nil, err
	}

	return message, nil
}

// GetMessages получает все сообщения чата
func (r *ChatRepo) GetMessages(ctx context.Context, chatID int64) ([]models.Message, error) {
	query := `
		SELECT id, chat_id, role, content, created_at
		FROM messages
		WHERE chat_id = $1
		ORDER BY created_at ASC
	`

	var messages []models.Message
	err := r.db.DB.SelectContext(ctx, &messages, query, chatID)
	if err != nil {
		return nil, err
	}

	return messages, nil
}

// SaveFeedback сохраняет обратную связь на сообщение
func (r *ChatRepo) SaveFeedback(ctx context.Context, messageID int64, rating int, comment string) (*models.Feedback, error) {
	feedback := &models.Feedback{
		MessageID: messageID,
		Rating:    rating,
		Comment:   comment,
		CreatedAt: time.Now(),
	}

	query := `
		INSERT INTO feedback (message_id, rating, comment, created_at)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`

	err := r.db.DB.QueryRowxContext(
		ctx, query, feedback.MessageID, feedback.Rating, feedback.Comment, feedback.CreatedAt,
	).Scan(&feedback.ID)

	if err != nil {
		return nil, err
	}

	return feedback, nil
}

// GetChatSummaries получает сводку о чатах с метаданными
func (r *ChatRepo) GetChatSummaries(ctx context.Context, userID string, limit, offset int) ([]models.ChatSummary, error) {
	query := `
		SELECT c.id, c.user_id, c.title, c.created_at,
			   cs.summary, cs.category,
			   COUNT(m.id) AS message_count,
			   AVG(f.rating) AS average_rating
		FROM chats c
		LEFT JOIN chat_summaries cs ON c.id = cs.chat_id
		LEFT JOIN messages m ON c.id = m.chat_id
		LEFT JOIN feedback f ON m.id = f.message_id
		WHERE c.user_id = $1
		GROUP BY c.id, cs.summary, cs.category
		ORDER BY c.updated_at DESC
		LIMIT $2 OFFSET $3
	`

	var summaries []models.ChatSummary
	err := r.db.DB.SelectContext(ctx, &summaries, query, userID, limit, offset)
	if err != nil {
		return nil, err
	}

	return summaries, nil
}
