package handlers

import (
	"net/http"

	"github.com/labstack/echo"

	//local
	db "github.com/Tavasiev/cws-backend/dbconn"
	"github.com/Tavasiev/cws-backend/models"
)

type inputPhone struct {
	Phone int `json:"phone"`
	User string `json:"user"`
}

 // CheckPhone Функция проверяет есть ли такой телефон в базе данных:
 // если нет нужно зарегестрироваться, 
 // а если есть нужно ввести пароль чтобы войти в свой аккаунт.
 // формат входного json'а:
 //{
//	"phone": 89888794747,
//	"user": "Client"
//}
func CheckPhone(c echo.Context) error {

	var inputJSON inputPhone
	err := c.Bind(&inputJSON)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Wrong data")
	}
	if inputJSON.User == "Client" {

		var Client models.Clients
		_, err = db.Conn.Query(&Client,"SELECT * FROM clients WHERE phone = ?",inputJSON.Phone)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}
		return echo.NewHTTPError(http.StatusOK, Client.Phone != inputJSON.Phone)

	} else if inputJSON.User == "Worker" {

		var Worker models.Workers
		_, err = db.Conn.Query(&Worker,"SELECT * FROM clients WHERE phone = ?",inputJSON.Phone)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}
		return echo.NewHTTPError(http.StatusOK, Worker.Phone != inputJSON.Phone)

	} else {
		return echo.NewHTTPError(http.StatusInternalServerError, "Wrong data")
	}

	
}
