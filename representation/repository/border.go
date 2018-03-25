// Package repository access data layer, transforms data into entities
package repository

import (
	"database/sql"
	"gitlab.com/andreluizmachado/go-challenge-ac001/representation/entity"
	"log"
)

type BorderRepository struct {
	Connection  *sql.DB
	Transaction *sql.Tx
}

// Store create a border
func (borderRepository *BorderRepository) Store(border *entity.Border) int {

	statement, err := borderRepository.Transaction.Prepare("insert into borders(city_id, border_city) values(?, ?)")

	if err != nil {
		log.Fatal(err)
	}

	result, err := statement.Exec(border.CityId, border.Border)

	if err != nil {
		log.Fatal(err)
	}

	id, err := result.LastInsertId()

	if err != nil {
		log.Fatal(err)
	}

	return int(id)
}

// StoreList create a list of borders
func (borderRepository *BorderRepository) StoreList(citiId int, borderList []int) {
	for _, border := range borderList {
		borderRepository.Store(&entity.Border{0, citiId, border})
	}
}

// DeleteByCityId Delete borders by city_id table field
func (borderRepository *BorderRepository) DeleteByCityId(cityId string) int {

	statement, err := borderRepository.Transaction.Prepare("DELETE FROM borders WHERE city_id = ?")

	if err != nil {
		log.Fatal(err)
	}

	result, err := statement.Exec(cityId)

	if err != nil {
		log.Fatal(err)
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		log.Fatal(err)
	}

	return int(rowsAffected)
}

// DeleteByBorder Delete borders by border_city table field
func (borderRepository *BorderRepository) DeleteByBorder(borderCity string) int {

	statement, err := borderRepository.Transaction.Prepare("DELETE FROM borders WHERE border_city = ?")

	if err != nil {
		log.Fatal(err)
	}

	result, err := statement.Exec(borderCity)

	if err != nil {
		log.Fatal(err)
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		log.Fatal(err)
	}

	return int(rowsAffected)
}

// DeleteAll delete all Borders
func (borderRepository *BorderRepository) DeleteAll() int {

	statement, err := borderRepository.Transaction.Prepare("DELETE FROM borders")

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

// NewBorderRepository returns a instance of border repository
func NewBorderRepository(dbConnection *sql.DB, transaction *sql.Tx) *BorderRepository {
	return &BorderRepository{dbConnection, transaction}
}
