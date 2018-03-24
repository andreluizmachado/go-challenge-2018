package action

import (
	"net/http"

	"github.com/labstack/echo"	

	"gitlab.com/andreluizmachado/go-challenge-ac001/infrastructure"

	"gitlab.com/andreluizmachado/go-challenge-ac001/representation/repository"	
)

func GetAllCities(c echo.Context) error {
	dbConnection := infrastructure.GetDbConnection()

	cityRepository := repository.NewCityRepository(dbConnection)

	cities := cityRepository.FindAll()

	if cities == nil {
		return c.JSON(http.StatusOK, cities)
	}

	return c.JSON(http.StatusOK, cities)
}