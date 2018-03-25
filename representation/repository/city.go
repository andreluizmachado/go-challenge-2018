// Package repository access data layer, transforms data into entities and entites into data
package repository

import (
	"database/sql"
	"gitlab.com/andreluizmachado/go-challenge-ac001/representation/entity"
	"log"
	"strconv"
)

type CityRepository struct {
	Connection  *sql.DB
	Transaction *sql.Tx
}

// Store create a city
func (cityRepositoy *CityRepository) Store(city *entity.City) int {

	statement, err := cityRepositoy.Transaction.Prepare("insert into cities(name) values(?)")

	if err != nil {
		log.Fatal(err)
	}

	result, err := statement.Exec(city.Name)

	if err != nil {
		log.Fatal(err)
	}

	id, err := result.LastInsertId()

	if err != nil {
		log.Fatal(err)
	}

	return int(id)
}

// Update update a city by id
func (cityRepositoy *CityRepository) Update(city *entity.City) bool {

	statement, err := cityRepositoy.Transaction.Prepare("UPDATE cities SET name = ? WHERE id= ?")

	if err != nil {
		log.Fatal(err)
	}

	result, err := statement.Exec(city.Name, city.Id)

	if err != nil {
		log.Fatal(err)
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		log.Fatal(err)
	}

	return rowsAffected == 1
}

// FindById return a city a your borders by id table field
func (cityRepositoy *CityRepository) FindById(id string) *entity.City {

	statement, err := cityRepositoy.Connection.Prepare("SELECT name, border_city FROM cities LEFT JOIN borders ON borders.city_id=cities.id WHERE cities.id = ?")

	if err != nil {
		log.Fatal(err)
	}

	var name string
	var borders []int = []int{}
	var isResultEmpty bool = true

	rows, err := statement.Query(id)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		isResultEmpty = false
		var borderCity int

		rows.Scan(&name, &borderCity)

		if borderCity > 0 {
			borders = append(borders, borderCity)
		}
	}

	if isResultEmpty {
		return nil
	}

	cityId, _ := strconv.Atoi(id)

	return &entity.City{cityId, name, borders}
}

// FindAll returns a list of cities with your borders
func (cityRepositoy *CityRepository) FindAll() *entity.Cities {

	statement, err := cityRepositoy.Connection.Prepare("SELECT cities.id, name, border_city FROM cities LEFT JOIN borders ON borders.city_id=cities.id ORDER BY cities.id ASC")

	if err != nil {
		log.Fatal(err)
	}

	var cities entity.Cities = entity.Cities{[]entity.City{}}
	var city entity.City = entity.City{}
	city.Borders = []int{}
	var currentCityId int = 0
	var cityPosition int = -1

	rows, err := statement.Query()

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		var borderCity int
		var name string
		var id int

		rows.Scan(&id, &name, &borderCity)

		if currentCityId != id {
			city.Id = id
			city.Name = name
			cities.Cities = append(cities.Cities, city)
			currentCityId = id
			cityPosition++
		}

		if borderCity > 0 {
			var borders []int = cities.Cities[cityPosition].Borders

			cities.Cities[cityPosition].Borders = append(borders, borderCity)
		}
	}

	return &cities
}

// Delete delete a City by id table field
func (cityRepository *CityRepository) Delete(id string) int {

	statement, err := cityRepository.Transaction.Prepare("DELETE FROM cities WHERE id = ?")

	if err != nil {
		log.Fatal(err)
	}

	result, err := statement.Exec(id)

	if err != nil {
		log.Fatal(err)
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		log.Fatal(err)
	}

	return int(rowsAffected)
}

// DeleteAll delete all cities
func (cityRepository *CityRepository) DeleteAll() int {

	statement, err := cityRepository.Transaction.Prepare("DELETE FROM cities")

	if err != nil {
		log.Fatal(err)
	}

	result, err := statement.Exec()

	if err != nil {
		log.Fatal(err)
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		log.Fatal(err)
	}

	return int(rowsAffected)
}

// NewCityRepository returns a instance of city repository
func NewCityRepository(dbConnection *sql.DB, transaction *sql.Tx) *CityRepository {
	return &CityRepository{dbConnection, transaction}
}
