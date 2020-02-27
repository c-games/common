package orderMsg

import (
	"gitlab.3ag.xyz/backend/common/mq/msg"
	"gitlab.3ag.xyz/backend/common/testutil"
	"testing"
)


func TestWithdrawData(t *testing.T) {
	data := &WithdrawData{
		Account: "account-name",
		Token: "random-string",
		OrderId: "a-order-encoding-id",
	}

	expect := `{` +
		`"account":"account-name",` +
		`"token":"random-string",` +
		`"order_id":"a-order-encoding-id"` +
		`}`

	testutil.Is(msg.ToJson(data), expect, t)
}

func TestFetchData(t *testing.T) {
	data := &FetchData{
		Account: "account-name",
		Token: "random-string",
		OrderId: "a-order-encoding-id",
	}

	expect := `{` +
		`"account":"account-name",` +
		`"token":"random-string",` +
		`"order_id":"a-order-encoding-id"` +
		`}`

	testutil.Is(msg.ToJson(data), expect, t)
}