FROM golang:1.21-alpine AS build

WORKDIR /app

# Копирование и скачивание зависимостей
COPY go.mod go.sum ./
RUN go mod download

# Копирование исходного кода
COPY . .

# Компиляция
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/server

# Финальный образ
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Копирование бинарника из предыдущего шага
COPY --from=build /app/main .
COPY --from=build /app/migrations ./migrations

EXPOSE 8080

CMD ["./main"]
