package database

import (
	"database/sql"

	"github.com/gustavohmsilva/Praticando-Go/Aula_XI/dataparser"
	// Driver for SQLite3 Database
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func init() {
	database, err := sql.Open("sqlite3", "miniWebStore.db")
	treatError(err)
	db = database
}

func treatError(err error) {
	if err != nil {
		panic(err)
	}
}

// InsertProduct insert a product in the SQLite3 database configured in init
func InsertProduct(p dataparser.Product) (int64, error) {
	stmt, err := db.Prepare(
		`INSERT INTO products
			(name, category, "desc")
		VALUES
			(?, ?, ?);
		`,
	)
	if err != nil {
		return 0, err
	}
	resp, err := stmt.Exec(p.Name, p.Category, p.Desc)
	if err != nil {
		return 0, err
	}
	id, err := resp.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}
