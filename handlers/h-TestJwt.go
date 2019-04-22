package handlers

import (
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/labstack/echo"
)

//TestJwt godoc
func TestJwt(c echo.Context) error {
	newToken, err := uuid.NewV4()
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, newToken.String())
}
