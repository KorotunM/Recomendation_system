package database

import (
	"database/sql"
	"log"

	// Импорт драйвера для SQLite
	_ "modernc.org/sqlite"
)

// DB - глобальная переменная для доступа к базе данных.
var DB *sql.DB

// InitDatabase - функция для инициализации подключения к базе данных.
func ConnectDB() {
	var err error

	// Открываем подключение к базе данных
	DB, err = sql.Open("sqlite", "./data/RecomendationSystem.sqlite3") // Путь к базе данных
	if err != nil {
		log.Fatalf("Ошибка при открытии базы данных: %v", err)
	}
	// Проверка подключения
	err = DB.Ping()
	if err != nil {
		log.Fatalf("Ошибка при проверке подключения к базе данных: %v", err)
	}

	log.Println("Подключение к базе данных успешно выполнено.")

	// Создание таблицы universities, если она не существует
}

// CloseDatabase - функция для закрытия подключения к базе данных.
func CloseDB() {
	if DB != nil {
		err := DB.Close()
		if err != nil {
			log.Printf("Ошибка при закрытии базы данных: %v", err)
		} else {
			log.Println("Подключение к базе данных успешно закрыто.")
		}
	}
}
