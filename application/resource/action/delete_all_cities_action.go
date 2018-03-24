package action

import (
	"net/http"

	"github.com/labstack/echo"	

	"gitlab.com/andreluizmachado/go-challenge-ac001/infrastructure"

	"gitlab.com/andreluizmachado/go-challenge-ac001/representation/repository"	
)

func DeleteAllCities(c echo.Context) error {
	dbConnection := infrastructure.GetDbConnection()

	cityRepository := repository.NewCityRepository(dbConnection)

	borderRepository := repository.NewBorderRepository(dbConnection)

	cityRepository.DeleteAll()

	borderRepository.DeleteAll()

	return c.NoContent(http.StatusOK)
}