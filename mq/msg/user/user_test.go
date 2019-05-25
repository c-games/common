package user

import (
	"gitlab.3ag.xyz/backend/common/mq/msg"
	"gitlab.3ag.xyz/backend/common/testutil"
	"testing"
)

func TestFetcData(t *testing.T) {
	data := &FetchData{
		Token: "",
	}
	expect := `{"token":""}`

	testutil.Is(msg.ToJson(data), expect, t)
}

func TestLoginData(t *testing.T) {
	data := &LoginData{
		Account: "syber",
		Password: "syberspwd",
		Token: "a-random-string",
		Ip: "127.0.0.1",
		Time: "2019-04-01",
		Platform: "PC",
	}
	expect := `{"account":"syber",` +
		`"password":"syberspwd",` +
		`"token":"a-random-string",` +
		`"ip":"127.0.0.1",` +
		`"time":"2019-04-01",` +
		`"platform":"PC"}`

	testutil.Is(msg.ToJson(data), expect, t)
}

func TestRegisterData(t *testing.T) {
	data := RegisterData{
		AgentId: "agent-id-number",
		Account: "syber",
		Name: "sybersname",
		Email: "syber@test.cg",
		Mobile: "00000000",
		Qq: "qq-number",
		Wechat: "wechat-number",
		Ip: "127.0.0.1",
		Time: "2019-04-01",
		Platform: "PC",
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

