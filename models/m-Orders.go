package models

import "time"

//Orders Таблица Заказов
type Orders struct {
	ID     int       `sql:",pk"`
	Date   time.Time `sql:"default:now()"`
	Rating string
	Status string

	WorkersID int `sql:"on_delete:RESTRICT, on_update: CASCADE"`
	Workers   *Workers

	ClientsID int `sql:"on_delete:RESTRICT, on_update: CASCADE"`
	Clients   *Clients
}
