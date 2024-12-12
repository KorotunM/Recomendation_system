package handlers

import (
	"html/template"
	"log"
	"net/http"
	"pragma/internal/database"
	"pragma/internal/models"
	"strconv"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	// Получение параметров из запроса
	city := r.URL.Query().Get("city")
	specialty := r.URL.Query().Get("specialty")
	form := r.URL.Query().Get("form")
	totalscore, _ := strconv.Atoi(r.URL.Query().Get("totalScore"))
	dormitory := r.URL.Query().Get("dormitory") == "yes"
	military := r.URL.Query().Get("military") == "yes"
	budget := r.URL.Query().Get("budget") == "yes"
	paid := r.URL.Query().Get("paid") == "yes"

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

	// Получение университетов с фильтрацией по параметрам
	universities := database.GetFilteredUniversities(city, specialty, form, totalscore, dormitoryStr, militaryStr, budgetStr, paidStr)
	// Получаем список сдаваемых предметов
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
