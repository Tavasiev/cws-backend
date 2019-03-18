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

//CreateModels Создает все модели в бд.
func CreateModels(c echo.Context) error {
	conf := configs.MakeConfig() // получение конфиг структуры

	db := pg.Connect(&pg.Options{
		Addr:     conf.DB.Addr,
		User:     conf.DB.User,
		Password: conf.DB.Password,
		Database: conf.DB.Database,
	})
	defer db.Close()

	for _, model := range []interface{}{&models.Cities{}, &models.Workers{}, &models.Clients{}, &models.Orders{}} {
		err := db.CreateTable(model, &orm.CreateTableOptions{
			FKConstraints: true,
		})

		if err != nil {
			panic(err)
		}
	}

	return c.String(http.StatusOK, "Created")
}
