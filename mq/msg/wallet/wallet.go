package wallet

import (
)

type WalletCommand string

const (
	Update WalletCommand = "update"
	Validate WalletCommand = "validate"
	Register WalletCommand = "register"
)

type UpdateData struct {
	Account string `json:"account"`
	UserId string `json:"user_id"`
	CreditChange float64 `json:"credit_change"`
}

type ValidateData struct {
	Account string `json:"account"`
	UserId string `json:"user_id"`
	ExpectCredit float64 `json:"expect_credit"`
}

type ValidateResponseData struct {
	Enough bool `json:"enough"`
}

type RegisterData struct {
	UserId string `json:"user_id"`
	Account string `json:"account"`
	DefaultCredit float64 `json:"default_credit"`
}

type ResponseData struct {
	Message string `json:"message"`
	Credit float64 `json:"credit"`
	ErrorCode int `json:"error_code"`
}
