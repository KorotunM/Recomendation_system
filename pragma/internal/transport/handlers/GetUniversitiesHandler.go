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
	specialty := r.URL.Query().Get("specialty")
	totalscoreStr := r.URL.Query().Get("totalScore")
	totalscore, _ := strconv.Atoi(r.URL.Query().Get("totalScore"))
	form := r.URL.Query().Get("form")
	dormitory := r.URL.Query().Get("dormitory") == "yes"
	military := r.URL.Query().Get("military") == "yes"
	budget := r.URL.Query().Get("budget") == "yes"
	paid := r.URL.Query().Get("paid") == "yes"

	// Определение значений строковых параметров
	dormitoryStr := "no"
	if dormitory {
		dormitoryStr = "yes"
	}

	militaryStr := "no"
	if military {
		militaryStr = "yes"
	}

	budgetStr := "no"
	if budget {
		budgetStr = "yes"
	}

	paidStr := "no"
	if paid {
		paidStr = "yes"
	}

	errors := service.ValidateInput(city, specialty, totalscoreStr)
	if service.HasErrors(errors) {
		// Если есть ошибки валидации, подготовим данные для отображения
		subjects := database.GetSubject()
		pageData := models.PageData{
			Universities: nil, // Пустой список университетов
			Subjects:     subjects,
			Errors:       errors, // Передача ошибок в шаблон
			City:         city,
			Specialty:    specialty,
			Form:         form,
			TotalScore:   totalscore,
			Dormitory:    dormitory,
			Military:     military,
			Budget:       budget,
			Paid:         paid,
		}

		// Загрузка шаблонов
		tmpl, err := template.ParseFiles("./web/templates/index.html", "./web/templates/university.html")
		if err != nil {
			http.Error(w, "Не удалось загрузить шаблоны", http.StatusInternalServerError)
			log.Printf("Ошибка загрузки шаблона: %v", err)
			return
		}

		// Выполнение шаблона с ошибками
		err = tmpl.ExecuteTemplate(w, "index", pageData)
		if err != nil {
			http.Error(w, "Ошибка выполнения шаблона", http.StatusInternalServerError)
			log.Printf("Ошибка выполнения шаблона: %v", err)
		}
		return
	}

	// Получение университетов с фильтрацией по параметрам
	// Если фильтры пустые, то возвращаем все университеты
	universities := database.GetFilteredUniversities(city, specialty, form, totalscore, dormitoryStr, militaryStr, budgetStr, paidStr)

	subjects := database.GetSubject()

	// Создание структуры данных для страницы
	pageData := models.PageData{
		Universities: universities,
		Subjects:     subjects,
		City:         city,
		Specialty:    specialty,
		Form:         form,
		TotalScore:   totalscore,
		Dormitory:    dormitory,
		Military:     military,
		Budget:       budget,
		Paid:         paid,
	}

	// Создание нового шаблона и загрузка всех необходимых файлов
	tmpl, err := template.ParseFiles("./web/templates/index.html", "./web/templates/university.html")
	if err != nil {
		http.Error(w, "Не удалось загрузить шаблоны", http.StatusInternalServerError)
		log.Printf("Ошибка загрузки шаблона: %v", err)
		return
	}

	// Выполнение шаблона "index" и передача данных об университетах
	err = tmpl.ExecuteTemplate(w, "index", pageData)
	if err != nil {
		http.Error(w, "Ошибка выполнения шаблона", http.StatusInternalServerError)
		log.Printf("Ошибка выполнения шаблона: %v", err)
	}
}
