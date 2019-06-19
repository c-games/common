package user

// User Error Code Defined
// TODO must be const
var (
	codeSuccess           int = 0
	codeTokenWrong        int = 1
	codeTokenExpired      int = 2
	codeAccountNotFound   int = 3
	codeWrongPassword     int = 4
	codeRequireFieldEmpty int = 5
	codeLogoutFailed      int = 6
	codeUnknownAgentId    int = 7
	codeIdNotExist        int = 8
	codeDbFailed          int = 9

)
func CodeSuccess() int {
	return codeSuccess
}
func CodeAccountExist() int {
	// NOTE 帳號存在也是用 Success
	return codeSuccess
}


func CodeTokenWrong() int {
	return codeTokenWrong
}
func CodeTokenExpired() int {
	return codeTokenExpired
}
//func CodeLogoutSuccess() int {
//	return codeLogoutSuccess
//}
func CodeFieldEmpty() int {
	return codeRequireFieldEmpty
}
func CodeAccountNotFound() int {
	return codeAccountNotFound
}
func CodeWrongPassword() int {
	return codeWrongPassword
}
func CodeLogoutFailed() int {
	return codeLogoutFailed
}
func CodeLoginTwice() int {
	return codeSuccess
}

func CodeUnknownAgentId() int {
	return codeUnknownAgentId
}

func CodeIdNotExist() int {
	return codeIdNotExist
}

func CodeDbFailed() int {
	return codeDbFailed
}