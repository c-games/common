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
	UserId int64 `json:"user_id"`
	CreditChange float64 `json:"credit_change"`
	Reason int `json:reason`
	ReasonMessage string `json:reason_message`
}

type ValidateData struct {
	UserId int64 `json:"user_id"`
	ExpectCredit float64 `json:"expect_credit"`
}

type RegisterData struct {
	UserId int64 `json:"user_id"`
	DefaultCredit float64 `json:"default_credit"`
}

type QueryData struct {
	UserId int64 `json:"user_id"`
}

type ResponseData struct {
	Credit float64 `json:"credit"`
}

// TODO must be const
var (
	codeSuccess         int = 0
	codeWalletNotFound  int = 1
	codeCreditNotEnough int = 2
	codeRegisterFailed  int = 3
	codeNoUpdateReason  int = 4
)
func CodeSuccess() int {
	return codeSuccess
}

func CodeWalletNotFound() int {
	return codeWalletNotFound
}

func CodeCreditNotEnough() int {
	return codeCreditNotEnough
}

func CodeRegisterFailed() int {
	return codeRegisterFailed
}

func CodeNoUpadetReason() int {
	return codeNoUpdateReason
}

// Reason
var (
	reasonOther    int = 0
	reasonBid      int = 1 // 下注
	reasonReward   int = 2 // 中獎
	reasonWithdraw int = 3 // 撤獎
	reasonRollback int = 4
	reasonUpdateOrder int = 5
)

func ReasonBid() int {
	return reasonBid
}

func ReasonReward() int {
	return reasonReward
}

func ReasonWithdraw() int {
	return reasonWithdraw
}

func ReasonRollback() int {
	return reasonRollback
}

func ReasonUpdateOrder() int {
	return reasonUpdateOrder
}