package database

import (
	"database/sql"
	"log"

	// Импорт драйвера для SQLite
	_ "modernc.org/sqlite" // Используйте этот драйвер, если возникают проблемы с CGO
	// Если у вас установлен GCC и вы предпочитаете использовать go-sqlite3:
	// _ "github.com/mattn/go-sqlite3"
)

// DB - глобальная переменная для доступа к базе данных.
var DB *sql.DB

// InitDatabase - функция для инициализации подключения к базе данных.
func InitDatabase() {
	var err error

	// Открываем подключение к базе данных
	DB, err = sql.Open("sqlite", "./universities.db") // Путь к базе данных
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
	createTableQuery := `
    CREATE TABLE IF NOT EXISTS universities (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT,
        city TEXT,
        type TEXT,
        specialty TEXT,
        tuition_fee INTEGER,
        min_score INTEGER,
        budget_places INTEGER,
        paid_places INTEGER,
        has_dormitory BOOLEAN,
        has_military BOOLEAN
    );
    `
	_, err = DB.Exec(createTableQuery)
	if err != nil {
		log.Fatalf("Ошибка при создании таблицы universities: %v", err)
	}
	log.Println("Таблица universities успешно создана или уже существует.")
}

// CloseDatabase - функция для закрытия подключения к базе данных.
func CloseDatabase() {
	if DB != nil {
		err := DB.Close()
		if err != nil {
			log.Printf("Ошибка при закрытии базы данных: %v", err)
		} else {
			log.Println("Подключение к базе данных успешно закрыто.")
		}
	}
}
