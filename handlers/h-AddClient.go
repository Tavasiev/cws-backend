package handlers

import (
	"log"
	"net/http"

	"github.com/labstack/echo"

	//local
	"github.com/Tavasiev/cws-backend/configs"
	"github.com/Tavasiev/cws-backend/dbconn"
	"github.com/Tavasiev/cws-backend/models"
)

// AddClient Добавляет инф о клиенте в таблицу Clients
// формат входного json'а:
//{
//	"phone": 89888794747,
//	"pass" : qwerty1 or md5
//	"initials": "Ivanon I. I.",
//	"city": "Vladikavkaz"
//}
func AddClient(c echo.Context) error {
	configs.InitConfigs("configs/config") // получение конфиг структуры

	var inputJSON models.Clients
	err := c.Bind(&inputJSON)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Wrong data")
	}

	db := dbconn.GetConnect()

	err = db.Insert(&models.Clients{

		Phone:      inputJSON.Phone,
		Initials:   inputJSON.Initials,
		CitiesCity: inputJSON.CitiesCity,
		Password:   inputJSON.Password,
	})

	if err != nil {
		return echo.NewHTTPError(http.StatusOK, err.Error())
	}

	token, err := models.CreateJwtToken()
	if err != nil {
		log.Println("Error creating Jwt token", err)
		return c.String(http.StatusInternalServerError, "something went wrong")
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "You registered sucsesfully",
		"token":   token,
	})
}
