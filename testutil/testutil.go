package testutil

import (
	"encoding/json"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"gitlab.3ag.xyz/backend/common/db"
	"gitlab.3ag.xyz/backend/common/fail"
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



func GenFakeDb(setMockFn func(sqlmock.Sqlmock)) *db.DBAdapter {
	fakeDb, mock, _ := sqlmock.New()
	fakeDbAdp := db.ConnectByDB(fakeDb)
	setMockFn(mock)
	return fakeDbAdp
}

func PackStructToByte(anyStruct interface{}) []byte {
	rlt, err := json.Marshal(anyStruct)
	fail.FailOnError(err, "marshal fail in test")
	return []byte(rlt)
}