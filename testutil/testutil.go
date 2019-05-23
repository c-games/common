package testutil

import (
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
