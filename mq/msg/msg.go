package msg

import (
	"encoding/json"
	"gitlab.3ag.xyz/backend/common/fail"
)

type CGMessage struct {
	Serial        int64               `json:"serial"`
	ResponseQueue string            `json:"response_queue"`
	Command       string            `json:"command"`
	WaitResponse  bool              `json:"wait_response"`
	Data          []json.RawMessage `json:"data"`
}

type CGResponseMessage struct {
	Serial        int64               `json:"serial"`
	Command       string            `json:"command"`
	ErrorCode     int               `json:"error_code"`
	Data          []json.RawMessage `json:"data"`
}

type LoggerMessage struct {
	Id string `json:"id"`
	Record json.RawMessage `json:"record"`
}

type IServiceData interface {
	// TODO 如果訂了這個就變成是 data type 都要 implement 了
	// ToJson() string
	// ToByteArray() []byte
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

func PackCgMessage(serial int64, data[]byte) []byte {
	cgMessage := &CGMessage{
		Serial: serial,
		Data: []json.RawMessage{data},
	}

	result, err := json.Marshal(cgMessage)
	fail.FailOnError(err, "Marshal Json failed")
	return result
}

func PackCgResponseMessage(cgMessage CGMessage, errorCode int, data []byte) CGResponseMessage {
	if data == nil {
		return CGResponseMessage{
			Serial: cgMessage.Serial,
			Command: cgMessage.Command,
			ErrorCode: errorCode,

		}
	} else {
		return CGResponseMessage{
			Serial: cgMessage.Serial,
			Command: cgMessage.Command,
			ErrorCode: errorCode,
			Data: []json.RawMessage{data},
		}
	}

}

func PackCgResponseMessageMany(cgMessage CGMessage, errorCode int, data []json.RawMessage) CGResponseMessage {

	return CGResponseMessage{
		Serial: cgMessage.Serial,
		Command: cgMessage.Command,
		ErrorCode: errorCode,
		Data: data,
	}

}

func SerializeCgResponseMessage(cgMessage CGMessage, errorCode int, data []byte) []byte {
	cgRes := PackCgResponseMessage(cgMessage, errorCode, data)
	result, err := json.Marshal(cgRes)
	fail.FailOnError(err, "Marshal Json failed")
	return result
}

func CodeUnexpectCommand() int {
	return -1
}