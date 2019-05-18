package user

import "time"

type UserCommand string


const (
	Register UserCommand = "register"
	Login    UserCommand = "login"
	Logout   UserCommand = "logout"
	Token    UserCommand = "token"
	Update   UserCommand = "update"
	Validate UserCommand = "validate"
	Fetch UserCommand = "fetchdata"
)

type FetchData struct {
	Account string `json:"account"`
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
	Account string `json:"account"`
	Token   string `json:"token"`
}


type ValidateData struct {
	Account  string `json:"account"`
	Token    string `json:"token"`
	UserId   string `json:"user_id"`
}


// TODO deprecated, move to msg
type ResponseData struct {
	Serial  string `json:"serial"`
	Message string `json:"message"`
	ErrorCode int `json:"error_code"`
	UserId string `json:"user_id"`
	Args interface{} `json:"args"`
}
