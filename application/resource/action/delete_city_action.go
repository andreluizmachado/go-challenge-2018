package action

import (
	"net/http"
	"log"

	"github.com/labstack/echo"	

	"gitlab.com/andreluizmachado/go-challenge-ac001/infrastructure"

	"gitlab.com/andreluizmachado/go-challenge-ac001/representation/repository"	
)

func DeleteCity(c echo.Context) error {

	cityId := c.Param("id")

	dbConnection := infrastructure.GetDbConnection()
	defer dbConnection.Close()

	transaction, err := dbConnection.Begin()
	if err != nil {
		log.Fatal(err)
	}	

	cityRepository := repository.NewCityRepository(dbConnection, transaction)

	borderRepository := repository.NewBorderRepository(dbConnection, transaction)

	result := cityRepository.Delete(cityId)

	borderRepository.DeleteByCityId(cityId)
	borderRepository.DeleteByBorder(cityId)

	transaction.Commit()

	if result < 1 {
		return c.NoContent(http.StatusNotFound)
	}

	return c.NoContent(http.StatusOK)
}