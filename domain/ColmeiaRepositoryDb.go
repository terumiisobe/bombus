package domain

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/terumiisobe/bombus/errs"
)

type ColmeiaRepositoryDb struct {
	client *sql.DB
}

func (d ColmeiaRepositoryDb) FindAll(status string, species string) ([]Colmeia, error) {

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
		return nil, err
	}

	colmeias := make([]Colmeia, 0)
	for rows.Next() {
		var c Colmeia
		err := rows.Scan(&c.ID, &c.ColmeiaID, &c.QRCode, &c.Species, &c.StartingDate, &c.Status)
		if err != nil {
			log.Println("Error while scanning colmeias " + err.Error())
			return nil, err
		}
		colmeias = append(colmeias, c)
	}

	return colmeias, nil
}

func (d ColmeiaRepositoryDb) ById(id string) (*Colmeia, *errs.AppError) {

	byIdSQL := "select * from colmeias where id = ?"

	row := d.client.QueryRow(byIdSQL, id)

	var c Colmeia
	err := row.Scan(&c.ID, &c.ColmeiaID, &c.QRCode, &c.Species, &c.StartingDate, &c.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Colmeia not found")
		} else {
			log.Println("Error while scanning colmeia " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}

	return &c, nil
}

func NewColmeiaRepositoryDB() ColmeiaRepositoryDb {

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

	fmt.Println("connected to mysql")

	return ColmeiaRepositoryDb{client}
}
