package postgres

import (
	"fmt"

	"backend/internal/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// DB обёртка над sqlx.DB
type DB struct {
	*sqlx.DB
}

// NewDB создаёт новое подключение к PostgreSQL
func NewDB(cfg config.DBConfig) (*DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode,
	)

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("could not connect to database: %w", err)
	}

	// Проверка соединения
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("could not ping database: %w", err)
	}

	return &DB{DB: db}, nil
}

// Close закрывает соединение с базой данных
func (db *DB) Close() error {
	return db.DB.Close()
}
