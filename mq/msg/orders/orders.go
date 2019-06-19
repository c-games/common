package orders

import (
	"encoding/json"
	"gitlab.3ag.xyz/backend/common/mq/msg"
	"time"
)

var (
	QueryOrder msg.ServiceCommand = msg.NewCommand("query_order")
	QueryOrderByRound msg.ServiceCommand = msg.NewCommand("query_order_by_round")
	QueryOrderByUser msg.ServiceCommand = msg.NewCommand("query_order_by_user")
	QueryOrderByAgent msg.ServiceCommand = msg.NewCommand("query_order_by_agent")
	UpdateOrder msg.ServiceCommand = msg.NewCommand("update_order")
	PlaceOrder msg.ServiceCommand = msg.NewCommand("place_order")
	OpenResult msg.ServiceCommand = msg.NewCommand("open_result")
	WithdrawResult msg.ServiceCommand = msg.NewCommand("withdraw_result")
)

func OrdersCommand(commandString string) msg.ServiceCommand {
	switch commandString {
	case "query_order":
		return QueryOrder
	case "query_order_by_round":
		return QueryOrderByRound
	case "query_order_by_user":
		return QueryOrderByUser
	case "query_order_by_agent":
		return QueryOrderByAgent
	case "update_order":
		return UpdateOrder
	case "place_order":
		return PlaceOrder
	case "open_result":
		return OpenResult
	case "withdraw_result":
		return WithdrawResult
	default:
		return msg.NullCommand
	}
}

type QueryOrderData struct {
	OrderId int64 `order_id`
}

type QueryOrderByRoundData struct {
	GameId int `game_id`
	Round int `round`
}

type QueryOrderByUserData struct {
	UserId int64 `user_id`
}

type QueryOrderByAgentData struct {
	AgentId int `agent_id`
}

type QueryResponseData struct {
    UserId   int64  `json:"user_id"`
    AgentId  int  `json:"agent_id"`
    OrderId  int64  `json:"order_id"`
    Round    int  `json:"round"`
    GameId   int `json:"game_id"`
    Target    string  `json:"target"`
    Odds      float64 `json:"odds"`
    Refund    float64 `json:"refund"`
    OrderCredit float64 `json:"order_credit"`
    OrderDate time.Time `json:"order_date"`
    PayoutCredit float64 `json:"payout_credit"`
    IsOpen       bool `json:"is_open"`
    OpenDate  time.Time `json:"open_date"`
}

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


type UpdateData struct {
	Id  int64 `json:"order_id"`
	ChangeCredit float64 `json:"change_credit"`
	Record string `json:"record"`
	// UpdateCredit bool `json:"update_credit"`
}

type UpdateResponse struct {
	IsSuccess bool `json:"is_success"`
	Error     string `json:"error"`
}

// TODO
type ReverseData struct {
	Account string `json:"account"`
	Token string `json:"token"`
}

// error code