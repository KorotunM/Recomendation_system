package models

type University struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	City         string `json:"city"`
	Type         string `json:"type"`          // Тип университета (например, государственный)
	Specialty    string `json:"specialty"`     // Направление (специальность)
	TuitionFee   int    `json:"tuition_fee"`   // Стоимость обучения
	MinScore     int    `json:"min_score"`     // Минимальный проходной балл
	BudgetPlaces int    `json:"budget_places"` // Количество бюджетных мест
	PaidPlaces   int    `json:"paid_places"`   // Количество платных мест
	HasDormitory bool   `json:"has_dormitory"` // Наличие общежития (да/нет)
	HasMilitary  bool   `json:"has_military"`  // Наличие военного учебного центра (да/нет)
}

type PageData struct {
	Universities []University
	Subjects     []Subject
	Errors       Errors
	City         string
	Specialty    string
	Form         string
	TotalScore   int
	Dormitory    bool
	Military     bool
	Budget       bool
	Paid         bool
}

type Subject struct {
	Id   int
	Name string
}

type Errors struct {
	CityError      string
	SpecialtyError string
	ScoreError     string
}
