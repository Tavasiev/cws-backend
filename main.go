package main

import (
	//"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	//local
	"github.com/Tavasiev/cws-backend/handlers"
)

func main() {

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/hello", handlers.Hello)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
