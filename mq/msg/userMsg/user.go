package userMsg

import (
	"encoding/json"
)

type QueryData struct {
	Id int64 `json:"user_id"`
}

type LoginData struct {
	AgentId  int    `json:"agent_id"`
	Account  string `json:"account"`
	Password string `json:"password"`
}

type LoginInfoData struct {
	UserId   int64  `json:"user_id"`
	Ip       string `json:"ip"`
	Platform int    `json:"platform"`
	Browser  int    `json:"browser"`
}

type RegisterData struct {
	AgentId  int    `json:"agent_id"`
	Account  string `json:"account"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Mobile   string `json:"mobile"`
	Qq       string `json:"qq"`
	Wechat   string `json:"wechat"`
	Ip       string `json:"ip"`
	Platform int    `json:"platform"`
}

type LogoutData struct {
	Id int64 `json:"user_id"`
}

type GeneralData struct {
	Id int64 `json:"user_id"`
}

type ValidateData struct {
	Token string `json:"token"`
}

type QueryUserData struct {
	UserId int64 `json:"user_id"`
	GameId int   `json:"game_id"`
}

type UpdateUserOddsData struct {
	UserId int64  `json:"user_id"`
	GameId int    `json:"game_id"`
	Odds   string `json:"odds"`
}

type UpdateUserLimitData struct {
	UserId int64  `json:"user_id"`
	GameId int    `json:"game_id"`
	Limit  string `json:"limit"`
}

type UpdateUserRefundData struct {
	UserId int64  `json:"user_id"`
	GameId int    `json:"game_id"`
	Refund string `json:"refund"`
}

type RegisterGameData struct {
	GameId int             `json:"game_id"`
	Name   string          `json:"name"`
	Odds   json.RawMessage `json:"odds"`
	Limit  json.RawMessage `json:"limit"`
	Refund json.RawMessage `json:"refund"`
}

type RegisterAgentData struct {
	Id            int    `json:"agent_id"`
	Account       string `json:"account"`
	Password      string `json:"password"`
	Name          string `json:"name"`
	SiteUrl       string `json:"site_url"`
	MasterAgentId int    `json:"master_agent_id"`
}

type QueryByAgent struct {
	AgentId int `json:"agent_id"`
}

type QueryAgentData struct {
	Id int `json:"agent_id"`
}

// Response data:
type QueryUserResponse struct {
	UserId int64           `json:"user_id"`
	GameId int             `json:"game_id"`
	Odds   json.RawMessage `json:"odds"`
}

type QueryResponse struct {
	UserId        int64   `json:"user_id"`
	AgentId       int     `json:"agent_id"`
	MasterAgentId int     `json:"master_agent_id"`
	Account       string  `json:"account"`
	Name          string  `json:"name"`
	Email         string  `json:"email"`
	Birthday      string  `json:"birthday"`
	Mobile        string  `json:"mobile"`
	Qq            string  `json:"qq"`
	Wechat        string  `json:"wechat"`
	Credit        float32 `json:"credit"`
}

type QueryStatusResponse struct {
	UserId      int64 `json:"user_id"`
	Status      int   `json:"status"`
	LoginStatus int   `json:"login_status"`
}

type LoginResponse struct {
	NewToken    string `json:"new_token"`
	TokenExpire string `json:"token_expire"`
}

type ValidateResponse struct {
	Id int64 `json:"user_id"`
	//	MasterAgentId int    `json:"master_agent_id"`
	AgentId int    `json:"agent_id"`
	Account string `json:"account"`
}

type QueryUserLimitResponse struct {
	UserId int64           `json:"user_id"`
	GameId int             `json:"game_id"`
	Limit  json.RawMessage `json:"limit"`
}

type QueryUserRefundResponse struct {
	UserId int64           `json:"user_id"`
	GameId int             `json:"game_id"`
	Refund json.RawMessage `json:"refund"`
}

type QueryAgentResponse struct {
	Id            int    `json:"agent_id"`
	Account       string `json:"account"`
	Name          string `json:"name"`
	MasterAgentId int    `json:"master_agent_id"`
	Role          int    `json:"role"`
}

type FindUserId struct {
	AgentId int    `json:"agent_id"`
	Account string `json:"account"`
}
