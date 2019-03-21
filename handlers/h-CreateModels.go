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
	configs.InitConfigs("configs/config") // получение конфиг структуры

	db := pg.Connect(&pg.Options{
		Addr:     configs.Cfg.DataBase.Addr,
		User:     configs.Cfg.DataBase.User,
		Password: configs.Cfg.DataBase.Password,
		Database: configs.Cfg.DataBase.DB,
	})
	defer db.Close()

	for _, model := range []interface{}{&models.Cities{}, &models.Workers{}, &models.Clients{}, &models.Orders{}} {
		err := db.CreateTable(model, &orm.CreateTableOptions{
			FKConstraints: true,
		})

		if err != nil {
			return echo.NewHTTPError(http.StatusOK, err.Error())
			//panic(err)
		}
	}

	return c.String(http.StatusOK, "Models Created")
}
