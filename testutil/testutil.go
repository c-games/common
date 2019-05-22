package testutil

import (
	"testing"
)

func TestFailIfErr(err error, t *testing.T) {
	if err != nil {
		t.Fail()}
}

func Is(result, expect interface{}, t *testing.T) {
	 if result != expect {
		t.Logf("\nresult: %s \nexpect: %s", result, expect)
		t.Fail()
	}
}
