package database

import (
	"log"
	"pragma/internal/models"
)

func GetSubject() []models.Subject {
	pattern := "SELECT * FROM subjects"

	// Получаем все строки из таблицы Subject
	rows, err := DB.Query(pattern)
	if err != nil {
		log.Printf("Ошибка выполнения запроса %v\n", err)
		return nil
	}
	defer rows.Close()

	// Приводим их к виду структуры Subject
	var subjects []models.Subject
	for rows.Next() {
		var sub models.Subject
		err := rows.Scan(&sub.Id, &sub.Name)
		if err != nil {
			log.Printf("Ошибка при чтении данных из строки: %v\n", err)
			continue
		}
		subjects = append(subjects, sub)
	}
	// Проверка на наличие ошибок, возникших во время итерации
	if err = rows.Err(); err != nil {
		log.Printf("Ошибка при обработке строк: %v", err)
	}
	return subjects
}
