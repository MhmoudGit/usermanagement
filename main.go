package main

import (
	"usermanagement/services/auth"
	"usermanagement/services/home"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Use(middleware.Recover())

	// Serving static files (css, html, images)
	e.Static("/static", "static")

	home.HomeRouter(e)
	auth.AuthRouter(e)

	e.Start("localhost:8000")
}
