package msg

import (
	"encoding/json"
	"github.com/c-games/common/fail"
)

type MessageData interface {
	Get(string) interface{}
	ToStruct()
}

func ToJson(d interface{}) string {
	return string(ToByteArray(d))
}

func ToByteArray(d interface{}) []byte {
	json, err := json.Marshal(d)
	fail.FailedOnError(err, "parse json failed")
	return json }
