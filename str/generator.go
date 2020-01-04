package str

func RepeatChar(char string, length int) []string {
	var rlt []string
	for i := 0 ; i < length ; i++ {
		rlt = append(rlt, char)
	}
	return rlt
}
