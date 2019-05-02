package models

//Clients Таблица клиентов
type Clients struct {
	ID       int    `sql:", pk"`
	Phone    int    `sql:", unique, notnull" json:"phone"`
	Password string `sql:",notnull" json:"pass"`
	Initials string `sql:",notnull" json:"initials"`

	CitiesCity string `sql:"on_delete:RESTRICT, on_update: CASCADE, notnull" json:"city"`
	Cities     *Cities
}
