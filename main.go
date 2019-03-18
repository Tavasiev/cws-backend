package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	//local
	"github.com/Tavasiev/cws-backend/configs"
	"github.com/Tavasiev/cws-backend/handlers"
)

func main() {
	conf := configs.MakeConfig() // получение конфиг структуры

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/CreateModels", handlers.CreateModels)
	e.GET("/DropModels", handlers.DropModels)
	e.POST("/AddCity", handlers.AddCity)

	// Start server
	e.Logger.Fatal(e.Start(conf.MainPort))
}
