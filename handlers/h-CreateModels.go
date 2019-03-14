package handlers

import (
	"net/http"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"github.com/labstack/echo"

	//local
	"github.com/Tavasiev/cws-backend/models"
)

//CreateModels Создает модели в бд
func CreateModels(c echo.Context) error {
	//Ссылка на документацию https://godoc.org/github.com/go-pg/pg#Connect
	db := pg.Connect(&pg.Options{
		User:     "user",
		Password: "passwrod",
		Database: "database",
	})
	defer db.Close()

	for _, model := range []interface{}{&models.Cities{}, &models.Workers{}, &models.Clients{}, &models.Orders{}} {
		err := db.CreateTable(model, &orm.CreateTableOptions{
			//Temp:          true, // create temp table
			FKConstraints: true,
		})

		if err != nil {
			panic(err)
		}
	}

	return c.String(http.StatusOK, "Созданно.")
}
