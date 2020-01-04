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

func JoinInt(arrInt []int, delemeter string) string {
	if len(arrInt) == 0 {
		return ""
	}

	intString := strconv.FormatInt(int64(arrInt[0]), 10)
	for _, id := range arrInt[1:] {
		intString = intString + delemeter + strconv.FormatInt(int64(id), 10)
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

// TODO fix length = 1
func JoinString(v []string, delimiter string) string {
	if len(v) == 0 {
		return ""
	}

	// workaround
	if len(v) == 1 {
		return v[0]
	}

	rlt := ""
	for _, v := range v {
		rlt = rlt + v + delimiter
	}
	rlt = rlt[:len(rlt) - len(delimiter)]
	return rlt
}