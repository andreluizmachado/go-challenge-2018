package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"gitlab.com/andreluizmachado/go-challenge-ac001/application/routes"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	routes.PutCitiesRoutes(e)

	// Start server
	e.Logger.Fatal(e.Start(":3000"))
}	