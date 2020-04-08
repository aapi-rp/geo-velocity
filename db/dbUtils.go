package db

import (
	"database/sql"
	"github.com/aapi-rp/geo-velocity/base"
	_ "github.com/mattn/go-sqlite3"
)

var SqliteConn *sql.DB

func InitData() (*sql.DB, error) {

	dbpath := base.DBPath()

	SqliteConn, err := sql.Open("sqlite3", dbpath)
	if err != nil {
		return nil, err
	}

	ct, err := SqliteConn.Prepare(createGVTable)

	if err != nil {
		return nil, err
	}

	ct.Exec()

	err = SqliteConn.Ping()

	if err != nil {
		return nil, err
	}
	return SqliteConn, nil
}

func InsertDBRow(query string, values ...interface{}) error {
	insert, err := SqliteConn.Prepare(query)

	if err != nil {
		return err
	}

	_, err = insert.Exec(values)

	if err != nil {
		return err
	}

	return nil
}
