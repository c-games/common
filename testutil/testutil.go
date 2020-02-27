package testutil

import (
	"encoding/json"
	"fmt"
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
