package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"gitlab.3ag.xyz/core/backend/common/fail"
	// "fmt"
)

type DBAdapter struct {
	Connect *sql.DB
}

var dbConfig string

func Init(config string) {
	dbConfig = config
}

func ConnectBy(dbUrl string) *DBAdapter {
	db, err := sql.Open("mysql", dbUrl)
	fail.FailOnError(err, "Failed to connect db")

	err = db.Ping() // This DOES open a connection if necessary. This makes sure the database is accessible
	fail.FailOnError(err, "Error on opening database connection:")

	rlt := DBAdapter{ Connect: db }
	return &rlt
}

func ConnectByDB(db *sql.DB) *DBAdapter {
	return &DBAdapter{ Connect: db}
}

func Connect() *DBAdapter {
	return ConnectBy(dbConfig)
}

func (d *DBAdapter) Exec(query string, args ...interface{}) sql.Result {
	var result sql.Result
	var err error

	if len(args) == 0 {
		result, err = d.Connect.Exec(query)
	} else {
		result, err = d.Connect.Exec(query, args...)
	}

	fail.FailOnError(err, "Failed to query")
	return result
}

// TODO query 時要改一下
func (d *DBAdapter) PrepareQuery(query string, args ...interface{}) *sql.Rows {
	stmt, err := d.Connect.Prepare(query)
	fail.FailOnError(err, "Failed to prepare")
	rows, err := stmt.Query(args...)
	fail.FailOnError(err, "Failed to query")
	return rows
}

func (d *DBAdapter) Close() error {
	return d.Connect.Close()
}

func (d *DBAdapter)QueryRow(query string, args ...interface{}) *sql.Row {
	row := d.Connect.QueryRow(query, args...)
	return row
}


// support functions
//
func IfNoRowOr(err error, fn func(), noRowFn func()) {

	if err == sql.ErrNoRows {
		fn()
	} else {
		noRowFn()
	}
}
