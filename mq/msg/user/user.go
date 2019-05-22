package user

import (
	"gitlab.3ag.xyz/core/backend/common/mq/msg"
	"time"
)


var (
	Register msg.ServiceCommand = msg.NewCommand("register")

	Login    msg.ServiceCommand = msg.NewCommand("login")

	Logout   msg.ServiceCommand = msg.NewCommand("logout")

	Token    msg.ServiceCommand = msg.NewCommand("token")

	Update   msg.ServiceCommand = msg.NewCommand("update")

	Validate msg.ServiceCommand = msg.NewCommand("validate")

	Fetch msg.ServiceCommand = msg.NewCommand("fetchdata")
)



// TODO rename SelfFetch
type FetchData struct {
	Token string `json:"token"`
}

type FetchResponseData struct {
	Account string      `json:"account"`
	LoginStatus string  `json:"login_status"`
	UserId string       `json:"user_id"`
	Token string        `json:"token"`
	Status string       `json:"status"`
	Name string         `json:"name"`
	Birthday time.Time  `json:"birthday"`
	Mobile string       `json:"mobile"`
	Qq string           `json:"qq"`
	Wechat string       `json:"wechat"`
	Email string        `json:"email"`
	ErrorCode int       `json:"error_code"`
}

type LoginData struct {
	Account  string `json:"account"`
	Password string `json:"password"`
	Token    string `json:"token"`
	Ip       string `json:"ip"`
	Time     string `json:"time"`
	Platform string `json:"platform"`
}

type RegisterData struct {
	AgentId string `json:"agent_id"`
	Account string `json:"account"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Mobile    string `json:"mobile"`
	Qq       string `json:"qq"`
	Wechat   string `json:"wechat"`
	Ip       string `json:"ip"`
	Time     string `json:"time"`
	Platform string `json:"platform"`
}

type LogoutData struct {
	Token   string `json:"token"`
}

type ValidateData struct {
	Token    string `json:"token"`
}