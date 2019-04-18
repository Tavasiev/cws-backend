package handlers

import (
	"net/http"

	"github.com/labstack/echo"

	//local
	"github.com/Tavasiev/cws-backend/configs"
	"github.com/Tavasiev/cws-backend/dbconn"
	"github.com/Tavasiev/cws-backend/models"
)

// AddWorker Добавляет нового работника в таблицу Workers
// формат входного json'а:
//{
//	"phone": 89888794747,
//	"pass" : qwerty1
//	"initials": "Ivanon I. I.",
//	"address": "Mamsurova 42",
//	"booth_number": 7,
//	"city": "Vladikavkaz"
//  "wosh_title : "string"
//}
func AddWorker(c echo.Context) error {
	configs.InitConfigs("configs/config") // получение конфиг структуры

	var inputJSON models.Workers
	err := c.Bind(&inputJSON)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Wrong data")
	}

	db := dbconn.GetConnect()

	err = db.Insert(&models.Workers{

		Phone:       inputJSON.Phone,
		Password:    inputJSON.Password,
		Initials:    inputJSON.Initials,
		Address:     inputJSON.Address,
		BoothNumber: inputJSON.BoothNumber,
		CitiesCity:  inputJSON.CitiesCity,
		WoshTitle:   inputJSON.WoshTitle,
	})

	if err != nil {
		return echo.NewHTTPError(http.StatusOK, err.Error())
		//panic(err)
	}

	return echo.NewHTTPError(http.StatusOK, "Worker Added")
}
