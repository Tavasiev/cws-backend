package handlers

import (
	"net/http"

	"github.com/Tavasiev/cws-backend/models"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"github.com/labstack/echo"
)

//DropModels Удаляет все модели из бд.
func DropModels(c echo.Context) error {
	db := pg.Connect(&pg.Options{
		User:     "user",
		Password: "password",
		Database: "database",
	})
	defer db.Close()

	for _, model := range []interface{}{&models.Orders{}, &models.Workers{}, &models.Clients{}, &models.Cities{}} {
		err := db.DropTable(model, &orm.DropTableOptions{})

		if err != nil {
			panic(err)
		}
	}

	return c.String(http.StatusOK, "Deleted/Dropped")
}
