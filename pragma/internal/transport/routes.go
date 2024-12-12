// internal/transport/routes.go
package transport

import (
	"net/http"
	"pragma/internal/transport/handlers"
)

func SetupRoutes() {
	// Настройка маршрута для главной страницы
	http.HandleFunc("/", handlers.IndexHandler)

	// Настройка маршрута для отображения университетов
	http.HandleFunc("/universities", handlers.GetUniversitiesHandler)

	// Настройка маршрутов для статических файлов (CSS, изображения и т.д.)
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./web/css"))))
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./web/assets"))))
}
