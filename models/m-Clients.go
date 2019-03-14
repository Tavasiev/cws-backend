package models

//Clients Таблица клиентов
type Clients struct {
	ID       int    `sql:", pk"`
	Phone    int    `sql:", unique, notnull"`
	Initials string `sql:",notnull"`
	Status   bool   `sql:",notnull"`

	CitiesCity string `sql:"on_delete:RESTRICT, on_update: CASCADE, notnull"`
	Cities     *Cities
}
