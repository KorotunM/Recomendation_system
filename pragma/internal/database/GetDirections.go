package database

import (
	"context"
	"log"
	"pragma/internal/models"
)

// GetDirections - возвращает направления для заданных факультетов с учётом фильтров
func GetDirections(facultyIDs []int, form string, totalScore int, budget, paid bool) []models.Direction {
	// Если нет факультетов для фильтрации, возвращаем пустой срез
	if len(facultyIDs) == 0 {
		return []models.Direction{}
	}

	query := `
		SELECT d.id, d.faculty_id, d.name, d.tuition_fee, d.min_score_budget, d.min_score_paid, d.budget_places, d.paid_places
		FROM directions d
		WHERE d.faculty_id = ANY($1)
	`

	args := []interface{}{facultyIDs}

	// Фильтрация по форме обучения
	if form != "all" && form != "" {
		query += " AND d.name ILIKE $2"
		args = append(args, "%"+form+"%")
	}

	// Фильтрация по бюджету
	if budget {
		query += " AND d.min_score_budget <= $3 AND d.budget_places > 0"
		args = append(args, totalScore)
	}

	// Фильтрация по платному
	if paid {
		query += " AND d.min_score_paid <= $4 AND d.paid_places > 0"
		args = append(args, totalScore)
	}

	// Выполнение запроса
	rows, err := DB.Query(context.Background(), query, args...)
	if err != nil {
		log.Printf("Ошибка при выполнении запроса к направлениям: %v", err)
		return nil
	}
	defer rows.Close()

	var directions []models.Direction

	// Сканируем данные
	for rows.Next() {
		var dir models.Direction
		err := rows.Scan(&dir.ID, &dir.FacultyID, &dir.Name, &dir.TuitionFee, &dir.MinScoreBudget, &dir.MinScorePaid, &dir.BudgetPlaces, &dir.PaidPlaces)
		if err != nil {
			log.Printf("Ошибка при сканировании направления: %v", err)
			continue
		}
		directions = append(directions, dir)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Ошибка при обработке строк направлений: %v", err)
	}

	return directions
}
