package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

func TestJwt(c echo.Context) error {
	return c.String(http.StatusOK, "you are here")
}
