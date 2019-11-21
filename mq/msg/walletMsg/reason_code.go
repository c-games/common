package walletMsg

const (
	Reason_Bid          int = 1 // 下注
	Reason_Reward       int = 2 // 中獎
	Reason_Withdraw     int = 3 // 撤獎
	Reason_Transfer_In  int = 6
	Reason_Transfer_Out int = 7
	Reason_Flat         int = 8  // 平局


)

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
