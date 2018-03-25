package action

import (
	"net/http"
	"strconv"
	"log"

	"github.com/labstack/echo"
	"gitlab.com/andreluizmachado/go-challenge-ac001/domain"
	"gitlab.com/andreluizmachado/go-challenge-ac001/infrastructure"
	"gitlab.com/andreluizmachado/go-challenge-ac001/representation/repository"
)

func GetValidPath(c echo.Context) error {

	cityIdOrigin, _ := strconv.Atoi(c.Param("origin"))

	cityIdDestinate, _ := strconv.Atoi(c.Param("destinate"))

	dbConnection := infrastructure.GetDbConnection()
	defer dbConnection.Close()

	transaction, err := dbConnection.Begin()
	if err != nil {
		log.Fatal(err)
	}	

	cityRepository := repository.NewCityRepository(dbConnection, transaction)

	travel := domain.NewTravel(cityRepository)

	travelEntiy := travel.GetPath(cityIdOrigin, cityIdDestinate)

	if (travelEntiy.Path == nil) {
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, travelEntiy)
}