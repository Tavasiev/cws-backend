package handlers

import (
	"net/http"

	"github.com/labstack/echo"

	//local
	"github.com/Tavasiev/cws-backend/models"
)

// LoginRefresh принимает рефреш токен и возвращает токен, новый рефреш токен, id пользователя
// формат входного json'а:
//{
//	"refresh" : "some refreshtoken"
//}
func LoginRefresh(c echo.Context) error {

	var loginresp models.LoginResponse
	var inputJSON models.LoginRefreshRequest

	err := c.Bind(&inputJSON)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Wrong data")
	}

	loginresp, err = models.RefreshJWTToken(inputJSON.RefreshToken)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "some problems")
	}

	return echo.NewHTTPError(http.StatusOK, loginresp)
}
