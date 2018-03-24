package action

import (
	"net/http"

	"github.com/labstack/echo"	

	"gitlab.com/andreluizmachado/go-challenge-ac001/infrastructure"

	"gitlab.com/andreluizmachado/go-challenge-ac001/representation/repository"	
)

func DeleteCity(c echo.Context) error {

	cityId := c.Param("id")

	dbConnection := infrastructure.GetDbConnection()

	cityRepository := repository.NewCityRepository(dbConnection)

	borderRepository := repository.NewBorderRepository(dbConnection)

	result := cityRepository.Delete(cityId)

	borderRepository.DeleteByCityId(cityId)

	if result < 1 {
		return c.NoContent(http.StatusNotFound)
	}

	return c.NoContent(http.StatusOK)
}