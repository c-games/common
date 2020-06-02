package userMsg

import (
	"github.com/c-games/common/mq/msg"
	"github.com/c-games/common/testutil"
	"testing"
)


func TestLoginData(t *testing.T) {
	data := &LoginData{
		AgentId: 1000,
		Account: "syber",
		Password: "syberspwd",
	}
	expect := `{` +
		`"agent_id":1000,` +
		`"account":"syber",` +
		`"password":"syberspwd"}`

	testutil.Is(msg.ToJson(data), expect, t)
}

func TestRegisterData(t *testing.T) {
	data := RegisterData{
		AgentId: 10001,
		Account: "syber",
		Name: "sybersname",
		Password: "pw",
		Email: "syber@test.cg",
		Mobile: "00000000",
		Qq: "qq-number",
		Wechat: "wechat-number",
		Ip: "127.0.0.1",
		Platform: 0,
	}
	expect := `{` +
		`"agent_id":10001,` +
		`"account":"syber",` +
		`"name":"sybersname",` +
		`"email":"syber@test.cg",` +
		`"password":"pw",` +
		`"mobile":"00000000",` +
		`"qq":"qq-number",` +
		`"wechat":"wechat-number",` +
		`"ip":"127.0.0.1",` +
		`"platform":0` +
		`}`

	testutil.Is(msg.ToJson(data), expect, t)
}

func TestLogotData(t *testing.T) {
	data := &ValidateData{
		Token: "",
	}
	expect := `{"token":""}`

	testutil.Is(msg.ToJson(data), expect, t)
}


func TestValidateData(t *testing.T) {
	data := &ValidateData{
		Token: "",
	}
	expect := `{"token":""}`

	testutil.Is(msg.ToJson(data), expect, t)
}

