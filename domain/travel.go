package domain

import (
	"gitlab.com/andreluizmachado/go-challenge-ac001/representation/entity"
	"gitlab.com/andreluizmachado/go-challenge-ac001/representation/repository"

	"strconv"
)

type Travel struct {
	CityRepository *repository.CityRepository
}

var citiesVerified []int
var path []int

func (travel *Travel) GetPath(cityIdOrigin int, cityIdDestinate int, byFilter []string) *entity.Travel {
	travelEntity := travel.getPathRecursively(cityIdOrigin, cityIdDestinate, byFilter)

	citiesVerified = []int{}
	path = []int{}

	return travelEntity
}

func NewTravel(cityRepository *repository.CityRepository) *Travel {
	return &Travel{cityRepository}
}

func (travel *Travel) getPathRecursively(cityIdOrigin int, cityIdDestinate int, byFilter []string) *entity.Travel {
	var allCitiesVerified bool = true

	path = append(path, cityIdOrigin)

	citiesVerified = append(citiesVerified, cityIdOrigin)

	city := travel.CityRepository.FindById(strconv.Itoa(cityIdOrigin))

	if travel.hasDestinate(city.Borders, cityIdDestinate) {

		path = append(path, cityIdDestinate)

		travelEntity := &entity.Travel{path}

		if travel.hasTwoFilters(byFilter) {
			return travelEntity
		}

		if travel.hasMoreThanThreeStops(travelEntity) {
			return travelEntity
		}

		firstStop, _ := strconv.Atoi(byFilter[0])
		secondStop, _ := strconv.Atoi(byFilter[1])

		if travel.hasDestinate(travelEntity.Path, firstStop) && travel.hasDestinate(travelEntity.Path, secondStop) {
			return travelEntity
		}

		travel.removePathElements(2)
	}

	for _, border := range city.Borders {
		if !travel.hasDestinate(citiesVerified, border) {
			allCitiesVerified = false
			nextTravel := travel.getPathRecursively(border, cityIdDestinate, byFilter)

			if nextTravel != nil {
				return nextTravel
			}
		}
	}

	if allCitiesVerified {
		travel.removePathElements(1)
		return nil
	}

	return nil
}

func (travel *Travel) removePathElements(length int) {
	if len(path) >= length {
		path = path[:len(path)-length]
	}
}

func (travel *Travel) hasTwoFilters(byFilter []string) bool {
	return len(byFilter) != 2
}

func (travel *Travel) hasMoreThanThreeStops(travelEntity *entity.Travel) bool {
	return len(travelEntity.Path) < 4
}

func (travel *Travel) hasDestinate(borders []int, cityIdDestinate int) bool {
	for _, border := range borders {
		if border == cityIdDestinate {
			return true
		}
	}
	return false
}
