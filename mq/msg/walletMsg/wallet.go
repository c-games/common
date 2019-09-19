package walletMsg

type OrderChangeRecord struct {
	OrderId int64
	Change  float32
	Reason  int
}

type OrderRecord struct {
	CreditChange    float32 `json:"credit_change"`
	Reason          int     `json:"reason"`
	ReasonMessage   string  `json:"reason_message"`
	GameId          int     `json:"game_id"`
	OrderId         int64   `json:"order_id"`
	TransferId      int64   `json:"transfer_id"`
	AgentTransferId string  `json:"agent_transfer_id"`
}

type UpdateData struct {
	UserId          int64               `json:"user_id"`
	CreditChange    float32             `json:"credit_change"`
	Reason          int                 `json:"reason"`
	ReasonMessage   string              `json:"reason_message"` // 暫時留下來
	Detail          []OrderChangeRecord `json:"detail"`
	GameId          int                 `json:"game_id"`
	OrderId         int64               `json:"order_id"`
	TransferId      int64               `json:"transfer_id"`
	AgentTransferId string              `json:"agent_transfer_id"`
}

type MultiUpdateData struct {
	UserId      int64         `json:"user_id"`
	TotalCredit float32       `json:"total_credit"`
	Detail      []OrderRecord `json:"detail"`
}

type ValidateData struct {
	UserId       int64   `json:"user_id"`
	ExpectCredit float32 `json:"expect_credit"`
}

type RegisterData struct {
	MasterAgentId int     `json:"master_agent_id"`
	AgentId       int     `json:"agent_id"`
	UserId        int64   `json:"user_id"`
	DefaultCredit float32 `json:"default_credit"`
}

type QueryData struct {
	UserId int64 `json:"user_id"`
}

// response data
type ResponseData struct {
	Credit float32 `json:"credit"`
}
