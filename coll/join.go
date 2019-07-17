package coll

import "strconv"

func JoinInt64(arrInt []int64, delemeter string) string {
	if len(arrInt) == 0 {
		return ""
	}

	intString := strconv.FormatInt(arrInt[0], 10)
	for _, id := range arrInt[1:] {
		intString = intString + delemeter + strconv.FormatInt(id, 10)
	}
	return intString
}

func JoinMapInt(m map[string]int, delimiter string) string {
	if len(m) == 0 {
		return ""
	}

	var arr []string

	for k, v := range m {
		arr = append(arr, k + "=" + strconv.Itoa(v))
	}

	rlt := arr[0] + delimiter
	for _, v := range arr[1:] {
		rlt = rlt + delimiter + v
	}
	return rlt
}