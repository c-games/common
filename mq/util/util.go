package util

import (
	"encoding/json"
	"github.com/streadway/amqp"
	"time"
)



// TODO 對這種可能很多 type 的，用 interface{} 真的好嗎 (不是一個開放 type，而是有限 type)
//func sendResponse(ch *amqp.Channel, targetQueueName string, responseData *userMsg.ResponseData) {

// TODO remove
func SendResponse(ch *amqp.Channel, targetQueueName string, responseData interface{}, serial string) {
	resJson, _ := json.Marshal(responseData)
	ch.Publish(
		targetQueueName,
		"", // route key
		false, //
		false, //
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(resJson),
			Timestamp:   time.Now(),
			AppId: serial,
		})
}

