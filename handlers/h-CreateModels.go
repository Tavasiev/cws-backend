package handlers

import (
	"net/http"

	"github.com/go-pg/pg/orm"
	"github.com/labstack/echo"

	//local
	"github.com/Tavasiev/cws-backend/configs"
	"github.com/Tavasiev/cws-backend/dbconn"
	"github.com/Tavasiev/cws-backend/models"
)

//CreateModels Создает все модели в бд.
func CreateModels(c echo.Context) error {
	configs.InitConfigs("configs/config") // получение конфиг структуры
	db := dbconn.GetConnect()

	for _, model := range []interface{}{&models.Cities{},
		&models.Workers{},
		&models.Clients{},
		&models.Orders{},
		&models.Sessions{}} {
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
