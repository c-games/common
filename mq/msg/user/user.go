package user

import (
	"encoding/json"
)

type QueryData struct {
	Id int64 `json:"user_id"`
}

type LoginData struct {
	MasterAgentId int `json:"master_agent_id"`
	AgentId  int    `json:"agent_id"`
	Account  string `json:"account"`
	Password string `json:"password"`
	Ip       string `json:"ip"`
	Platform int `json:"platform"`
}

type RegisterData struct {
	MasterAgentId int    `json:"master_agent_id"`
	AgentId       int    `json:"agent_id"`
	Account       string `json:"account"`
	Name          string `json:"name"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	Mobile        string `json:"mobile"`
	Qq            string `json:"qq"`
	Wechat        string `json:"wechat"`
	Ip            string `json:"ip"`
	Platform      int    `json:"platform"`
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
	UserId    int64  `json:"user_id"`
	GameId int    `json:"game_id"`
	Odds      string `json:"odds"`
}

type UpdateUserLimitData struct {
	UserId int64  `json:"user_id"`
	GameId int    `json:"game_id"`
	Limit  string `json:"limit"`
}

type UpdateUserRefundData struct {
	UserId int64  `json:"user_id"`
	GameId int    `json:"game_id"`
	Refund  string `json:"refund"`
}

type RegisterGameData struct {
	GameId int `json:"game_id"`
	Name int `json:"name"`
	Odds json.RawMessage `json:"odds"`
	Limit json.RawMessage `json:"json"`
	Refund json.RawMessage`json:"refund"`
}

type RegisterAgentData struct {
	Id int `json:"id"`
	Name string `json:"name"`
	SiteUrl string `json:"site_url"`
}

// Response data:
type QueryUserResponse struct {
	UserId int64  `json:"user_id"`
	GameId int    `json:"game_id"`
	Odds   string `json:"odds"`
}

type QueryResponse struct {
	UserId   int64     `json:"user_id"`
	AgentId  int       `json:"agent_id"`
	MasterAgentId int `json:"master_agent_id"`
	Account  string    `json:"account"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Birthday string `json:"birthday"`
	Mobile   string    `json:"mobile"`
	Qq       string    `json:"qq"`
	Wechat   string    `json:"wechat"`
}

type LoginResponse struct {
	MasterAgentId int `json:"master_agent_id"`
	AgentId     int       `json:"agent_id"`
	UserId      int64  `json:"user_id"`
	Account     string    `json:"account"`
	NewToken    string    `json:"new_token"`
	TokenExpire string `json:"token_expire"`
}

type ValidateResponse struct {
	Id int64 `json:"user_id"`
}

type QueryUserLimitResponse struct {
	UserId int64  `json:"user_id"`
	GameId int    `json:"game_id"`
	Limit  string `json:"limit"`
}

type QueryUserRefundResponse struct {
	UserId int64  `json:"user_id"`
	GameId int    `json:"game_id"`
	Refund  string `json:"refund"`
}

