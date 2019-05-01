package handlers

import (
	"net/http"

	"github.com/labstack/echo"

	//local
	db "github.com/Tavasiev/cws-backend/dbconn"
	"github.com/Tavasiev/cws-backend/models"
)

type inputPhone struct {
	Phone int
	User string
}

 // CheckPhone Функция проверяет есть ли такой телефон в базе данных:
 // если нет нужно зарегестрироваться, 
 // а если есть нужно ввести пароль чтобы войти в свой аккаунт.
func CheckPhone(c echo.Context) error {

	var inputJSON inputPhone
	err := c.Bind(&inputJSON)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Wrong data")
	}
	if inputJSON.User == "Client" {

		var Client models.Clients
		db.Conn.Model(&Client).Where("phone = ?",inputJSON.Phone).Select()
		return echo.NewHTTPError(http.StatusOK, Client.Phone != inputJSON.Phone)

	} else if inputJSON.User == "Worker" {

		var Worker models.Workers
		db.Conn.Model(&Worker).Where("phone = ?",inputJSON.Phone).Select()
		return echo.NewHTTPError(http.StatusOK, Worker.Phone != inputJSON.Phone)

	} else {
		return echo.NewHTTPError(http.StatusInternalServerError, "Wrong data")
	}

	
}
