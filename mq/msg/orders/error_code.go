package orders

// User Error Code Defined
// TODO must be const
var (
	codeSuccess           int = 0
	codeWrongOrderId      int = 1

	codeDbFailed          int = 9

)
func CodeSuccess() int {
	return codeSuccess
}
func CodeAccountExist() int {
	// NOTE 帳號存在也是用 Success
	return codeSuccess
}

func CodeWrongOrderId() int {
	return codeWrongOrderId
}


func CodeDbFailed() int {
	return codeDbFailed
}