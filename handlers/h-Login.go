package handlers

import (
	"net/http"

	"github.com/labstack/echo"

	//local
	db "github.com/Tavasiev/cws-backend/dbconn"
	"github.com/Tavasiev/cws-backend/models"
)

type inputPass struct {
	Password string
	Phone int
	User string
}

// Login Функция проверяет правильность логина и пароля, а так же выдаёт токен 
func Login(c echo.Context) error {

	var inputJSON inputPass
	err := c.Bind(&inputJSON)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Wrong data")
	}

	if inputJSON.User == "Client" {

		var Client models.Clients
		err = db.Conn.Model(&Client).Where("phone = ?",inputJSON.Phone).Select()
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}

		if inputJSON.Password == Client.Password {
			return echo.NewHTTPError(http.StatusOK, Client.Password) // JWT.AuthenticateUser(Client))
		}
	} else if inputJSON.User == "Worker" {

		var Worker models.Workers
		err = db.Conn.Model(&Worker).Where("phone = ?",inputJSON.Phone).Select()
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}
		if inputJSON.Password == Worker.Password {
			return echo.NewHTTPError(http.StatusOK, Worker.Password) //JWT.AuthenticateUser(Worker))
		}
	}
}