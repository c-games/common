package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gitlab.3ag.xyz/backend/common/coll"
	"gitlab.3ag.xyz/backend/common/fail"
	"gitlab.3ag.xyz/backend/common/str"
	"reflect"
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

func (d *DBAdapter)Query(query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := d.Connect.Query(query, args...)
	return rows, err
}

func (d *DBAdapter)QueryRow(query string, args ...interface{}) *sql.Row {
	row := d.Connect.QueryRow(query, args...)
	return row
}

func (adp *DBAdapter) PrepareExec(query string, args...interface{}) sql.Result {
	stmt, err := adp.Connect.Prepare(query)
	fail.FailOnError(err, "Failed to prepare")

	rlt, err := stmt.Exec(args...)

	fail.FailOnError(err, "Failed to exec")
	return rlt
}

func (adp *DBAdapter) Prepare(query string) *sql.Stmt {
	stmt, err := adp.Connect.Prepare(query)
	fail.FailOnError(err, "Failed to prepare")
	return stmt
}

// support functions
// TODO Deprecated
func IfNoRowOr(err error, fn func(), noRowFn func()) {
	if err == sql.ErrNoRows {
		noRowFn()
	} else {
		fn()
	}
}

func GenDropTable(s interface{}) string {
	rfs := reflect.TypeOf(s)
	if rfs.Name() == "" {
		return "DROP TABLE "
	} else {
		return "DROP TABLE `" + str.Pascal2Snake(rfs.Name()) + "`;"
	}
}

func GenCreateTable(s interface{}) string {
	rfs := reflect.TypeOf(s)
	var sqlString string
	if rfs.Name() == "" {
		sqlString = "CREATE TABLE "
	} else {

		sqlString = "CREATE TABLE `" + str.Pascal2Snake(rfs.Name()) + "` "
	}


	fields := ""
	var pk []string
	for idx := 0 ; idx < rfs.NumField() ; idx++ {
		f := rfs.Field(idx)
		name := "`" + str.Pascal2Snake(f.Name) + "`"

		fields = fields + name + " " + f.Tag.Get("sql") + ",\n"

		_, ok := f.Tag.Lookup("pk")
		if ok {
			pk = append(pk, name)
		}

	}

	if len(fields) > 0 && len(pk) == 0{
		fields = fields[:len(fields) - 2]
	}

	pkStr := ""
	if len(pk) != 0 {
		pkStr = "PRIMARY KEY " + "(" + coll.JoinString(pk, ",") + ")"
	}


	sqlString = sqlString + "(\n" + fields + pkStr + ") ENGINE=INNODB CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;"

	fmt.Println(sqlString)
	return sqlString
}

func GenCreateIndex(s interface{}) string {

	rfs := reflect.TypeOf(s)
	sqlString := ""
	if rfs.Name() == "" {
		return ""
	}

	tableName := str.Pascal2Snake(rfs.Name())

	index := make(map[string][]string)
	var keys []string
	//var idx map[string][]string
	for idx := 0 ; idx < rfs.NumField() ; idx++ {
		f := rfs.Field(idx)
		indexName, ok := f.Tag.Lookup("index")
		if ok {
			_, ok := index[indexName]
			fieldName := "`" + str.Pascal2Snake(f.Name) + "`"
			if !ok {
				keys = append(keys, indexName)
				index[indexName] = []string{fieldName}
			} else {
				index[indexName] = append(index[indexName], fieldName)
			}

		}

	}

	// NOTE golang 的 map 不保證順序，所以要自己處理 keys
	for _, indexName := range keys {
		indexSet := index[indexName]
		idxs := coll.JoinString(indexSet, ",")
		sqlString = sqlString + "CREATE INDEX " + indexName + " ON `" + tableName + "` (" + idxs + ");"
	}

	return sqlString
}


func CompareParams(params []interface{}, expectParamCount int, expectParamTypes []reflect.Kind) error {
	totalParams := len(params)
	if totalParams != expectParamCount {
		panic("stmt 所需參數不同")
	}
	// check type
	for idx, param := range params {
		expectType := expectParamTypes[idx]
		paramType := reflect.TypeOf(param)

		if expectType != paramType.Kind() {
			panic("query stmt 資料型態有錯")
		}
	}
	return nil
}

// NOTE 為了要共用 sql.rows 和 sql.row 的 Scan
// ---------------------------------------------------

type SqlRowLike interface {
	Scan(...interface{}) error
}

type Scannable interface {
	Scan(SqlRowLike) error
}

func QueryCondition(queryResult Scannable, rowLike SqlRowLike) error {
	err := queryResult.Scan(rowLike)
	if err == nil {
		return nil
	} else if err == sql.ErrNoRows  {
		return nil
	} else {
		// rows 的 Scan 有可能會因為沒有先 call Next() 而出錯，那需要自己處理，所以直接 return err
		return err
	}
}

func QueryCondition2(queryResultType reflect.Type, rowLike SqlRowLike) (interface{}, error) {
	newInstent := reflect.New(queryResultType)
	elementInstent := newInstent.Elem().Interface()
	// TODO 要 check elementInstent 有 Scannable
	err := QueryCondition(elementInstent.(Scannable), rowLike)
	return elementInstent, err
}

