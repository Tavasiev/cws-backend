package models

//Workers Таблица работников автомоек
type Workers struct {
	ID          int    `sql:",pk"`
	Phone       int    `sql:",unique"`
	Initials    string `sql:",notnull"`
	Address     string
	BoothNumber int
	Status      bool

	CitiesCity string `sql:"on_delete:RESTRICT, on_update: CASCADE"`
	Cities     *Cities
}
