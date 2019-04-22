package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	//local
	"github.com/Tavasiev/cws-backend/configs"
	"github.com/Tavasiev/cws-backend/dbconn"
	"github.com/Tavasiev/cws-backend/handlers"
)

func main() {

	// получение конфиг структуры
	configs.InitConfigs("configs/config")

	// подключение к бд
	err := dbconn.Connect()
	if err != nil {
		panic(err)
	}
	defer dbconn.CloseDbConnection(dbconn.Conn)

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/CreateModels", handlers.CreateModels)
	e.GET("/DropModels", handlers.DropModels)
	e.POST("/AddCity", handlers.AddCity)
	e.POST("/AddWorker", handlers.AddWorker)
	e.POST("/AddClient", handlers.AddClient)
	e.POST("/CheckPhone", handlers.CheckPhone)

	// Start server
	e.Logger.Fatal(e.Start(configs.Cfg.Server.MainPort))
}
