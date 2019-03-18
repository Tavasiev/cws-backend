package handlers

import (
	"net/http"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"github.com/labstack/echo"

	//local
	"github.com/Tavasiev/cws-backend/configs"
	"github.com/Tavasiev/cws-backend/models"
)

//DropModels Удаляет все модели из бд.
func DropModels(c echo.Context) error {
	conf := configs.MakeConfig() // получение конфиг структуры

	db := pg.Connect(&pg.Options{
		Addr:     conf.DB.Addr,
		User:     conf.DB.User,
		Password: conf.DB.Password,
		Database: conf.DB.Database,
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
