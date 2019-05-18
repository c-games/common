package user

import (
	"testing"
)

func Test_Get_1(t *testing.T) {
	loginData := LoginData {
		Account: "syber",
		Password: "imyourfather",
	}
	if loginData.Get("account") != "syber" {
		t.Error("account not equal syber")
	} else {
		t.Log("pass")
	}


	if loginData.Get("password") != "imyourfather" {
		t.Error("account not equal imyourfather")
	} else {
		t.Log("pass")
	}
}
