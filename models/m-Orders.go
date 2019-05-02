package models

import "time"

//Orders Таблица Заказов
type Orders struct {
	ID          int       `sql:",pk"`
	Date        time.Time `sql:"default:now()"`
	Rating      int
	Status      string
	ClientPhone int `json:"client_phone"`
	WorkerPhone int `json:"worker_phone"`

	WorkersID int `sql:"on_delete:RESTRICT, on_update: CASCADE"`
	Workers   *Workers

	ClientsID int `sql:"on_delete:RESTRICT, on_update: CASCADE"`
	Clients   *Clients
}
