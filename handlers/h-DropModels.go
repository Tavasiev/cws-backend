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

//DropModels Удаляет все модели из бд.
func DropModels(c echo.Context) error {
	configs.InitConfigs("configs/config") // получение конфиг структуры
	db := dbconn.GetConnect()

	for _, model := range []interface{}{&models.Orders{},
		&models.Workers{},
		&models.Clients{},
		&models.Cities{},
		&models.Sessions{}} {
		err := db.DropTable(model, &orm.DropTableOptions{})

		if err != nil {
			return echo.NewHTTPError(http.StatusOK, err.Error())
			//panic(err)
		}
	}

	return c.String(http.StatusOK, "Models Deleted/Dropped")
}
