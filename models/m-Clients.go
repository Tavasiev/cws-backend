package models

//Clients Таблица клиентов
type Clients struct {
	ID       int    `sql:", pk"`
	Phone    int    `sql:", unique, notnull" json:"phone"`
	Initials string `sql:",notnull" json:"initials"`
	Status   bool   `sql:",notnull, default:false"`

	CitiesCity string `sql:"on_delete:RESTRICT, on_update: CASCADE, notnull" json:"city"`
	Cities     *Cities
}
