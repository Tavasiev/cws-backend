package models

import "time"

//Orders Таблица Заказов
type Orders struct {
	ID     int       `sql:",pk"`
	Date   time.Time `sql:",notnull"`
	Rating string
	Status string `sql:",notnull"`

	WorkersID int `sql:"on_delete:RESTRICT, on_update: CASCADE, notnull"`
	Workers   *Workers

	ClientsID int `sql:"on_delete:RESTRICT, on_update: CASCADE, notnull"`
	Clients   *Clients
}
