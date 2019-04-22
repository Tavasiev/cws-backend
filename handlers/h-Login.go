package handlers

import (
	"log"
	"net/http"

	"github.com/labstack/echo"

	//local

	"github.com/Tavasiev/cws-backend/models"
)

func Login(c echo.Context) error {

	var inputJSON models.LoginRequest
	err := c.Bind(&inputJSON)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Wrong data")
	}
	//db := dbconn.GetConnect()

	response, err := models.AuthenticateUser(inputJSON)
	if err != nil {
		log.Println("Error creating Jwt token", err)
		return c.String(http.StatusInternalServerError, "something went wrong")
	}

	return c.JSON(http.StatusOK, response)
}
