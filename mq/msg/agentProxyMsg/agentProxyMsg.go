package agentProxyMsg

import "encoding/json"

//http
type CGHTTPRequestCommand struct {
	AgentId int64    `json:"agent_id"`
	Command int    `json:"command"`
	Data    string `json:"data"`
}

type CGHTTPResponseCommand struct {
	ErrorCode int             `json:"error_code"`
	Message   string          `json:"message"`
	Data      []json.RawMessage `json:"data"`
}

//http command
type LoginHttpRequest struct {
	AgentId  int64    `json:"agent_id"`
	Account  string `json:"account"`
	Password string `json:"password"`
	Ip       string `json:"ip"`
	Platform int    `json:"platform"`
}

type LoginHttpResponse struct {
	Token string `json:"token"`
	Url   string `json:"url"`
}

type RegisterHttpRequest struct {
	AgentId  int64    `json:"agent_id"`
	Account  string `json:"account"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Mobile   string `json:"mobile"`
	QQ       string `json:"qq"`
	Wechat   string `json:"wechat"`
	Ip       string `json:"ip"`
	Platform int    `json:"platform"`
}

type LogoutHttpRequest struct {
	Account string `json:"account"`
}

type TransferInHttpRequest struct {
	Account         string  `json:"account"`
	ChangeCredit    float32 `json:"credit_change"`
	AgentTransferId string  `json:"agent_transfer_id"`
}

type TransferInHttpResponse struct {
	TransferId      int64   `json:"transfer_id"`
	Account         string  `json:"account"`
	ChangeCredit    float32 `json:"credit_change"`
	AgentTransferId string  `json:"agent_transfer_id"`
}

type TransferOutHttpRequest struct {
	Account         string  `json:"account"`
	ChangeCredit    float32 `json:"credit_change"`
	AgentTransferId string  `json:"agent_transfer_id"`
}

type TransferOutHttpResponse struct {
	TransferId      int64   `json:"transfer_id"`
	Account         string  `json:"account"`
	ChangeCredit    float32 `json:"credit_change"`
	AgentTransferId string  `json:"agent_transfer_id"`
}

type QueryCreditHttpRequest struct {
	Account string `json:"account"`
}

type QueryCreditHttpResponse struct {
	Account string  `json:"account"`
	Credit  float32 `json:"credit"`
}

type QueryUserWalletChangeHttpRequest struct {
	Account   string `json:"account"`
	BeginDate string `json:"begin_date"`
	EndDate   string `json:"end_date"`
}

type QueryAllWalletChangeHttpRequest struct {
	BeginDate string `json:"begin_date"`
	EndDate   string `json:"end_date"`
}

type QueryWalletChangeByAgentTransferIdHttpRequest struct {
	AgentTransferId string `json:"agent_transfer_id"`
}


type QueryWalletChangeHttpResponse struct {
	Account         string  `json:"account"`
	ChangeCredit    float32 `json:"credit_change"`
	BalanceCredit   float32 `json:"balance_credit"`
	TransferId      int64   `json:"transfer_id"`
	AgentTransferId string  `json:"agent_transfer_id"`
	Date            string  `json:"date"`
}

type QueryUserOrderHttpRequest struct {
	Account   string `json:"account"`
	BeginDate string `json:"begin_date"`
	EndDate   string `json:"end_date"`
}

type QueryAllOrderHttpRequest struct {
	BeginDate string `json:"begin_date"`
	EndDate   string `json:"end_date"`
}

type QueryOrderHttpResponse struct {
	Account      string  `json:"account"`
	OrderId      int64   `json:"order_id"`
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

type QueryUserStatusHttpRequest struct {
	Account string `json:"account"`
}

type QueryUserStatusHttpResponse struct {
	Account     string `json:"account"`
	LoginStatus int    `json:"login_status"`
	Status      int    `json:"status"`
}
