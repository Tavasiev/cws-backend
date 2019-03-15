package handlers

import (
	"net/http"

	"github.com/Tavasiev/cws-backend/models"
	"github.com/go-pg/pg"
	"github.com/labstack/echo"
)

// AddCity Добавляет город в таблицу Cities
// формат входного json'а:
// {
// 	"city": "YourCity"
// }
func AddCity(c echo.Context) error {
	var inputJSON models.Cities
	err := c.Bind(&inputJSON)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Wrong data")
	}

	db := pg.Connect(&pg.Options{
		User:     "user",
		Password: "password",
		Database: "database",
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
