package models

import "time"

// Chat представляет собой чат между пользователем и ассистентом
type Chat struct {
	ID        int64     `db:"id" json:"id"`
	UserID    string    `db:"user_id" json:"user_id"`
	Title     string    `db:"title" json:"title"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

// Message представляет собой сообщение в чате
type Message struct {
	ID        int64     `db:"id" json:"id"`
	ChatID    int64     `db:"chat_id" json:"chat_id"`
	Role      string    `db:"role" json:"role"` // user или assistant
	Content   string    `db:"content" json:"content"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

// Feedback представляет собой обратную связь от пользователя
type Feedback struct {
	ID        int64     `db:"id" json:"id"`
	MessageID int64     `db:"message_id" json:"message_id"`
	Rating    int       `db:"rating" json:"rating"` // от 1 до 5
	Comment   string    `db:"comment" json:"comment,omitempty"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

// Document представляет собой документ в базе знаний
type Document struct {
	ID        int64     `db:"id" json:"id"`
	Title     string    `db:"title" json:"title"`
	Content   string    `db:"content" json:"content"`
	Source    string    `db:"source" json:"source"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

// ChatSummary представляет сводку чата с метаданными
type ChatSummary struct {
	ID            int64     `db:"id" json:"id"`
	UserID        string    `db:"user_id" json:"user_id"`
	Title         string    `db:"title" json:"title"`
	Summary       string    `db:"summary" json:"summary"`
	Category      string    `db:"category" json:"category"`
	CreatedAt     time.Time `db:"created_at" json:"created_at"`
	MessageCount  int       `db:"message_count" json:"message_count"`
	AverageRating float64   `db:"average_rating" json:"average_rating,omitempty"`
}
