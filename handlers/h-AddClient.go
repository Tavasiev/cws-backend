package handlers

import (
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/labstack/echo"

	//local
	db "github.com/Tavasiev/cws-backend/dbconn"
	"github.com/Tavasiev/cws-backend/models"
)

//
// AddClient Добавляет инф о клиенте в таблицу Clients
// формат входного json'а:
//{
//	"phone": 89888794747,
//	"pass" : qwerty1 or md5
//	"initials": "Ivanon I. I.",
//	"city": "Vladikavkaz"
//}
func AddClient(c echo.Context) error {

	var inputJSON models.Clients
	var login models.LoginResponse
	err := c.Bind(&inputJSON)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Wrong data")
	}

	uuid, err := uuid.NewV4()
	if err != nil {
		return err
	}

	err = db.Conn.Insert(&models.Clients{
		UUID:       uuid.String(),
		Phone:      inputJSON.Phone,
		Initials:   inputJSON.Initials,
		CitiesCity: inputJSON.CitiesCity,
		Password:   inputJSON.Password,
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

	err = login.GenerateJWT(inputJSON)
	if err != nil {
		return err
	}

	return echo.NewHTTPError(http.StatusOK, login)
}
