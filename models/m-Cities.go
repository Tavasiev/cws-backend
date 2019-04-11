package models

//Cities Таблица городов, в которых запущено приложение
type Cities struct {
	City string `sql:",pk, unique"`
}
