package database

import (
	"log"
	"pragma/internal/models"
)

// GetFilteredData - объединяет данные об университетах, факультетах и направлениях
func GetFilteredData(city, form string, totalScore int, dormitory, military, budget, paid bool, selectedSubjects []int) ([]models.University, []models.Faculty, []models.Direction) {

	// Получение университетов
	universities := GetFilteredUniversities(city, dormitory, military)

	// Извлекаем IDs университетов для фильтрации факультетов
	universityIDs := []int{}
	for _, uni := range universities {
		universityIDs = append(universityIDs, uni.ID)
	}

	// Получение факультетов для выбранных университетов
	faculties := GetFaculties(universityIDs, selectedSubjects)

	// Извлекаем IDs факультетов для фильтрации направлений
	facultyIDs := []int{}
	for _, fac := range faculties {
		facultyIDs = append(facultyIDs, fac.ID)
	}

	// Получение направлений для выбранных факультетов
	directions := GetDirections(facultyIDs, form, totalScore, budget, paid)

	// Логирование для отладки
	log.Printf("Universities: %d, Faculties: %d, Directions: %d", len(universities), len(faculties), len(directions))

	return universities, faculties, directions
}
