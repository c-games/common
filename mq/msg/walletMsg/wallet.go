package walletMsg

type UpdateData struct {
	MasterAgentId int     `json:"master_agent_id"`
	AgentId       int     `json:"agent_id"`
	UserId        int64   `json:"user_id"`
	CreditChange  float64 `json:"credit_change"`
	Reason        int     `json:"reason"`
	ReasonMessage string  `json:"reason_message"`
}

type ValidateData struct {
	MasterAgentId int     `json:"master_agent_id"`
	AgentId       int     `json:"agent_id"`
	UserId        int64   `json:"user_id"`
	ExpectCredit  float64 `json:"expect_credit"`
}

type RegisterData struct {
	MasterAgentId int     `json:"master_agent_id"`
	AgentId       int     `json:"agent_id"`
	UserId        int64   `json:"user_id"`
	DefaultCredit float64 `json:"default_credit"`
}

type QueryData struct {
	MasterAgentId int   `json:"master_agent_id"`
	AgentId       int   `json:"agent_id"`
	UserId        int64 `json:"user_id"`
}

// response data

type ResponseData struct {
	Credit float64 `json:"credit"`
}
