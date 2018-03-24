package action

import (
	"net/http"

	"github.com/labstack/echo"	

	"gitlab.com/andreluizmachado/go-challenge-ac001/infrastructure"

	"gitlab.com/andreluizmachado/go-challenge-ac001/representation/repository"	
)

func GetCity(c echo.Context) error {

	cityId := c.Param("id")

	dbConnection := infrastructure.GetDbConnection()

	cityRepository := repository.NewCityRepository(dbConnection)

	city := cityRepository.FindById(cityId)

	if city == nil {
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, cityRepository.FindById(cityId))
}