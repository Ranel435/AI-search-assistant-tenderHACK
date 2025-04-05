package llm

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"backend/internal/config"
)

// Request представляет запрос к LLM API
type Request struct {
	Prompt    string `json:"prompt"`
	Context   string `json:"context,omitempty"`
	MaxTokens int    `json:"max_tokens,omitempty"`
}

// Response представляет ответ от LLM API
type Response struct {
	GeneratedText string `json:"generated_text"`
	Error         string `json:"error,omitempty"`
}

// Service предоставляет функциональность для работы с LLM
type Service struct {
	client    *http.Client
	baseURL   string
	maxTokens int
	retries   int
}

// NewService создаёт новый сервис для взаимодействия с LLM
func NewService(cfg config.LLMConfig) *Service {
	return &Service{
		client: &http.Client{
			Timeout: time.Duration(cfg.Timeout) * time.Second,
		},
		baseURL:   cfg.URL,
		maxTokens: cfg.MaxTokens,
		retries:   3, // Количество повторных попыток
	}
}

// GenerateAnswer отправляет запрос к LLM API и получает ответ
func (s *Service) GenerateAnswer(ctx context.Context, query, context string) (string, error) {
	reqBody := Request{
		Prompt:    query,
		Context:   context,
		MaxTokens: s.maxTokens,
	}

	var lastErr error
	for attempt := 0; attempt < s.retries; attempt++ {
		reqJSON, err := json.Marshal(reqBody)
		if err != nil {
			return "", fmt.Errorf("failed to marshal request: %w", err)
		}

		req, err := http.NewRequestWithContext(
			ctx,
			http.MethodPost,
			s.baseURL+"/generate", // URL должен соответствовать эндпоинту FastAPI
			bytes.NewBuffer(reqJSON),
		)
		if err != nil {
			return "", fmt.Errorf("failed to create request: %w", err)
		}

		req.Header.Set("Content-Type", "application/json")

		resp, err := s.client.Do(req)
		if err != nil {
			return "", fmt.Errorf("failed to send request: %w", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
		}

		var result Response
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			return "", fmt.Errorf("failed to decode response: %w", err)
		}

		if result.Error != "" {
			return "", fmt.Errorf("LLM error: %s", result.Error)
		}

		if err == nil && resp.StatusCode == http.StatusOK {
			// Успешный запрос, обрабатываем ответ
			return result.GeneratedText, nil
		}

		lastErr = err
		time.Sleep(time.Duration(attempt+1) * 500 * time.Millisecond) // Увеличивающаяся задержка
	}

	if lastErr != nil {
		return "", fmt.Errorf("failed after %d attempts: %w", s.retries, lastErr)
	}

	return "", nil
}
