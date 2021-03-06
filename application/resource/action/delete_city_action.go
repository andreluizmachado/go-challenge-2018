// Package action all actions of api starts where
package action

import (
	"log"
	"net/http"

	"github.com/labstack/echo"

	"gitlab.com/andreluizmachado/go-challenge-ac001/infrastructure"

	"gitlab.com/andreluizmachado/go-challenge-ac001/representation/repository"
)

// DeleteCity delete a city action of DELETE /city/:id
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
