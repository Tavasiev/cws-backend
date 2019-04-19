package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	//local
	"github.com/Tavasiev/cws-backend/configs"
	"github.com/Tavasiev/cws-backend/dbconn"
	h "github.com/Tavasiev/cws-backend/handlers"
)

func main() {

	// получение конфиг структуры
	configs.InitConfigs("configs/config")

	// подключение к бд
	err := dbconn.Connect()
	if err != nil {
		panic(err)
	}

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes

	jwtGroup := e.Group("/api/auth")
	jwtGroup.POST("/newclient", h.AddClient)
	jwtGroup.POST("/newworker", h.AddWorker)
	//jwtGroup.POST("/login", h.Login)
	//jwtGroup.POST("/refresh", h.LoginRefresh)

	// JWT middleware
	o := e.Group("/api")
	o.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS512",
		SigningKey:    []byte("mySecret"),
	}))

	e.GET("/CreateModels", h.CreateModels)
	e.GET("/DropModels", h.DropModels)
	e.POST("/AddCity", h.AddCity)

	o.GET("/main", h.TestJwt)

	// Start server
	e.Logger.Fatal(e.Start(configs.Cfg.Server.MainPort))
}
