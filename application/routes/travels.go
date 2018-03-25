package routes

import (
	"github.com/labstack/echo"

	"gitlab.com/andreluizmachado/go-challenge-ac001/application/resource/action"
)

func PutTravelsRoutes(e *echo.Echo) {
	e.GET("/city/:origin/travel/:destinate", action.GetValidPath)
}
