package random

import(
	"math/rand"
	"time"
)



var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func Init() {
	rand.Seed(time.Now().UnixNano())
}

func RandStringRunes(n int) string {
	Init()
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func GenString(offset int) string {
	return RandStringRunes(offset)
}
