package main

import (
	"github.com/Tavasiev/cws-backend/handlers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/CreateModels", handlers.CreateModels)
	e.GET("/DropModels", handlers.DropModels)
	e.POST("/AddCity", handlers.AddCity)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
