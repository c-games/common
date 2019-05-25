package msg

import (
	"encoding/json"
	"gitlab.3ag.xyz/backend/common/fail"
)

type CGMessage struct {
	Type          string             `json:type`
	Serial        string             `json:"serial"`
	RouteId       string             `json:"id"`
	ResponseName  Service            `json:"service_name"`
	Command       string             `json:"command"`
	WaitResponse  bool               `json:"wait_response"`
	Data          []json.RawMessage  `json:"data"`
}

type IServiceData interface {
//	ToJson() string
	//ToByteArray() []byte
}

type MessageData interface {
	Get(string) interface{}
	ToStruct()
}

func ToStruct(body []byte) CGMessage {
	var cgMsg CGMessage
	err := json.Unmarshal(body, &cgMsg)
	fail.FailOnError(err, "unmarshal json failed")
	return cgMsg
}


func ToJson(d interface{}) string {
	return string(ToByteArray(d))
}

func ToByteArray(d interface{}) []byte {
	json, err := json.Marshal(d)
	fail.FailOnError(err, "parse json failed")
	return json
}