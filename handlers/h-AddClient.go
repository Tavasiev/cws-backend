package handlers

import (
	"net/http"

	"github.com/go-pg/pg"
	"github.com/labstack/echo"

	//local
	"github.com/Tavasiev/cws-backend/configs"
	"github.com/Tavasiev/cws-backend/models"
)

// AddClient Добавляет инф о клиенте в таблицу Clients
// формат входного json'а:
//{
//	"phone": 89888794747,
//	"initials": "Ivanon I. I.",
//	"city": "Vladikavkaz"
//}
func AddClient(c echo.Context) error {
	configs.InitConfigs("configs/config") // получение конфиг структуры

	var inputJSON models.Clients
	err := c.Bind(&inputJSON)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Wrong data")
	}

	db := pg.Connect(&pg.Options{
		Addr:     configs.Cfg.DataBase.Addr,
		User:     configs.Cfg.DataBase.User,
		Password: configs.Cfg.DataBase.Password,
		Database: configs.Cfg.DataBase.DB,
	})
	defer db.Close()

	err = db.Insert(&models.Clients{

		Phone:      inputJSON.Phone,
		Initials:   inputJSON.Initials,
		CitiesCity: inputJSON.CitiesCity,
	})

	if err != nil {
		return echo.NewHTTPError(http.StatusOK, err.Error())
		//panic(err)
	}

	return echo.NewHTTPError(http.StatusOK, "Client Added")
}
