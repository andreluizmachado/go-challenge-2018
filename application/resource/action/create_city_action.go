// Package action all actions of api starts where
package action

import (
	"log"
	"net/http"

	"strconv"

	"github.com/labstack/echo"

	"gitlab.com/andreluizmachado/go-challenge-ac001/representation/entity"

	"gitlab.com/andreluizmachado/go-challenge-ac001/infrastructure"

	"gitlab.com/andreluizmachado/go-challenge-ac001/representation/repository"
)

// CreateCity create a city action of POST /city
func CreateCity(c echo.Context) error {
	city := new(entity.City)

	if err := c.Bind(city); err != nil {
		log.Println("bind city problems")
		return err
	}

	dbConnection := infrastructure.GetDbConnection()
	defer dbConnection.Close()

	transaction, err := dbConnection.Begin()
	if err != nil {
		log.Fatal(err)
	}

	cityRepository := repository.NewCityRepository(dbConnection, transaction)

	borderRepository := repository.NewBorderRepository(dbConnection, transaction)

	citiId := cityRepository.Store(city)
	city.Id = citiId

	borderRepository.StoreList(city.Id, city.Borders)

	transaction.Commit()

	c.Response().Header().Set("Location", "/city/"+strconv.Itoa(citiId))
	return c.JSON(http.StatusCreated, city)
}
