package models

//Workers Таблица работников автомоек
type Workers struct {
	ID          int    `sql:",pk"`
	Phone       int    `sql:",unique"`
	Initials    string `sql:",notnull"`
	Address     string `sql:",notnull"`
	BoothNumber int    `sql:",notnull"`
	Status      bool   `sql:",notnull"`

	CitiesCity string `sql:"on_delete:RESTRICT, on_update: CASCADE"`
	Cities     *Cities
}
