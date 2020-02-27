package userMsg

import (
	"gitlab.3ag.xyz/backend/common/mq/msg"
	"gitlab.3ag.xyz/backend/common/testutil"
	"testing"
)


func TestLoginData(t *testing.T) {
	data := &LoginData{
		Account: "syber",
		Password: "syberspwd",
	}
	expect := `{"account":"syber",` +
		`"password":"syberspwd"}`

	testutil.Is(msg.ToJson(data), expect, t)
}

func TestRegisterData(t *testing.T) {
	data := RegisterData{
		AgentId: 10001,
		Account: "syber",
		Name: "sybersname",
		Email: "syber@test.cg",
		Mobile: "00000000",
		Qq: "qq-number",
		Wechat: "wechat-number",
		Ip: "127.0.0.1",
		Platform: 0,
	}
	expect := `{` +
		`"agent_id":"agent-id-number",` +
		`"account":"syber",` +
		`"name":"sybersname",` +
		`"email":"syber@test.cg",` +
		`"mobile":"00000000",` +
		`"qq":"qq-number",` +
		`"wechat":"wechat-number",` +
		`"ip":"127.0.0.1",` +
		`"time":"2019-04-01",` +
		`"platform":"PC"` +
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

