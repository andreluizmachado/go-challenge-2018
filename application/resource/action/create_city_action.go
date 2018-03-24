package action

import (
	"net/http"
	"fmt"

	"github.com/labstack/echo"	

	"gitlab.com/andreluizmachado/go-challenge-ac001/representation/entity"

	"gitlab.com/andreluizmachado/go-challenge-ac001/infrastructure"

	"gitlab.com/andreluizmachado/go-challenge-ac001/representation/repository"	
)

func CreateCity(c echo.Context) error {
	city := new(entity.City)

	if err := c.Bind(city); err!=nil {
		fmt.Println("bind city problems");
		return err
	}

	dbConnection := infrastructure.GetDbConnection()

	cityRepository := repository.NewCityRepository(dbConnection)

	borderRepository := repository.NewBorderRepository(dbConnection)

	citiId := cityRepository.Store(city);

	for _, border := range city.Borders {
		borderRepository.Store(&entity.Border{0, citiId, border});
	}

	city.Id = citiId;
	c.Response().Header().Set("Location", "/city/" + fmt.Sprintf("%d", citiId) )
	return c.JSON(http.StatusCreated, city)
}