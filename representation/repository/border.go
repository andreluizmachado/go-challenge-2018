package repository

import (
	"database/sql"
	"gitlab.com/andreluizmachado/go-challenge-ac001/representation/entity"
	"log"
)

type BorderRepository struct {
	Connection *sql.DB
}


func (borderRepositoy *BorderRepository) Store(border *entity.Border) int {

	statement, err := borderRepositoy.Connection.Prepare("insert into borders(city_id, border_city) values(?, ?)")

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

func (borderRepository *BorderRepository) StoreList(citiId int, borderList []int ) {
	for _, border := range borderList {
		borderRepository.Store(&entity.Border{0, citiId, border});
	}	
}

func (borderRepositoy *BorderRepository) DeleteByCity(city *entity.City) int {

	statement, err := borderRepositoy.Connection.Prepare("DELETE FROM borders WHERE city_id = ?")

	if err != nil {
		log.Fatal(err)
	}

	result, err := statement.Exec(city.Id)

	if err != nil {
		log.Fatal(err)
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		log.Fatal(err)
	}

	return int(rowsAffected)
}

func NewBorderRepository(dbConnection *sql.DB) *BorderRepository {
	return &BorderRepository{dbConnection}
}