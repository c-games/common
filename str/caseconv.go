package str

import "strings"

func Pascal2Snake(str string) string {

	if !IsUppercase(rune(str[0])) {
		panic("not PascalCase")
	}

	rlt := ""
	for _, c := range str {
		if IsUppercase(c) {
			rlt += "_" + strings.ToLower(string(c))
		} else if IsNumber(c) {
			rlt += "_" + strings.ToLower(string(c))
		} else {
			rlt += string(c)
		}
	}

	return rlt[1:]
}


func IsUppercase(c rune) bool {
	return  c >= 'A' && c <= 'Z'
}

func IsNumber(c rune) bool {
	return c >= '0' && c <= '9'
}