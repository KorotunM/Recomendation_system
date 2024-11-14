// internal/transport/routes.go
package transport

import (
	"net/http"
)

// SetupRoutes инициализирует маршруты для веб-приложения
func SetupRoutes() {
	// Настройка маршрута для главной страницы
	http.HandleFunc("/", IndexHandler)

	// Настройка маршрута для отображения университетов
	http.HandleFunc("/universities", GetUniversitiesHandler)

	// Настройка маршрутов для статических файлов (CSS, изображения и т.д.)
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./web/css"))))
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./web/assets"))))
}
