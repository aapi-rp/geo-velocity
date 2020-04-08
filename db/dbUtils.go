package db

import (
	"database/sql"
	"github.com/aapi-rp/geo-velocity/base"
	"github.com/aapi-rp/geo-velocity/logger"
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

func InsertDBRow(query string, values ...interface{}) {
	insert, err := SqliteConn.Prepare(query)

	if err != nil {
		logger.Error("Prepare insert event failed: ", err)
	}

	_, err = insert.Exec(values)

	if err != nil {
		logger.Error("Insert row event failed: ", err)
	}
}
