package search

import (
	"context"
	"strings"

	"backend/internal/db/models"
	"backend/internal/db/postgres"
)

// Service предоставляет функциональность для поиска в базе знаний
type Service struct {
	documentRepo postgres.DocumentRepository
}

// NewService создаёт новый сервис для поиска
func NewService(db *postgres.DB) *Service {
	return &Service{
		documentRepo: postgres.NewDocumentRepo(db),
	}
}

// Search ищет релевантные документы по запросу
func (s *Service) Search(ctx context.Context, query string) ([]models.Document, error) {
	// Предобработка запроса: удаление лишних пробелов, приведение к нижнему регистру
	query = strings.TrimSpace(strings.ToLower(query))
	if query == "" {
		return nil, nil
	}

	// Поиск документов (максимум 5)
	return s.documentRepo.SearchDocuments(ctx, query, 5)
}

// PrepareContext подготавливает контекст из найденных документов
func (s *Service) PrepareContext(docs []models.Document) string {
	if len(docs) == 0 {
		return ""
	}

	var sb strings.Builder

	for i, doc := range docs {
		if i > 0 {
			sb.WriteString("\n\n")
		}

		sb.WriteString("Документ: ")
		sb.WriteString(doc.Title)
		sb.WriteString("\nИсточник: ")
		sb.WriteString(doc.Source)
		sb.WriteString("\nСодержание: ")

		// Если текст слишком длинный, обрезаем его
		content := doc.Content
		if len(content) > 1000 {
			content = content[:1000] + "..."
		}
		sb.WriteString(content)
	}

	return sb.String()
}

// GetSourceReferences извлекает ссылки на источники из документов
func (s *Service) GetSourceReferences(docs []models.Document) []string {
	sources := make([]string, 0, len(docs))
	for _, doc := range docs {
		sources = append(sources, doc.Source)
	}
	return sources
}
