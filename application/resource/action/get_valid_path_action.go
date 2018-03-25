// Package action all actions of api starts where
package action

import (
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"gitlab.com/andreluizmachado/go-challenge-ac001/domain"
	"gitlab.com/andreluizmachado/go-challenge-ac001/infrastructure"
	"gitlab.com/andreluizmachado/go-challenge-ac001/representation/repository"
)

// GetValidPath get a path base on a travel of the
// city i to city 2, action GET /city/2/travel/1[?by=2&by=2]
func GetValidPath(c echo.Context) error {

	cityIdOrigin, _ := strconv.Atoi(c.Param("origin"))

	byFilter := c.Request().URL.Query()["by"]

	cityIdDestinate, _ := strconv.Atoi(c.Param("destinate"))

	dbConnection := infrastructure.GetDbConnection()
	defer dbConnection.Close()

	transaction, err := dbConnection.Begin()
	if err != nil {
		log.Fatal(err)
	}

	cityRepository := repository.NewCityRepository(dbConnection, transaction)

	travel := domain.NewTravel(cityRepository)

	travelEntiy := travel.GetPath(cityIdOrigin, cityIdDestinate, byFilter)

	if travelEntiy == nil {
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, travelEntiy)
}
