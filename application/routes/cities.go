package routes

import (
	"github.com/labstack/echo"

	"gitlab.com/andreluizmachado/go-challenge-ac001/application/resource/action"
)

func PutCitiesRoutes(e *echo.Echo) {
	cityGroup := e.Group("/city")
	cityGroup.GET("/:id", action.GetCity)
	cityGroup.POST("", action.CreateCity)
	cityGroup.PUT("/:id", action.UpdateCity)
	cityGroup.DELETE("/:id", action.DeleteCity)


	citiesGroup := e.Group("/cities")
	citiesGroup.GET("", action.GetAllCities)
	citiesGroup.DELETE("", action.DeleteAllCities)	
}
