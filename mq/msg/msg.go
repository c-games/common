package msg

import (
	"gitlab.3ag.xyz/backend/common/fail"
	"encoding/json"
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
	return json
}
