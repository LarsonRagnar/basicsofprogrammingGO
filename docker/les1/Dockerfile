# ============================================
# Stage 1: Build (сборка приложения)
# ============================================
FROM golang:1.21-alpine AS builder

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем файлы зависимостей (для кэширования слоя)
COPY go.mod go.sum ./

# Загружаем зависимости (этот слой кэшируется, если go.mod/go.sum не изменились)
RUN go mod download

# Копируем исходный код
COPY *.go ./

# Собираем бинарник с оптимизациями:
# CGO_ENABLED=0 - отключаем CGO для статической сборки
# -ldflags="-w -s" - убираем отладочную информацию и символы (уменьшает размер)
# -trimpath - убирает пути из бинарника
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags="-w -s" \
    -trimpath \
    -o /app/server \
    .

# ============================================
# Stage 2: Runtime (финальный образ)
# ============================================
# Используем alpine образ - минимальный размер (~5MB)
FROM alpine:latest

# Создаем nonroot пользователя для безопасности
RUN addgroup -g 1000 appgroup && \
    adduser -D -u 1000 -G appgroup appuser

# Копируем только собранный бинарник из stage 1
COPY --from=builder /app/server /server

# Устанавливаем права на выполнение
RUN chmod +x /server

# Открываем порт (информация для документации)
EXPOSE 8080

# Переключаемся на nonroot пользователя
USER appuser

# Запускаем приложение
ENTRYPOINT ["/server"]
