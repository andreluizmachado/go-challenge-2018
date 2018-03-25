package action

import (
	"log"
	"net/http"

	"github.com/labstack/echo"

	"gitlab.com/andreluizmachado/go-challenge-ac001/infrastructure"

	"gitlab.com/andreluizmachado/go-challenge-ac001/representation/repository"
)

func DeleteAllCities(c echo.Context) error {
	dbConnection := infrastructure.GetDbConnection()
	defer dbConnection.Close()

	transaction, err := dbConnection.Begin()
	if err != nil {
		log.Fatal(err)
	}

	cityRepository := repository.NewCityRepository(dbConnection, transaction)

	borderRepository := repository.NewBorderRepository(dbConnection, transaction)

	cityRepository.DeleteAll()

	borderRepository.DeleteAll()

	transaction.Commit()

	return c.NoContent(http.StatusOK)
}
