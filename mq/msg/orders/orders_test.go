package orders

import (
	"encoding/json"
	"gitlab.3ag.xyz/backend/common/mq/msg"
	"gitlab.3ag.xyz/backend/common/testutil"
	"testing"
)

func TestBidData(t *testing.T) {
	data := &BidData{
		Token: "random-string",
		Dollar: 1.1,
		GameId: "game-id-string",
		Periods: 1234,
		Content: json.RawMessage(`{"k":"v"}`),
	}

	expect := `{` +
		`"token":"random-string",` +
		`"dollar":1.1,` +
		`"game_id":"game-id-string",` +
		`"periods":1234,` +
		`"content":{"k":"v"}` +
		`}`

	testutil.Is(msg.ToJson(data), expect, t)
}

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