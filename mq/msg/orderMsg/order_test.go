package orderMsg

import (
	"github.com/c-games/common/mq/msg"
	"github.com/c-games/common/testutil"
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