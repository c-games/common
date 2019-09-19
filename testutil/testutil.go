package testutil

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"gitlab.3ag.xyz/backend/common/db"

	"testing"
)


func TestFailIfErr(t *testing.T, err error, message string) {
	if err != nil {
		t.Logf("%s", message)
		t.Log(err)
		t.Fail()
	}
}

func TestFailIfErrf(t *testing.T, err error, format string, args ...interface{}) {
	TestFailIfErr(t, err, fmt.Sprintf(format, args))
}

func Is(result, expect interface{}, t *testing.T) {
	 if result != expect {
		t.Logf("\nresult: %s \nexpect: %s", result, expect)
		t.Fail()
	}
}

func _wrapFakeDb(fakeDb *sql.DB, mock sqlmock.Sqlmock, setMockFn func(sqlmock.Sqlmock)) *db.DBAdapter {
	fakeDbAdp := db.ConnectByDB(fakeDb)
	setMockFn(mock)
	return fakeDbAdp
}

func GenFakeDb(setMockFn func(sqlmock.Sqlmock)) *db.DBAdapter { //
	fakeDb, mock, _ := sqlmock.New()
	return _wrapFakeDb(fakeDb, mock, setMockFn)
}

func GenFakeDb_QueryMatcherEqual(setMockFn func(sqlmock.Sqlmock)) *db.DBAdapter {
	fakeDb, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	return _wrapFakeDb(fakeDb, mock, setMockFn)
}

func PackStructToByte(anyStruct interface{}) []byte {
	rlt, err := json.Marshal(anyStruct)
	if err != nil {
		panic("marshal fail in test, err = " + err.Error())
	}
	return []byte(rlt)
}

// ref: https://stackoverflow.com/a/31596110
func AssertPanic(t *testing.T, f func()) {
    defer func() {
        if r := recover(); r == nil {
            t.Errorf("The code did not panic")
        }
    }()
    f()
}
