package handlers

import (
	"net/http"

	"github.com/go-pg/pg"
	"github.com/labstack/echo"

	//local
	"github.com/Tavasiev/cws-backend/configs"
	"github.com/Tavasiev/cws-backend/models"
)

// AddCity Добавляет город в таблицу Cities
// формат входного json'а:
// {
// 	"city": "YourCity"
// }
func AddCity(c echo.Context) error {
	conf := configs.MakeConfig() // получение конфиг структуры

	var inputJSON models.Cities
	err := c.Bind(&inputJSON)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Wrong data")
	}

	db := pg.Connect(&pg.Options{
		Addr:     conf.DB.Addr,
		User:     conf.DB.User,
		Password: conf.DB.Password,
		Database: conf.DB.Database,
	})
	defer db.Close()

	err = db.Insert(&models.Cities{
		City: inputJSON.City,
	})

	if err != nil {
		panic(err)
	}

	return echo.NewHTTPError(http.StatusOK, "Added")
}
