package models

//Clients Таблица клиентов Initials -это инициалы
type Clients struct {
	ID       int    `sql:",pk"`
	Phone    int    `sql:",unique"`
	Initials string `sql:",notnull"`
	Status   bool

	CitiesCity string `sql:"on_delete:RESTRICT, on_update: CASCADE"`
	Cities     *Cities
}
