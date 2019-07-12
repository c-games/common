package user

import (
	"encoding/json"
	"gitlab.3ag.xyz/backend/common/mq/msg"
	"time"
)

var (
	Register        msg.ServiceCommand = msg.NewCommand("register")
	Login           msg.ServiceCommand = msg.NewCommand("login")
	Logout          msg.ServiceCommand = msg.NewCommand("logout")
	Token           msg.ServiceCommand = msg.NewCommand("token")
	Validate        msg.ServiceCommand = msg.NewCommand("validate")
	Query           msg.ServiceCommand = msg.NewCommand("query")
	Update          msg.ServiceCommand = msg.NewCommand("update")
	QueryUserOdds   msg.ServiceCommand = msg.NewCommand("query_user_odds")
	QueryUserLimit  msg.ServiceCommand = msg.NewCommand("query_user_limit")
	QueryUserRefund  msg.ServiceCommand = msg.NewCommand("query_user_refund")
	UpdateUserOdds  msg.ServiceCommand = msg.NewCommand("update_user_odds")
	UpdateUserLimit msg.ServiceCommand = msg.NewCommand("update_user_limit")
	UpdateUserRefund  msg.ServiceCommand = msg.NewCommand("update_user_refund")
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

// TODO deprecated update struct 用不到
type UpdateData struct {
	UserId   int64     `json:"user_id"`
	// GameId   int       `json:"game_id"`
	Account  string    `json:"account"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	Mobile   string    `json:"mobile"`
	Qq       string    `json:"qq"`
	Wechat   string    `json:"wechat"`
	Ip       string    `json:"ip"`
	Time     time.Time `json:"time"`
	Platform string    `json:"platform"`
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
	Birthday time.Time `json:"birthday"`
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
	TokenExpire time.Time `json:"token_expire"`
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

