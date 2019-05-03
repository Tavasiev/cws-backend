package handlers

import (
	"net/http"

	"github.com/labstack/echo"

	//local

	db "github.com/Tavasiev/cws-backend/dbconn"
	"github.com/Tavasiev/cws-backend/models"
)

type inputPass struct {
	Password string `json:"pass"`
	Phone    int    `json:"phone"`
	User     string `json:"user"`
}

// Login Функция проверяет правильность логина и пароля, а так же выдаёт токен
// формат входного json'а:
//{
//	"pass" : "qwerty1",
//	"phone": 89888794747,
//	"user": "Client"
//}
func Login(c echo.Context) error {

	var login models.LoginResponse
	var inputJSON inputPass

	err := c.Bind(&inputJSON)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Wrong data")
	}

	if inputJSON.User == "Client" {

		var Client models.Clients
		_, err = db.Conn.Query(&Client, "SELECT * FROM clients WHERE phone = ?", inputJSON.Phone)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}

		if inputJSON.Password == Client.Password {

			err = models.ExpireUserTokens(Client.ID)
			if err != nil {
				return err
			}

			err = login.NewRefreshToken(Client.ID)
			if err != nil {
				return err
			}

			err = login.GenerateJWT(Client)
			if err != nil {
				return err
			}

			return echo.NewHTTPError(http.StatusOK, login) // JWT.AuthenticateUser(Client))
		}
	} else if inputJSON.User == "Worker" {

		var Worker models.Workers
		_, err = db.Conn.Query(&Worker, "SELECT * FROM workers WHERE phone = ?", inputJSON.Phone)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}
		if inputJSON.Password == Worker.Password {
			return echo.NewHTTPError(http.StatusOK, Worker.Password) //JWT.AuthenticateUser(Worker))
		}
	}
	return echo.NewHTTPError(http.StatusOK, login)
}
