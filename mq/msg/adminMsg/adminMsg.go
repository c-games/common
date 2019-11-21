package adminMsg

import "encoding/json"

type LoginData struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

type LogoutData struct {
	Token   string `json:"token"`
	Account string `json:"account"`
}

type FetchInfoData struct {
	//UserId int `json:"user_id"`
	Token   string `json:"token"`
	Account string `json:"account"`
}

type ValidateData struct {
	Token   string `json:"token"`
	Account string `json:"account"`
}

type CreateUserData struct {
	// Id          int    `json:"id"`
	Token       string `json:"token"`
	Account     string `json:"account"`
	Password    string `json:"password"`
	Name        string `json:"name"`
	Email       string `json:"name"`
	Category    int    `json:"category"`
	Key         string `json:"key"`
	Belongs     int    `json:"belongs"`
	AccountType string `json:"account_type"` // head_office or master_agent or agent or sub_account
}

type QueryData struct {
	GameId  int `json:"game_id"`
	AgentId int `json:"agent_id"`
}

type FetchAgentConfig struct {
	AgentId int `json:"agent_id"`
}

type CreatePermissionData struct {
	// TODO
}

type FetchAllMasterAgent struct {
	// Token string `json:"token"` // HeadOffice level token
	Id int `json:"id"`
}

type FetchAllAgent struct {
	// Token string `json:"token"` // HeadOffice level token
	Id int `json:"id"`
}

type FetchAllSubAccount struct {
	Token string `json:"token"` // HeadOffice level token
}

// response
type LoginResponse struct {
	Token string `json:"token"`
}

type LogoutResponse struct {
}

type QueryResponse struct {
	AgentId int    `json:"agent_id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Role    int    `json:"role"`
	IsSub   bool   `json:"is_sub"`
}

type QueryLimitResponse struct {
	AgentId int             `json:"agent_id"`
	GameId  int             `json:"game_id"`
	Limit   json.RawMessage `json:"limit"`
}

type QueryOddsResponse struct {
	AgentId int             `json:"agent_id"`
	GameId  int             `json:"game_id"`
	Odds    json.RawMessage `json:"odds"`
}

type QueryRefundResponse struct {
	AgentId int             `json:"agent_id"`
	GameId  int             `json:"game_id"`
	Refund  json.RawMessage `json:"refund"`
}

type FetchInfoResponse struct {
	Id         int      `json:"id"`
	Account    string   `json:"account"`
	Name       string   `json:"name"`
	Email      string   `json:"email"`
	Status     int      `json:"status"`
	Category   int      `json:"category"`
	Key        string   `json:"key"`
	Role       int      `json:"role"`
	Permission []string `json:"permits"`
	Belongs    int      `json:"belongs"`
}

type FetchAllMasterAgentResponse struct {
	Id          int     `json:"id"`
	Account     string  `json:"account"`
	Name        string  `json:"name"`
	Email       string  `json:"email"`
	Status      int     `json:"status"`
	Category    int     `json:"category"`
	LoginStatus int     `json:"login_status"`
	Role        int     `json:"role"`
	Parent      int     `json:"parent"`
	Credit      float32 `json:"credit"`
	CreateDate  string  `json:"create_time"`
}

type FetchAllAgentResponse FetchAllMasterAgentResponse
