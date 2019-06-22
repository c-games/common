package wallet

import "gitlab.3ag.xyz/backend/common/mq/msg"

var (
	Update msg.ServiceCommand = msg.NewCommand("update")
	Validate msg.ServiceCommand = msg.NewCommand("validate")
	Register msg.ServiceCommand = msg.NewCommand("register")
	Query msg.ServiceCommand = msg.NewCommand("query")
	)

func WalletCommand(commandString string) msg.ServiceCommand {
	switch commandString {
	case "update":
		return Update
	case "validate":
		return Validate
	case "register":
		return Register
	case "query":
		return Query
	default:
		return msg.NullCommand
	}
}

type UpdateData struct {
	UserId        int64   `json:"user_id"`
	CreditChange  float64 `json:"credit_change"`
	Reason        int     `json:"reason"`
	ReasonMessage string  `json:"reason_message"`
}

type ValidateData struct {
	UserId       int64   `json:"user_id"`
	ExpectCredit float64 `json:"expect_credit"`
}

type RegisterData struct {
	UserId        int64   `json:"user_id"`
	DefaultCredit float64 `json:"default_credit"`
}

type QueryData struct {
	UserId int64 `json:"user_id"`
}

type ResponseData struct {
	Credit float64 `json:"credit"`
}
