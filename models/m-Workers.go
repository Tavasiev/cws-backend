package models

//Workers Таблица работников автомоек
type Workers struct {
	ID          int    `sql:",pk"`
	Phone       int    `sql:",unique" json:"phone"`
	Initials    string `sql:",notnull" json:"initials"`
	Address     string `sql:",notnull" json:"address"`
	BoothNumber int    `sql:",notnull" json:"booth_number"`
	Status      bool   `sql:",notnull, default:false"`

	CitiesCity string `sql:"on_delete:RESTRICT, on_update: CASCADE" json:"city"`
	Cities     *Cities
}
