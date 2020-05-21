package orderMsg

import (
	"encoding/json"
)

type QueryOrderData struct {
	OrderId int64 `json:"order_id"`
}

type QueryOrderByRoundData struct {
	GameId int   `json:"game_id"`
	Round  int64 `json:"round"`
}

type QueryMultipleOrder struct {
	Id     int64 `json:"id"`
	IdType int   `json:"id_type"`

	GameId int   `json:"game_id"`
	Round  int64 `json:"round"`

	BeginDate string `json:"begin_date"`
	EndDate   string `json:"end_date"`
	Date      string `json:"date"`

	Limit int `json:"limit"`
	Page  int `json:"page"`
}

type QueryOrderByUserData struct {
	UserId    int64  `json:"user_id"`
	BeginDate string `json:"begin_date"`
	EndDate   string `json:"end_date"`
}

type QueryOrderByAgentData struct {
	MasterAgentId int64    `json:"master_agent_id"`
	AgentId       int64    `json:"agent_id"`
	BeginDate     string   `json:"begin_date"`
	EndDate       string   `json:"end_date"`
}

type PlaceOrderData struct {
	UserId       int64   `json:"user_id"`
	GameId       int     `json:"game_id"`
	Round        int64   `json:"round"`
	Target       string  `json:"target"`
	Value        string  `json:"value"`
	Odds         float32 `json:"odds"`
	Refund       float32 `json:"refund"`
	OrderCredit  float32 `json:"order_credit"`
	EffectCredit float32 `json:"effect_credit"`
}

type WithdrawData struct {
	Account string `json:"account"`
	Token   string `json:"token"`
	OrderId string `json:"order_id"`
}

type DrawResultData struct {
	GameId  int     `json:"game_id"`
	Round   int64   `json:"round"`
	OrderId int64   `json:"order_id"`
	WinLose int     `json:"win_lose"`
	Payout  float32 `json:"payout"`
}

type FetchData struct {
	Account string `json:"account"`
	Token   string `json:"token"`
	OrderId string `json:"order_id"`
}

type UpdateData struct {
	Id           int64   `json:"order_id"`
	ChangeCredit float32 `json:"change_credit"`
	Target       string  `json:"target"`
}

// response
type FetchResponseData struct {
	ErrorCode int             `json:"error_code"`
	OrderId   string          `json:"order_id"`
	UserId    string          `json:"user_id"`
	GameId    string          `json:"game_id"`
	Content   json.RawMessage `json:"content"`
	Dollar    float32         `json:"dollar"`
}

type UpdateResponse struct {
	IsSuccess bool   `json:"is_success"`
	Error     string `json:"error"`
}

type QueryResponseData struct {
	OrderId      int64   `json:"order_id"`
	AgentId      int64   `json:"agent_id"`
	AgentAccount string  `json:"agent_account"`
	UserId       int64   `json:"user_id"`
	Account      string  `json:"account"`
	Round        int64   `json:"round"`
	GameId       int     `json:"game_id"`
	Target       string  `json:"target"`
	Value        string  `json:"value"`
	Odds         float32 `json:"odds"`
	Refund       float32 `json:"refund"`
	OrderCredit  float32 `json:"order_credit"`
	EffectCredit float32 `json:"effect_credit"`
	OrderDate    string  `json:"order_date"`
	PayoutCredit float32 `json:"payout_credit"`
	IsOpen       int     `json:"is_open"`
	OpenDate     string  `json:"open_date"`
}

type QueryOrderByAgentAndGameData struct {
	AgentId int64   `json:"agent_id"`
	GameId  int     `json:"game_id"`
	Round   int64   `json:"round"`
}

type QueryOrderByUserAndGameData struct {
	UserId int   `json:"user_id"`
	GameId int   `json:"game_id"`
	Round  int64 `json:"round"`
}

type PlaceOrderResponse struct {
	Id int64 `json:"order_id"`
}

type DrawResultResponse struct {
	OrderId      int64 `json:"order_id"`
	WalletUpdate bool  `json:"wallet_update"`
}

type WithdrawResultResponse struct {
	Success bool    `json:"success"`
	GameId  int     `json:"game_id"`
	Round   int64   `json:"round"`
	Orders  []int64 `json:"orders"`
}

// TODO rename response
type WithdrawResultData struct {
	GameId int   `json:"game_id"`
	Round  int64 `json:"round"`
}

// TODO
type ReverseData struct {
	Account string `json:"account"`
	Token   string `json:"token"`
}

func ConvertIsOpen(isDraw bool) int {
	if isDraw {
		return 1
	} else {
		return 0
	}
}
