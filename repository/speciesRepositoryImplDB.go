package repository

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"bombus/domain"
	"bombus/errs"

	_ "github.com/go-sql-driver/mysql"
)

type SpeciesRepositoryImplDB struct {
	client *sql.DB
}

// TODO: refactor errors
// TODO: refactor to a centralized connection with DB
func NewSpeciesRepositoryDB() SpeciesRepositoryImplDB {
	client, err := sql.Open("mysql", "bombus_usr:bombuspass@tcp(localhost:3306)/bombus?parseTime=true")
	if err != nil {
		panic(err)
	}

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	err = client.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to MySQL")

	return SpeciesRepositoryImplDB{client}
}

func (d SpeciesRepositoryImplDB) FindAll() ([]domain.Species, *errs.AppError) {
	var rows *sql.Rows
	var err error

	findAllSql := "select * from species"
	rows, err = d.client.Query(findAllSql)
	if err != nil {
		e := errs.NewDatabaseError(err.Error())
		log.Println("[SpeciesRepositoryDB | FindAll]", e.Message)
		return nil, e
	}

	species := make([]domain.Species, 0)
	for rows.Next() {
		var id int
		var scientificName, commonName string
		err := rows.Scan(&id, &scientificName, &commonName)
		if err != nil {
			e := errs.NewDatabaseError(err.Error())
			log.Println("[SpeciesRepositoryDB | FindAll]", e.Message)
			return nil, e
		}
		s := domain.NewSpecies(id, scientificName, commonName)
		species = append(species, s)
	}

	return species, nil
}

func (d SpeciesRepositoryImplDB) ById(id string) (*domain.Species, *errs.AppError) {
	byIdSQL := "select * from species where id = ?"

	row := d.client.QueryRow(byIdSQL, id)

	var idFound int
	var scientificName, commonName string
	err := row.Scan(&idFound, &scientificName, &commonName)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("species")
		} else {
			e := errs.NewDatabaseError(err.Error())
			log.Println("[SpeciesRepositoryDB | ById]", e.Message)
			return nil, e
		}
	}

	s := domain.NewSpecies(idFound, scientificName, commonName)
	return &s, nil
}
