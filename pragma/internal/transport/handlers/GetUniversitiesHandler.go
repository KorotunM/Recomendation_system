package handlers

import (
	"html/template"
	"log"
	"net/http"
	"pragma/internal/database"
	"pragma/internal/models"
	"pragma/internal/service"
	"strconv"
)

func GetUniversitiesHandler(w http.ResponseWriter, r *http.Request) {
	// Получение параметров из запроса
	city := r.URL.Query().Get("city")
	form := r.URL.Query().Get("form")
	totalscoreStr := r.URL.Query().Get("totalScore")
	totalscore, _ := strconv.Atoi(totalscoreStr)
	dormitory := r.URL.Query().Get("dormitory") == "yes"
	military := r.URL.Query().Get("military") == "yes"
	budget := r.URL.Query().Get("budget") == "yes"
	paid := r.URL.Query().Get("paid") == "yes"

	offsetParam := r.URL.Query().Get("offset")
	offset, err := strconv.Atoi(offsetParam)
	if err != nil || offset < 0 {
		offset = 0
	}

	// Получение выбранных предметов
	subjectIDs := r.URL.Query()["subjects"]
	selectedSubjects := []int{}

	// Конвертация предметов в int
	for _, idStr := range subjectIDs {
		id, err := strconv.Atoi(idStr)
		if err == nil {
			selectedSubjects = append(selectedSubjects, id)
		}
	}

	log.Printf("Offset: %d", offset)

	// Валидация данных
	errors := service.ValidateInput(city, totalscoreStr)
	if service.HasErrors(errors) {
		subjects := database.GetSubject()

		pageData := models.PageData{
			Universities: nil,
			Subjects:     subjects,
			Errors:       errors,
			City:         city,
			Form:         form,
			TotalScore:   totalscore,
			Dormitory:    dormitory,
			Military:     military,
			Budget:       budget,
			Paid:         paid,
		}

		// Загрузка шаблона
		tmpl, err := template.ParseFiles("./web/templates/index.html", "./web/templates/university.html")
		if err != nil {
			http.Error(w, "Не удалось загрузить шаблоны", http.StatusInternalServerError)
			log.Printf("Ошибка загрузки шаблона: %v", err)
			return
		}

		err = tmpl.ExecuteTemplate(w, "index", pageData)
		if err != nil {
			http.Error(w, "Ошибка выполнения шаблона", http.StatusInternalServerError)
			log.Printf("Ошибка выполнения шаблона: %v", err)
		}
		return
	}

	// Получение данных (университеты, факультеты, направления)
	universities, faculties, directions, hasMore := database.GetFilteredData(city, form, totalscore, dormitory, military, budget, paid, selectedSubjects, offset)

	subjects := database.GetSubject()

	// Формируем данные для шаблона
	pageData := models.PageData{
		Universities:     universities,
		Faculties:        faculties,
		Directions:       directions,
		Subjects:         subjects,
		City:             city,
		Form:             form,
		TotalScore:       totalscore,
		Dormitory:        dormitory,
		Military:         military,
		Budget:           budget,
		Paid:             paid,
		SelectedSubjects: selectedSubjects,
		Offset:           offset,
		HasMore:          hasMore,
	}

	// Загрузка шаблона
	tmpl, err := template.ParseFiles("./web/templates/index.html", "./web/templates/university.html")
	if err != nil {
		http.Error(w, "Не удалось загрузить шаблоны", http.StatusInternalServerError)
		log.Printf("Ошибка загрузки шаблона: %v", err)
		return
	}

	err = tmpl.ExecuteTemplate(w, "index", pageData)
	if err != nil {
		http.Error(w, "Ошибка выполнения шаблона", http.StatusInternalServerError)
		log.Printf("Ошибка выполнения шаблона: %v", err)
	}
}
