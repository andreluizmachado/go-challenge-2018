package domain

import (
	"gitlab.com/andreluizmachado/go-challenge-ac001/representation/repository"
	"gitlab.com/andreluizmachado/go-challenge-ac001/representation/entity"

	"strconv"
)

type Travel struct {
	CityRepository *repository.CityRepository
}

var citiesVerified []int
var path []int

func (travel *Travel) GetPath(cityIdOrigin int, cityIdDestinate int) *entity.Travel {
	travelEntity := travel.getPathRecursively(cityIdOrigin, cityIdDestinate)	
	
	citiesVerified = []int{}
	path = []int{}

	return travelEntity
}

func (travel *Travel) getPathRecursively(cityIdOrigin int, cityIdDestinate int) *entity.Travel {
	path = append(path, cityIdOrigin)

	citiesVerified = append(citiesVerified, cityIdOrigin)

	city := travel.CityRepository.FindById(strconv.Itoa(cityIdOrigin))

	if travel.hasDestinate(city.Borders, cityIdDestinate) {

		path = append(path, cityIdDestinate)

		return &entity.Travel{path}
	}

	for _, border := range city.Borders {
		if !travel.hasDestinate(citiesVerified, border) {
			nextTravel := travel.GetPath(border, cityIdDestinate)

			if nextTravel != nil {
				return nextTravel
			}
		}
	}

	path = path[:len(path)-1]
	return nil
}

func (travel *Travel) hasDestinate(borders []int, cityIdDestinate int) bool {
	for _, border := range borders {
		if border == cityIdDestinate {
			return true
		}
	}
	return false
}



func NewTravel(cityRepository *repository.CityRepository) *Travel {
	return &Travel{cityRepository}	
}