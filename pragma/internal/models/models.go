package models

// Университет
type University struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	City         string `json:"city"`
	HasDormitory bool   `json:"has_dormitory"` // Наличие общежития
	HasMilitary  bool   `json:"has_military"`  // Наличие военной кафедры
	Rating       int    `json:"rating"`        // Рейтинг вуза
}

// Факультет
type Faculty struct {
	ID           int    `json:"id"`
	UniversityID int    `json:"university_id"`
	Name         string `json:"name"`
}

// Направление (Программа обучения)
type Direction struct {
	ID             int    `json:"id"`
	FacultyID      int    `json:"faculty_id"`
	Name           string `json:"name"`
	TuitionFee     int    `json:"tuition_fee"`
	MinScoreBudget int    `json:"min_score_budget"`
	MinScorePaid   int    `json:"min_score_paid"`
	BudgetPlaces   int    `json:"budget_places"`
	PaidPlaces     int    `json:"paid_places"`
}

// Предмет
type Subject struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Избранное (ВУЗы)
type Favorite struct {
	ID           int `json:"id"`
	UserID       int `json:"user_id"`
	UniversityID int `json:"university_id"`
}

// История поиска
type SearchHistory struct {
	ID         int    `json:"id"`
	UserID     int    `json:"user_id"`
	Filters    string `json:"filters"` // JSON строка с фильтрами
	SearchedAt string `json:"searched_at"`
}

// Структура для передачи данных в шаблоны
type PageData struct {
	Universities     []University `json:"universities"`
	Faculties        []Faculty    `json:"faculties"`
	Directions       []Direction  `json:"directions"`
	Subjects         []Subject    `json:"subjects"`
	Errors           Errors       `json:"errors"`
	City             string       `json:"city"`
	Form             string       `json:"form"`
	TotalScore       int          `json:"total_score"`
	Dormitory        bool         `json:"dormitory"`
	Military         bool         `json:"military"`
	Budget           bool         `json:"budget"`
	Paid             bool         `json:"paid"`
	SelectedSubjects []int        `json:"selected_subjects"`
	Offset           int
	HasMore          bool
}

// Ошибки валидации
type Errors struct {
	CityError  string `json:"city_error"`
	ScoreError string `json:"score_error"`
}
