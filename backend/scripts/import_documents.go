package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"backend/internal/config"
	"backend/internal/db/postgres"
)

// Простая функция для импорта документов в базу знаний
func main() {
	// Загрузка конфигурации
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Подключение к БД
	db, err := postgres.NewDB(cfg.DB)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Создаем репозиторий для работы с документами
	docRepo := postgres.NewDocumentRepo(db)

	// Путь к директории с файлами документов
	docsDir := "data/documents"
	if len(os.Args) > 1 {
		docsDir = os.Args[1]
	}

	// Обходим все файлы в директории
	err = filepath.Walk(docsDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Пропускаем директории
		if info.IsDir() {
			return nil
		}

		// Пропускаем не текстовые файлы
		if !strings.HasSuffix(strings.ToLower(path), ".txt") {
			return nil
		}

		// Читаем содержимое файла
		content, err := ioutil.ReadFile(path)
		if err != nil {
			log.Printf("Error reading file %s: %v", path, err)
			return nil
		}

		// Определяем заголовок и источник
		title := filepath.Base(path)
		source := fmt.Sprintf("Файл: %s", title)

		// Сохраняем документ в базу
		_, err = docRepo.CreateDocument(context.Background(), title, string(content), source)
		if err != nil {
			log.Printf("Error creating document %s: %v", path, err)
			return nil
		}

		log.Printf("Imported document: %s", path)
		return nil
	})

	if err != nil {
		log.Fatalf("Error walking through docs directory: %v", err)
	}

	log.Println("Document import completed successfully")
}
