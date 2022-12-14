package main

import (
	"alterra-agmc/config"
	"alterra-agmc/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	config.StartDB()
	e := echo.New()

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}` + "\n",
	}))
	e.Use(middleware.Recover())

	// Routes
	routes.User(e.Group(""))
	routes.Book(e.Group(""))

	e.Logger.Fatal(e.Start(":8080"))
}
