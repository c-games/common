package orders

import "encoding/json"

type OrdersCommand string

const (
	Bid OrdersCommand = "bid" // 下單
	Withdraw OrdersCommand = "withdraw" // 撤單
	Fetch OrdersCommand = "fetch" // 撈訂單
	Reverse OrdersCommand = "reverse" // 撤將
	Update OrdersCommand = "update" // TODO
	// TODO 開奬
)

// TODO 下單 api 要改設計
type BidData struct {
	Account string `json:"account"`
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


type UpdateData struct {
	Account string `json:"account"`
	Token string `json:"token"`
}



type ReverseData struct {
	Account string `json:"account"`
	Token string `json:"token"`
}

