package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

var DB *pgxpool.Pool

// ConnectDB - инициализация подключения к PostgreSQL
func ConnectDB() {
	// Загрузка переменных окружения из .env
	err := godotenv.Load()
	if err != nil {
		log.Println("Файл .env не найден, используем системные переменные")
	}

	// Получение данных для подключения
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	sslMode := os.Getenv("SSL_MODE")

	// Формирование строки подключения
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		dbUser, dbPassword, dbHost, dbPort, dbName, sslMode)

	// Настройка пула соединений
	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		log.Fatalf("Ошибка при парсинге строки подключения: %v", err)
	}

	config.MaxConns = 10 // Максимальное количество подключений
	config.MinConns = 2  // Минимальное количество подключений
	config.MaxConnLifetime = 30 * time.Minute
	config.MaxConnIdleTime = 15 * time.Minute

	// Подключение к базе данных
	DB, err = pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Fatalf("Ошибка при подключении к базе данных: %v", err)
	}

	// Проверка подключения
	err = DB.Ping(context.Background())
	if err != nil {
		log.Fatalf("Ошибка при проверке подключения: %v", err)
	}

	log.Println("Подключение к базе данных PostgreSQL успешно установлено.")
}

// CloseDB - закрытие подключения к базе данных
func CloseDB() {
	if DB != nil {
		DB.Close()
		log.Println("Подключение к PostgreSQL успешно закрыто.")
	}
}
