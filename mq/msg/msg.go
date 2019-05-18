package msg

import (
	"encoding/json"
	"gitlab.3ag.xyz/core/backend/common/fail"
)

type Content struct {
	Type          string             `json:type`
	Serial        string             `json:"serial"`
	RouteId       string             `json:"id"`
	// UserId        string             `json:"user"`
	ServiceName   Service            `json:"service_name"`
	Command       string             `json:"command"`
	Data          []json.RawMessage  `json:"data"`
	WaitResponse  bool               `json:"wait_response"`
}

type Service string
const (
	Order  Service = "order"
	User   Service = "user"
	Wallet Service = "wallet"
)

type MessageData interface {
	Get(string) interface{}
	ToStruct()
}


func ToStruct(body []byte) Content {
	var reqMsg Content
	err := json.Unmarshal(body, &reqMsg)
	fail.FailOnError(err, "unmarshal json failed")
	return reqMsg
}

// TODO remove
type ResponseData struct {
	Serial  string `json:serial`
	OriginRouteId string `json:origin_route_id`
	Service string `json:service`
	UserId string `json:user_id`
	OriginCommand string `json:origin_command`
	ErrorCode int `json:error_code`
	Args interface{} `json:args`
}
