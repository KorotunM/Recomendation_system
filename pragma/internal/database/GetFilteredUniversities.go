package database

import (
	"context"
	"log"
	"pragma/internal/models"
	"strconv"
)

const Limit = 4

// GetFilteredUniversities - получение университетов с фильтрацией и пагинацией
func GetFilteredUniversities(city string, dormitory, military bool, offset int) []models.University {
	query := `
		SELECT u.id, u.name, u.city, u.has_dormitory, u.has_military, r.rank_position
		FROM universities u
		LEFT JOIN university_ratings r ON u.id = r.university_id
		WHERE 1=1
	`

	args := []interface{}{}
	argIndex := 1

	// Фильтр по городу
	if city != "" {
		query += " AND u.city ILIKE $" + strconv.Itoa(argIndex)
		args = append(args, "%"+city+"%")
		argIndex++
	}

	// Фильтр по общежитию
	if dormitory {
		query += " AND u.has_dormitory = true"
	}

	// Фильтр по военной кафедре
	if military {
		query += " AND u.has_military = true"
	}

	// Сортировка по рейтингу
	query += " GROUP BY u.id, r.rank_position"
	query += " ORDER BY r.rank_position ASC NULLS LAST, u.name ASC"

	// Пагинация
	query += " LIMIT $1 OFFSET $2"
	args = append(args, Limit, offset)

	// Выполнение запроса
	rows, err := DB.Query(context.Background(), query, args...)
	if err != nil {
		log.Printf("Ошибка при выполнении запроса: %v", err)
		return nil
	}
	defer rows.Close()

	var universities []models.University

	// Сканирование данных
	for rows.Next() {
		var uni models.University
		var rankPosition *int

		err := rows.Scan(
			&uni.ID,
			&uni.Name,
			&uni.City,
			&uni.HasDormitory,
			&uni.HasMilitary,
			&rankPosition,
		)
		if err != nil {
			log.Printf("Ошибка при сканировании данных: %v", err)
			continue
		}

		if rankPosition != nil {
			uni.Rating = *rankPosition
		} else {
			uni.Rating = 0 // Университет без рейтинга
		}

		universities = append(universities, uni)
	}

	return universities
}
