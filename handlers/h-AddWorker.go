package handlers

import (
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/labstack/echo"

	//local
	db "github.com/Tavasiev/cws-backend/dbconn"
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

	var inputJSON models.Workers
	var login models.LoginResponse

	err := c.Bind(&inputJSON)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Wrong data")
	}

	uuid, err := uuid.NewV4()
	if err != nil {
		return err
	}

	err = db.Conn.Insert(&models.Workers{
		UUID:        uuid.String(),
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
	}

	err = models.ExpireUserTokens(uuid.String())
	if err != nil {
		return err
	}

	err = login.NewRefreshToken(uuid.String())
	if err != nil {
		return err
	}

	err = login.GenerateJWTWorker(inputJSON)
	if err != nil {
		return err
	}

	return echo.NewHTTPError(http.StatusOK, login)
}
