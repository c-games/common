package user

import (
	"gitlab.3ag.xyz/backend/common/mq/msg"
	"time"
)

var (
	Register msg.ServiceCommand = msg.NewCommand("register")
	Login    msg.ServiceCommand = msg.NewCommand("login")
	Logout   msg.ServiceCommand = msg.NewCommand("logout")
	Token    msg.ServiceCommand = msg.NewCommand("token")
	Validate msg.ServiceCommand = msg.NewCommand("validate")
	Query msg.ServiceCommand = msg.NewCommand("query")
	QueryUserOdds msg.ServiceCommand = msg.NewCommand("query_user_odds")
	QueryUserLimit msg.ServiceCommand = msg.NewCommand("query_user_limit")
	Update   msg.ServiceCommand = msg.NewCommand("update")
	UpdateUserOdds   msg.ServiceCommand = msg.NewCommand("update_user_odds")
	UpdateUserLimit   msg.ServiceCommand = msg.NewCommand("update_user_limit")
)



type QueryData struct {
	Id int64 `json:"user_id"`
}


type QueryResponse struct {
	UserId int64       `json:"user_id"`
	AgentId int        `json:"agent_id"`
	Account string      `json:"account"`
	Name string         `json:"name"`
	Email string        `json:"email"`
	Birthday time.Time  `json:"birthday"`
	Mobile string       `json:"mobile"`
	Qq string           `json:"qq"`
	Wechat string       `json:"wechat"`
}

type LoginData struct {
	AgentId  int `json:"agent_id"`
	Account  string `json:"account"`
	Password string `json:"password"`
	Ip       string `json:"ip"`
	Platform string `json:"platform"`
}

type LoginResponse struct {
	AgentId  int `json:"agent_id"`
	Account  string `json:"account"`
	NewToken string `json:"new_token"`
	TokenExpire time.Time `json:"token_expire"`
}

type RegisterData struct {
	AgentId  int `json:"agent_id"`
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
	Id int64 `json:user_id`
}

type GeneralData struct {
	Id int64 `json:user_id`
}

type ValidateData struct {
	Token    string `json:"token"`
}

type ValidateResponse struct {
	Id    int64 `json:"user_id"`
}

type QueryUserData struct {
	UserId int64 `json:"user_id"`
	LotteryId int `json:"lottery_id"`
}

type UpdateData struct {
	UserId int64 `json:"user_id"`
	LotteryId int `json:"lottery_id"`
	Account  string `json:"account"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Mobile   string `json:"mobile"`
	Qq       string `json:"qq"`
	Wechat   string `json:"wechat"`
	Ip       string `json:"ip"`
	Time     time.Time `json:"time"`
	Platform string `json:"platform"`
}

type UpdateUserOddsData struct {
	UserId int64 `json:"user_id"`
	LotteryId int `json:"lottery_id"`
	Odds string `json:"odds"`
}

type UpdateUserLimitData struct {
	UserId int64 `json:"user_id"`
	LotteryId int `json:"lottery_id"`
	Limit string `json:"limit"`
}

type QueryUserResponse struct {
	UserId int64 `json:"user_id"`
	LotteryId int `json:"lottery_id"`
	Odds string `json:"odds"`
}

type QueryUserLimitResponse struct {
	UserId int64 `json:"user_id"`
	LotteryId int `json:"lottery_id"`
	Limit string `json:"limit"`
}


func UserCommand(commandString string) msg.ServiceCommand {
	switch commandString{
	case "register":
		return Register
	case "login":
		return Login
	case "logout":
		return Logout
	case "token":
		return Token
	case "validate":
		return Validate
	case "query":
		return Query
	case "query_user_odds":
		return QueryUserOdds
	case "query_user_limit":
		return QueryUserLimit
	case "update":
		return Update
	case "update_user_odds":
		return UpdateUserOdds
	case "update_user_limit":
		return UpdateUserLimit
	default:
		return msg.NullCommand
	}
}
