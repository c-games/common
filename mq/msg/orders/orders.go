package orders

import (
	"encoding/json"
	"gitlab.3ag.xyz/backend/common/mq/msg"
)

var (
	// 下單
	Bid msg.ServiceCommand = msg.NewCommand("bid")
	// 撤單
	Withdraw msg.ServiceCommand = msg.NewCommand("withdraw")
	// 撈訂單
	Fetch msg.ServiceCommand = msg.NewCommand("fetch")
	// 撤將
	Reverse msg.ServiceCommand = msg.NewCommand("reverse")
	// TODO remove
	Update msg.ServiceCommand = msg.NewCommand("update")
	// TODO 開奬
)


type BidData struct {
	Token string   `json:"token"`
	Dollar float64 `json:"dollar"`
	GameId string `json:"game_id"`
	Periods int `json:"periods"`
	Content json.RawMessage `json:"content"`
}

type WithdrawData struct {
	Account string     `json:"account"`
	Token string       `json:"token"`
	OrderId string `json:"order_id"`
}

type FetchData struct {
	Account string `json:"account"`
	Token string `json:"token"`
	OrderId string `json:"order_id"`
}

type FetchResponseData struct {
	ErrorCode int `json:"error_code"`
	OrderId string `json:"order_id"`
	UserId string `json:"user_id"`
	GameId string `json:"game_id"`
	Content json.RawMessage `json:"content"`
	Dollar float64 `json:"dollar"`
}

// TODO
type UpdateData struct {
	Account string `json:"account"`
	Token string `json:"token"`
}

// TODO
type ReverseData struct {
	Account string `json:"account"`
	Token string `json:"token"`
}

