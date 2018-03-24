package routes

import (
	"net/http"

	"github.com/labstack/echo"

	"gitlab.com/andreluizmachado/go-challenge-ac001/application/resource/action"
)

func PutCitiesRoutes(e *echo.Echo) {
	cityGroup := e.Group("/city")
	cityGroup.GET("/:id", action.GetCity)
	cityGroup.POST("", action.CreateCity)
	cityGroup.PUT("/:id", action.UpdateCity)
	cityGroup.DELETE(":id", hello)


	citiesGroup := e.Group("/cities")
	citiesGroup.GET("", hello)
	citiesGroup.DELETE("", hello)	
}


// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}