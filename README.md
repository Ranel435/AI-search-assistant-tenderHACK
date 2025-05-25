# 🧠 AI Search Assistant — TenderHACK 2024

Интеллектуальный ассистент для поиска по тендерной документации. Проект разработан на хакатоне **TenderHACK**. Использует векторный поиск и LLM-модель для генерации ответов на естественном языке, позволяя быстро находить релевантную информацию в загруженных документах.
## ⚙️ Команда MISISsippi
- Хаметшин Ранэль backend
- Усков Максим frontend
- Проскурин Андрей backend/designer
- Филатов Артур ML
- Крылов Александр ML

## ⚙️ Презентация
- https://drive.google.com/file/d/1i3TCish8pSQuL2XMKTIhXJJnm6C0YYZO/view?usp=sharing

## ⚙️ Стек технологий

**Backend (Golang):**
- REST API с использованием `Gin`
- Сервисная архитектура
- PostgreSQL + миграции
- Векторный поиск + чат с LLM (обёртка в `internal/services/llm`)
- Swagger-документация

**Frontend (React + Vite):**
- Чат-интерфейс с историей сообщений
- Поддержка нескольких чатов
- Адаптивная верстка
- Подключение к backend через REST API

**DevOps:**
- Docker / Docker Compose
- Nginx как reverse proxy
- Makefile и bash-скрипты
- Автоматический импорт документов

## 🚀 Как запустить

1. Клонируйте проект:
   ```bash
   git clone https://github.com/Ranel435/AI-search-assistant-tenderHACK.git
   cd AI-search-assistant-tenderHACK
   ```

2. Соберите и запустите контейнеры:
   ```bash
   docker-compose up --build
   ```

3. Перейдите по адресу:
   - **Frontend**: `http://localhost`
   - **Swagger UI**: `http://localhost/api/docs/index.html`

4. Чтобы загрузить документы:
   ```bash
   docker exec -it <backend-container> go run scripts/import_documents.go
   ```

## 📂 Структура проекта

```
.
├── backend/         # API, логика, векторный поиск, чат
│   ├── internal/    # Сервисы, роутеры, БД, LLM
│   ├── cmd/         # Точка входа main.go
│   └── scripts/     # Импорт документов
├── frontend/        # Vite + React SPA
│   └── src/         # UI-компоненты и логика чата
├── nginx/           # Конфиг nginx
├── docker-compose.yml
```

## 📚 Возможности

- 🔍 Поиск по загруженным PDF-документам
- 💬 Поддержка диалога в стиле ChatGPT
- 🧠 Генерация ответов LLM на основе семантического поиска
- 🗂 Управление несколькими чатами
- 🔧 Расширяемая архитектура (можно легко заменить модель/векторный движок)

## 🧪 Swagger

Файл `backend/docs/swagger.yaml` описывает все API endpoint'ы и может быть просмотрен в браузере после запуска по адресу:  
📄 `http://localhost/api/docs/index.html`

## 📝 Лицензия

Проект распространяется под лицензией **Unlicense** — вы можете свободно использовать, копировать и изменять код.
