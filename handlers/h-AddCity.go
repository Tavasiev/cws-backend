package handlers

import (
	"net/http"

	"github.com/labstack/echo"

	//local
	"github.com/Tavasiev/cws-backend/configs"
	"github.com/Tavasiev/cws-backend/dbconn"
	"github.com/Tavasiev/cws-backend/models"
)

////
// AddCity Добавляет город в таблицу Cities
// формат входного json'а:
// {
// 	"city": "YourCity"
// }
func AddCity(c echo.Context) error {
	configs.InitConfigs("configs/config") // получение конфиг структуры

	var inputJSON models.Cities
	err := c.Bind(&inputJSON)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Wrong data")
	}

	db := dbconn.GetConnect()

	err = db.Insert(&models.Cities{
		City: inputJSON.City,
	})

	if err != nil {
		return echo.NewHTTPError(http.StatusOK, err.Error())
	}

	return echo.NewHTTPError(http.StatusOK, "City Added")
}
