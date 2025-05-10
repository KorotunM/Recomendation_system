package database

import (
	"context"
	"log"
	"pragma/internal/models"
)

// GetFilteredUniversities - возвращает университеты с учётом базовых фильтров
func GetFilteredUniversities(city string, dormitory, military bool) []models.University {
	query := `
		SELECT u.id, u.name, u.city, u.has_dormitory, u.has_military, r.rank_position
		FROM universities u
		LEFT JOIN university_ratings r ON u.id = r.university_id
		WHERE 1=1
	`

	args := []interface{}{}

	// Фильтр по городу
	if city != "" {
		query += " AND u.city ILIKE $1"
		args = append(args, "%"+city+"%")
	}

	// Фильтр по общежитию
	if dormitory {
		query += " AND u.has_dormitory = true"
	}

	// Фильтр по военной кафедре
	if military {
		query += " AND u.has_military = true"
	}

	query += " GROUP BY u.id, r.rank_position"

	// Выполнение запроса
	rows, err := DB.Query(context.Background(), query, args...)
	if err != nil {
		log.Printf("Ошибка при выполнении запроса к университетам: %v", err)
		return nil
	}
	defer rows.Close()

	var universities []models.University

	// Сканируем результат
	for rows.Next() {
		var uni models.University
		var rankPosition int

		err := rows.Scan(&uni.ID, &uni.Name, &uni.City, &uni.HasDormitory, &uni.HasMilitary, &rankPosition)
		if err != nil {
			log.Printf("Ошибка при сканировании университета: %v", err)
			continue
		}

		uni.Rating = rankPosition
		universities = append(universities, uni)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Ошибка при обработке строк университета: %v", err)
	}

	return universities
}
