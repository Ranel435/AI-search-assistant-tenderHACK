FROM golang:1.21-alpine AS build

WORKDIR /app

# Копирование и скачивание зависимостей
COPY go.mod go.sum ./
RUN go mod download

# Копирование исходного кода
COPY . .

# Установка Swag
RUN go install github.com/swaggo/swag/cmd/swag@v1.16.1

# Отключить Go proxy из-за проблем с сетью
ENV GOPROXY=direct

# Генерация Swagger документации
RUN swag init -g cmd/server/main.go

# Сборка приложения
RUN go build -o app cmd/server/main.go

# Финальный образ
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Копирование бинарника из предыдущего шага
COPY --from=build /app/app .
COPY --from=build /app/migrations ./migrations
COPY --from=build /app/docs ./docs

EXPOSE 8080

CMD ["./app"]
