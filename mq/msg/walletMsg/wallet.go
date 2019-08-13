package walletMsg

type UpdateData struct {
//	MasterAgentId   int     `json:"master_agent_id"`
//	AgentId         int     `json:"agent_id"`
	UserId          int64   `json:"user_id"`
	CreditChange    float32 `json:"credit_change"`
	Reason          int     `json:"reason"`
	ReasonMessage   string  `json:"reason_message"`
	GameId          int     `json:"game_id"`
	OrderId         int64   `json:"order_id"`
	TransferId      int64   `json:"transfer_id"`
	AgentTransferId string  `json:"agent_transfer_id"`
}

type ValidateData struct {
	UserId        int64   `json:"user_id"`
	ExpectCredit  float32 `json:"expect_credit"`
}

type RegisterData struct {
	MasterAgentId int     `json:"master_agent_id"`
	AgentId       int     `json:"agent_id"`
	UserId        int64   `json:"user_id"`
	DefaultCredit float32 `json:"default_credit"`
}

type QueryData struct {
//	MasterAgentId int   `json:"master_agent_id"`
//	AgentId       int   `json:"agent_id"`
	UserId        int64 `json:"user_id"`
}

// response data

type ResponseData struct {
	Credit float32 `json:"credit"`
}
