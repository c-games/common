package orderMsg

import (
	"gitlab.3ag.xyz/backend/common/mq/msg"
	"gitlab.3ag.xyz/backend/common/testutil"
	"testing"
)


func TestWithdrawData(t *testing.T) {
	data := &WithdrawData{
		Token: "random-string",
		OrderId: "a-order-encoding-id",
	}

	expect := `{` +
		`"token":"random-string",` +
		`"order_id":"a-order-encoding-id"` +
		`}`

	testutil.Is(msg.ToJson(data), expect, t)
}

func TestFetchData(t *testing.T) {
	data := &FetchData{
		Token: "random-string",
		OrderId: "a-order-encoding-id",
	}

	expect := `{` +
		`"token":"random-string",` +
		`"order_id":"a-order-encoding-id"` +
		`}`

	testutil.Is(msg.ToJson(data), expect, t)
}