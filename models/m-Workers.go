package models

//Workers Таблица работников автомоек
type Workers struct {
	ID          int    `sql:",pk"`
	UUID        string `sql:", unique" json:"uuid"`
	Phone       int    `sql:",unique" json:"phone"`
	Password    string `sql:",notnull" json:"pass"`
	Initials    string `sql:",notnull" json:"initials"`
	Address     string `sql:",notnull" json:"addr"`
	BoothNumber int    `sql:",notnull" json:"booth_number"`
	WoshTitle   string `json:"wosh_title"`
	Status      bool   `sql:",notnull, default:false"`

	CitiesCity string `sql:"on_delete:RESTRICT, on_update: CASCADE" json:"city"`
	Cities     *Cities
}
