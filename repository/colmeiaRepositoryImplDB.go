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

type ColmeiaRepositoryImplDB struct {
	client *sql.DB
}

func NewColmeiaRepositoryDB() ColmeiaRepositoryImplDB {
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

	return ColmeiaRepositoryImplDB{client}
}

func (d ColmeiaRepositoryImplDB) FindAll(status string, species string) ([]domain.Colmeia, *errs.AppError) {
	var rows *sql.Rows
	var err error

	findAllSql := "select * from colmeias"
	if status != "" {
		findAllSql += " where status_id = ?"
		rows, err = d.client.Query(findAllSql, status)
	} else if species != "" {
		findAllSql += " where species_id = ?"
		rows, err = d.client.Query(findAllSql, species)
	} else {
		rows, err = d.client.Query(findAllSql)
	}
	if err != nil {
		log.Println("Error while getting colmeias table: " + err.Error())
		return nil, errs.NewDatabaseError(err.Error())
	}

	colmeias := make([]domain.Colmeia, 0)
	for rows.Next() {
		var c domain.Colmeia
		err := rows.Scan(&c.ID, &c.ColmeiaID, &c.QRCode, &c.Species, &c.StartingDate, &c.Status)
		if err != nil {
			log.Println("Error while scanning colmeias " + err.Error())
			return nil, errs.NewDatabaseError(err.Error())
		}
		colmeias = append(colmeias, c)
	}

	return colmeias, nil
}

func (d ColmeiaRepositoryImplDB) ById(id string) (*domain.Colmeia, *errs.AppError) {
	byIdSQL := "select * from colmeias where id = ?"

	row := d.client.QueryRow(byIdSQL, id)

	var c domain.Colmeia
	err := row.Scan(&c.ID, &c.ColmeiaID, &c.QRCode, &c.Species, &c.StartingDate, &c.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return &c, errs.NewNotFoundError("Colmeia not found")
		} else {
			log.Println("Error while scanning colmeia " + err.Error())
			return &c, errs.NewDatabaseError("Unexpected database error")
		}
	}

	return &c, nil
}

func (d ColmeiaRepositoryImplDB) Create(colmeia domain.Colmeia) *errs.AppError {
	createSQL := "INSERT INTO colmeias (colmeia_id, qr_code, species_id, starting_date, status_id, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?)"

	_, err := d.client.Exec(createSQL,
		colmeia.ColmeiaID,
		colmeia.QRCode,
		colmeia.Species,
		colmeia.StartingDate,
		colmeia.Status)
	if err != nil {
		log.Println("Error while creating colmeia: " + err.Error())
		return errs.NewDatabaseError(err.Error())
	}

	return nil
}
