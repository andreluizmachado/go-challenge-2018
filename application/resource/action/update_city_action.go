// Package action all actions of api starts where
package action

import (
	"net/http"

	"github.com/labstack/echo"

	"log"

	"gitlab.com/andreluizmachado/go-challenge-ac001/infrastructure"

	"gitlab.com/andreluizmachado/go-challenge-ac001/representation/repository"

	"gitlab.com/andreluizmachado/go-challenge-ac001/representation/entity"

	"strconv"
)

// UpdateCity update a city action of PUT /city/:id
func UpdateCity(c echo.Context) error {

	cityId := c.Param("id")

	city := new(entity.City)

	if err := c.Bind(city); err != nil {
		return err
	}

	city.Id, _ = strconv.Atoi(cityId)

	dbConnection := infrastructure.GetDbConnection()
	defer dbConnection.Close()

	transaction, err := dbConnection.Begin()
	if err != nil {
		log.Fatal(err)
	}

	cityRepository := repository.NewCityRepository(dbConnection, transaction)

	borderRepository := repository.NewBorderRepository(dbConnection, transaction)

	result := cityRepository.Update(city)

	borderRepository.DeleteByCityId(cityId)

	borderRepository.StoreList(city.Id, city.Borders)

	transaction.Commit()
	if result == false {
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, city)
}
