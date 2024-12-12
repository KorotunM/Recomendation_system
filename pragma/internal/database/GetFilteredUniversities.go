package database

import (
	"log"
	"pragma/internal/models"

	// Импорт драйвера SQLite

	_ "modernc.org/sqlite"
)

// GetFilteredUniversities возвращает все университеты из базы данных
func GetFilteredUniversities(city, specialty, form string, totalScore int, dormitory, military, budget, paid string) []models.University {
	query := `SELECT * FROM universities WHERE 1=1`
	args := []interface{}{}

	// Добавление фильтров, если они указаны
	if city != "" {
		query += " AND city LIKE ?"
		args = append(args, "%"+city+"%")
	}
	if specialty != "" {
		query += " AND specialty LIKE ?"
		args = append(args, "%"+specialty+"%")
	}
	if form != "" && form != "all" {
		query += " AND type LIKE ?"
		args = append(args, form)
	}
	if totalScore != 0 {
		query += " AND min_score <= ?"
		args = append(args, totalScore)
	}
	if dormitory == "yes" {
		query += " AND has_dormitory = 1"
	}
	if military == "yes" {
		query += " AND has_military = 1"
	}
	if budget == "yes" {
		query += " AND budget_places > 0"
	}
	if paid == "yes" {
		query += " AND paid_places > 0"
	}

	// Выполнение SQL-запроса
	rows, err := DB.Query(query, args...)
	if err != nil {
		log.Printf("Ошибка при выполнении запроса: %v", err)
		return nil
	}
	defer rows.Close()

	// Срез для хранения университетов
	var universities []models.University

	// Обход строк результата запроса и заполнение структур University
	for rows.Next() {
		var uni models.University
		err := rows.Scan(&uni.ID, &uni.Name, &uni.City, &uni.Type, &uni.Specialty, &uni.TuitionFee, &uni.MinScore, &uni.BudgetPlaces, &uni.PaidPlaces, &uni.HasDormitory, &uni.HasMilitary)
		if err != nil {
			log.Printf("Ошибка при чтении данных из строки: %v", err)
			continue
		}
		universities = append(universities, uni)
	}

	// Проверка на наличие ошибок, возникших во время итерации
	if err = rows.Err(); err != nil {
		log.Printf("Ошибка при обработке строк: %v", err)
	}

	return universities
}
