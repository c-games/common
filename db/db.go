package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"gitlab.3ag.xyz/core/backend/common/fail"
	// "fmt"
)

type DB struct {
	Conn *sql.DB
}

type ConnPool struct {
	// TODO
}

var dbConfig string

func Init(config string) {
	dbConfig = config
}

//TODO golang 沒有 overload, 所以先這樣寫
func AnotherConnect(dbUrl string) *DB {
	db, err := sql.Open("mysql", dbUrl)
	fail.FailOnError(err, "Failed to connect db")

	err = db.Ping() // This DOES open a connection if necessary. This makes sure the database is accessible
	// fail.FailOnError(err, fmt.Sprintf("Error on opening database connection: %s", err.Error()))
	fail.FailOnError(err, "Failed to ping db")

	rlt := DB { Conn: db }
	return &rlt
}

func Connect() *DB {
	return AnotherConnect(dbConfig)
}

func (d *DB) Exec(query string, args ...interface{}) sql.Result {
	result, err := d.Conn.Exec(query, args...)
	fail.FailOnError(err, "Failed to query")
	return result
}

// TODO query 時要改一下
func (d *DB) PrepareQuery(query string, args ...interface{}) *sql.Rows {
	stmt, err := d.Conn.Prepare(query)
	rows, err := stmt.Query(args...)
	fail.FailOnError(err, "Failed to query")
	return rows
}

func (d *DB)QueryRow(query string, args ...interface{}) *sql.Row {
	row := d.Conn.QueryRow(query, args...)
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
