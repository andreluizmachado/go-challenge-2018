package action

import (
	"net/http"

	"github.com/labstack/echo"	

	"gitlab.com/andreluizmachado/go-challenge-ac001/infrastructure"

	"gitlab.com/andreluizmachado/go-challenge-ac001/representation/repository"	

	"gitlab.com/andreluizmachado/go-challenge-ac001/representation/entity"

	"strconv"
)

func UpdateCity(c echo.Context) error {

	cityId := c.Param("id")

	city := new(entity.City)

	if err := c.Bind(city); err!=nil {
		return err
	}

	city.Id, _ = strconv.Atoi(cityId)

	dbConnection := infrastructure.GetDbConnection()

	cityRepository := repository.NewCityRepository(dbConnection)

	borderRepository := repository.NewBorderRepository(dbConnection)

	result := cityRepository.Update(city)

	borderRepository.DeleteByCity(city)

	borderRepository.StoreList(city.Id, city.Borders)

	if result == false {
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, city)
}