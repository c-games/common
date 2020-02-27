package walletMsg

import (
	"gitlab.3ag.xyz/backend/common/testutil"
	"gitlab.3ag.xyz/backend/common/mq/msg"
	"testing"
)

func TestUpdateData(t *testing.T) {
	data := &UpdateData{
		UserId: 1234,
		CreditChange: 1.1,
	}

	expect := `{` +
		`"user_id":1234,` +
		`"credit_change":1.1` +
		`}`

	testutil.Is(msg.ToJson(data), expect, t)
}


func  TestValidateData(t *testing.T) {
	data := &ValidateData{
		UserId: 1234,
		ExpectCredit: 10000.1,
	}

	expect := `{` +
		`"user_id":1234,` +
		`"expect_credit":10000.1` +
		`}`

	testutil.Is(msg.ToJson(data), expect, t)
}


func TestRegisterData(t *testing.T) {
	data := &RegisterData{
		UserId: 1234,
		DefaultCredit: 10000.0,
	}

	expect := `{` +
		`"user_id":1234,` +
		`"default_credit":10000` +
		`}`

	testutil.Is(msg.ToJson(data), expect, t)
}
