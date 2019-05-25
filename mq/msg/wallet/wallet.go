package wallet

import "gitlab.3ag.xyz/backend/common/mq/msg"

var (
	Update msg.ServiceCommand = msg.NewCommand("update")
	Validate msg.ServiceCommand = msg.NewCommand("validate")
	Register msg.ServiceCommand = msg.NewCommand("register")
)

type UpdateData struct {
	UserId int `json:"user_id"`
	CreditChange float64 `json:"credit_change"`
}

type ValidateData struct {
	UserId int `json:"user_id"`
	ExpectCredit float64 `json:"expect_credit"`
}

type ValidateResponseData struct {
	Enough bool `json:"enough"`
}

type RegisterData struct {
	UserId int `json:"user_id"`
	DefaultCredit float64 `json:"default_credit"`
}

type ResponseData struct {
	Message string `json:"message"`
	Credit float64 `json:"credit"`
	ErrorCode int `json:"error_code"`
}
