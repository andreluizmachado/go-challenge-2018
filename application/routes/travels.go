// Package routes this package configurate the routes of the api
package routes

import (
	"github.com/labstack/echo"

	"gitlab.com/andreluizmachado/go-challenge-ac001/application/resource/action"
)

// PutTravelsRoutes config travels routes
func PutTravelsRoutes(e *echo.Echo) {
	e.GET("/city/:origin/travel/:destinate", action.GetValidPath)
}
