package orderMsg

// User Error Code Defined

const (
	Code_Success            int = 0
	Code_WrongOrderId       int = 1
	Code_CreditNotEnough    int = 2
	Code_WrongGameIdOrRound int = 3
	Code_InvalidateUser     int = 4
	Code_DbFailed           int = 8
)