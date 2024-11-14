package main

import (
	"log"
	"net/http"
	"pragma/internal/database"
	"pragma/internal/transport"
)

func main() {
	// Инициализация базы данных
	database.InitDatabase()
	defer database.CloseDatabase()

	// Настройка маршрутов
	transport.SetupRoutes()

	// Запуск сервера
	log.Println("Запуск сервера на http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Ошибка при запуске сервера: %v", err)
	}
}
