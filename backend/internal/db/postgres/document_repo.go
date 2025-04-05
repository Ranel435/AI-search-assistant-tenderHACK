package postgres

import (
	"context"

	"backend/internal/db/models"
)

// DocumentRepository интерфейс для работы с документами
type DocumentRepository interface {
	CreateDocument(ctx context.Context, title, content, source string) (*models.Document, error)
	SearchDocuments(ctx context.Context, query string, limit int) ([]models.Document, error)
	GetDocument(ctx context.Context, id int64) (*models.Document, error)
}

// DocumentRepo имплементация DocumentRepository
type DocumentRepo struct {
	db *DB
}

// NewDocumentRepo создаёт новый репозиторий для работы с документами
func NewDocumentRepo(db *DB) DocumentRepository {
	return &DocumentRepo{db: db}
}

// CreateDocument создаёт новый документ
func (r *DocumentRepo) CreateDocument(ctx context.Context, title, content, source string) (*models.Document, error) {
	document := &models.Document{
		Title:   title,
		Content: content,
		Source:  source,
	}

	query := `
		INSERT INTO documents (title, content, source, created_at)
		VALUES ($1, $2, $3, NOW())
		RETURNING id, created_at
	`

	err := r.db.DB.QueryRowxContext(
		ctx, query, document.Title, document.Content, document.Source,
	).Scan(&document.ID, &document.CreatedAt)

	if err != nil {
		return nil, err
	}

	return document, nil
}

// SearchDocuments ищет документы по запросу (полнотекстовый поиск)
func (r *DocumentRepo) SearchDocuments(ctx context.Context, query string, limit int) ([]models.Document, error) {
	// Используем полнотекстовый поиск PostgreSQL
	searchQuery := `
		SELECT id, title, content, source, created_at
		FROM documents
		WHERE to_tsvector('russian', content) @@ plainto_tsquery('russian', $1)
		ORDER BY ts_rank(to_tsvector('russian', content), plainto_tsquery('russian', $1)) DESC
		LIMIT $2
	`

	var documents []models.Document
	err := r.db.DB.SelectContext(ctx, &documents, searchQuery, query, limit)
	if err != nil {
		return nil, err
	}

	return documents, nil
}

// GetDocument получает документ по ID
func (r *DocumentRepo) GetDocument(ctx context.Context, id int64) (*models.Document, error) {
	query := `SELECT id, title, content, source, created_at FROM documents WHERE id = $1`

	var document models.Document
	err := r.db.DB.GetContext(ctx, &document, query, id)
	if err != nil {
		return nil, err
	}

	return &document, nil
}
