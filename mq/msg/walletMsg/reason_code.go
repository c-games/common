package walletMsg

const (
	Reason_Other int = 0
	Reason_Bid          int = 1 // 下注
	Reason_Reward       int = 2 // 中獎
	Reason_Withdraw     int = 3 // 撤獎
	Reason_Transfer_In  int = 6
	Reason_Transfer_Out int = 7
	Reason_Flat         int = 8  // 平局
Reason_MaxRange int = 9

)

func IsAvailableReason(reason int) bool {
	switch reason {
	case Reason_Bid, Reason_Reward, Reason_Withdraw, Reason_Transfer_In, Reason_Transfer_Out, Reason_Flat:
		return true
	default:
		return false
	}
}


func GetReasonString(reasonCode int) string {
	switch reasonCode {
	case Reason_Bid:
		return "下注"
	case Reason_Reward:
		return "中奬"
	case Reason_Withdraw:
		return "撤獎"
	case Reason_Transfer_In:
		return "入帳"
	case Reason_Transfer_Out:
		return "出帳"
	case Reason_Flat:
		return "開獎平局"

	default:
		panic("不合法的 wallet reason")
	}
}
