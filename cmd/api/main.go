package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/ravilock/naive-banking-system/api"
)

func main() {

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	apiGroup := e.Group("/api")
	apiGroup.GET("/healthcheck", api.Healthcheck)

	// Start server
	e.Logger.Fatal(e.Start(":9191"))
}
