package msg

import (
	"encoding/json"
	"fmt"
	appErr "gitlab.3ag.xyz/backend/common/app/err"
	"gitlab.3ag.xyz/backend/common/fail"
)

type CGMessage struct {
	Serial        int64             `json:"serial"`
	ResponseQueue string            `json:"response_queue"`
	Command       string            `json:"command"`
	WaitResponse  bool              `json:"wait_response"`
	Data          []json.RawMessage `json:"data"`
}

func (msg *CGMessage) String() string {
	return fmt.Sprintf("Serial:%v, ResQueue:%s, Command:%s, WaitRes:%v, Data:%s",
		msg.Serial, msg.ResponseQueue, msg.Command, msg.WaitResponse, msg.Data)
}

type CGResponseMessage struct {
	Serial       int64             `json:"serial"`
	Command      string            `json:"command"`
	ErrorCode    int               `json:"error_code"`
	ErrorMessage string            `json:"error_message"`
	Data         []json.RawMessage `json:"data"`
}

func (msg *CGResponseMessage) String() string {
	return fmt.Sprintf("Serial: %v, Command:%s, ErrorCode:%v, ErrorMsg:%s, Data:%s",
		msg.Serial, msg.Command, msg.ErrorCode, msg.ErrorMessage, msg.Data)
}

type IServiceData interface {
	// TODO 如果訂了這個就變成是 data type 都要 implement 了
	// ToJson() string
	// ToByteArray() []byte
}

type PrintRecord struct {
	Serial  int64  `json:"serial"`
	Time    string `json:"time"`
	Service string `json:"service"`
	Level   int    `json:"level"` // 0: info 1: warn 2: debug 3: error 4: fatal
	Message string `json:"message"`
}

type MessageData interface {
	Get(string) interface{}
	ToStruct()
}

func (msg *CGMessage) FirstData() json.RawMessage {
	return msg.Data[0]
}

func ToStruct(body []byte) CGMessage {
	var cgMsg CGMessage
	err := json.Unmarshal(body, &cgMsg)
	fail.FailedOnError(err, "unmarshal json failed")
	return cgMsg
}

func ToJson(d interface{}) string {
	return string(ToByteArray(d))
}

func ToByteArray(d interface{}) []byte {
	json, err := json.Marshal(d)
	fail.FailedOnError(err, "parse json failed")
	return json
}

func PackCgMessage(serial int64, data []byte) []byte {
	cgMessage := &CGMessage{
		Serial: serial,
		Data:   []json.RawMessage{data},
	}

	result, err := json.Marshal(cgMessage)
	fail.FailedOnError(err, "Marshal Json failed")
	return result
}

func PackCgResponseError(cgMessage CGMessage, errorCode int, errMessage string) CGResponseMessage {
	if errorCode == appErr.Code_Success {
		panic("wrong error code")
	}
	return CGResponseMessage{
		Serial:       cgMessage.Serial,
		Command:      cgMessage.Command,
		ErrorCode:    errorCode,
		ErrorMessage: errMessage,
		Data:         []json.RawMessage{},
	}

}

// TODO rename
func PackCgResponseMessage2(cgMessage CGMessage, errorCode int, data interface{}) CGResponseMessage {
	// TODO check data is not []byte
	// data must be a struct with json tag
	var resData []json.RawMessage
	if data == nil {
		resData = []json.RawMessage{}
	} else {
		jsonRes, err := json.Marshal(data)
		fail.FailedOnError(err, "marshal failed")
		resData = []json.RawMessage{jsonRes}
	}
	return CGResponseMessage{
		Serial:    cgMessage.Serial,
		Command:   cgMessage.Command,
		ErrorCode: errorCode,
		Data:      resData,
	}
}
func PackCgResponseMessageMany2(cgMessage CGMessage, errorCode int, data []interface{}) CGResponseMessage {
	// TODO check data is not []byte
	// data must be a struct with json tag
	var encodeData []json.RawMessage
	if len(data) == 0 {
		encodeData = []json.RawMessage{}
	} else {
		for _, datum := range data {
			jsonDatum, err := json.Marshal(datum)
			fail.FailedOnError(err, "marshal failed")
			encodeData = append(encodeData, jsonDatum)
		}
	}

	return CGResponseMessage{
		Serial:    cgMessage.Serial,
		Command:   cgMessage.Command,
		ErrorCode: errorCode,
		Data:      encodeData,
	}
}

func PackCgResponseMessage(cgMessage CGMessage, errorCode int, data []byte) CGResponseMessage {
	if data == nil {
		return CGResponseMessage{
			Serial:    cgMessage.Serial,
			Command:   cgMessage.Command,
			ErrorCode: errorCode,
			Data:      []json.RawMessage{},
		}
	} else {
		return CGResponseMessage{
			Serial:    cgMessage.Serial,
			Command:   cgMessage.Command,
			ErrorCode: errorCode,
			Data:      []json.RawMessage{data},
		}
	}

}

func PackCgResponseMessageMany(cgMessage CGMessage, errorCode int, data []json.RawMessage) CGResponseMessage {

	return CGResponseMessage{
		Serial:    cgMessage.Serial,
		Command:   cgMessage.Command,
		ErrorCode: errorCode,
		Data:      data,
	}

}

func SerializeCgResponseMessage(cgMessage CGMessage, errorCode int, data []byte) []byte {
	cgRes := PackCgResponseMessage(cgMessage, errorCode, data)
	result, err := json.Marshal(cgRes)
	fail.FailedOnError(err, "Marshal Json failed")
	return result
}
