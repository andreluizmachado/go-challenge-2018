package repository

import (
	"strconv"
	"database/sql"
	"gitlab.com/andreluizmachado/go-challenge-ac001/representation/entity"
	"log"
)

type CityRepository struct {
	Connection *sql.DB
}


func (cityRepositoy *CityRepository) Store(city *entity.City) int {

	statement, err := cityRepositoy.Connection.Prepare("insert into cities(name) values(?)")

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


func (cityRepositoy *CityRepository) Update(city *entity.City) bool {

	statement, err := cityRepositoy.Connection.Prepare("UPDATE cities SET name = ? WHERE id= ?")

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


func (cityRepositoy *CityRepository) FindById(id string) *entity.City {

	statement, err := cityRepositoy.Connection.Prepare("SELECT name, border_city FROM cities INNER JOIN borders ON borders.city_id=cities.id WHERE cities.id = ?")

	if err != nil {
		log.Fatal(err)
	}

    var name string
	var borders []int
	var isResultEmpty bool = true

	rows, err := statement.Query(id)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		isResultEmpty = false
		var borderCity string

		err = rows.Scan(&name, &borderCity)
		if err != nil {
			log.Fatal(err)
		}
	
		borderCityInt, _ := strconv.Atoi(borderCity)

		borders = append(borders, borderCityInt)
	}

	if isResultEmpty {
		return nil
	}

	cityId, _ := strconv.Atoi(id)

	return &entity.City{cityId, name, borders}
}

func (cityRepository *CityRepository) Delete(id string) int {

	statement, err := cityRepository.Connection.Prepare("DELETE FROM cities WHERE id = ?")

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

func NewCityRepository(dbConnection *sql.DB) *CityRepository {
	return &CityRepository{dbConnection}
}