// Package action all actions of api starts where
package action

import (
	"log"
	"net/http"

	"github.com/labstack/echo"

	"gitlab.com/andreluizmachado/go-challenge-ac001/infrastructure"

	"gitlab.com/andreluizmachado/go-challenge-ac001/representation/repository"
)

// GetAllCities get all cities action of GET /cities
func GetAllCities(c echo.Context) error {
	dbConnection := infrastructure.GetDbConnection()
	defer dbConnection.Close()

	transaction, err := dbConnection.Begin()
	if err != nil {
		log.Fatal(err)
	}

	cityRepository := repository.NewCityRepository(dbConnection, transaction)

	cities := cityRepository.FindAll()

	return c.JSON(http.StatusOK, cities)
}
