# Используем официальное минимальное изображение Go
FROM golang:1.23.2 as builder

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /cmd

# Копируем Go-модули и скачиваем зависимости
COPY go.mod go.sum ./
RUN go mod download

# Копируем весь исходный код проекта в контейнер
COPY . .

# Компилируем приложение
RUN go build -o main ./cmd/main.go

# Используем минимальный образ для запуска
FROM alpine:latest

# Устанавливаем нужные зависимости (например, SSL)
RUN apk --no-cache add ca-certificates

# Устанавливаем рабочую директорию
WORKDIR /root/

# Копируем скомпилированное приложение из стадии сборки
COPY --from=builder /cmd/main .
COPY --from=builder /cmd/web ./web
COPY --from=builder /cmd/data ./data

# Устанавливаем порт, который будет использовать приложение
EXPOSE 8080

# Команда для запуска приложения
CMD ["./main"]
