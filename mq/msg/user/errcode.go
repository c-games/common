package user

// User Error Code Defined
const (
	Code_Success              int = 0
	Code_TokenWrong           int = 1
	Code_TokenExpired         int = 2
	Code_AccountNotFound      int = 3
	Code_WrongPassword        int = 4
	Code_RequireFieldEmpty    int = 5
	Code_LogoutFailed         int = 6
	Code_UnknownAgentId       int = 7
	Code_IdNotExist           int = 8
	Code_AccountAllreadyExist int = 9
	Code_ZeroDataUpdate       int = 10
	Code_DbFailed             int = 11
)
