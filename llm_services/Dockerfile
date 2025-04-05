FROM python:3.10-slim

WORKDIR /app

# Установка зависимостей
RUN pip install --no-cache-dir fastapi uvicorn transformers torch accelerate

# Копирование исходного кода
COPY app.py /app/

# Копирование модели
COPY llm-model /app/llm-model

# Порт для API
EXPOSE 8000

# Запуск сервера
CMD ["python", "app.py"]
