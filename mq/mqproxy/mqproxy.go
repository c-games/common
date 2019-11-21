package mqproxy

import (
	"encoding/json"
	"github.com/spf13/viper"
	"github.com/streadway/amqp"
	"gitlab.3ag.xyz/backend/common/app"
	"gitlab.3ag.xyz/backend/common/fail"
	"gitlab.3ag.xyz/backend/common/mq"
	"gitlab.3ag.xyz/backend/common/mq/msg"
	"time"
)

var MqProxy = make(chan app.ProxyData)

var chMap = make(map[int64]chan msg.CGResponseMessage)
var responseQueueName string
var orderQueue string
var userQueue string
var walletQueue string
var loggerQueue string
var gameQueue string

func StartListenMqProxy(chAdp mq.IChannelAdapter) {
	for req := range MqProxy {
		chMap[req.ReqData.Serial] = req.ResChan

		body, err := json.Marshal(req.ReqData)
		fail.FailedOnError(err, "marshal failed")

		err = chAdp.Publish("", req.Target, false, false,
			amqp.Publishing{
				ContentType: "application/json",
				Timestamp:   time.Now(),
				Body:        body,
			})
	}
}

func StartListenMqResponse(chAdp mq.IChannelAdapter) {
	responseQueueName = viper.GetString("response_queue")
	orderQueue = viper.GetString("order_queue")
	walletQueue = viper.GetString("wallet_queue")
	userQueue = viper.GetString("user_queue")
	loggerQueue = viper.GetString("logger_queue")
	gameQueue = viper.GetString("game_queue")

	if responseQueueName == "" {
		panic("lost config response_queue")
	}
	adminResponseQueue, err := chAdp.QueueDeclare(responseQueueName, true, false, false, false, nil)

	if err != nil {
		fail.FailedOnError(err, "Create user response queue failed")
	}
	respQueue := adminResponseQueue.Consume("", false, false, false, false, nil)
	go func() {
		/*defer CapturePanic(func(r interface{}) {
			fmt.Printf("panic: %s", r)
		})*/

		for resp := range respQueue {
			var respMsg msg.CGResponseMessage
			_ = json.Unmarshal(resp.Body, &respMsg)
			resp.Ack(false)
			// fmt.Printf("%+v\n", string(resp.Body))

			if ch, ok := chMap[respMsg.Serial]; !ok {
				// NOTE 如果不 ok 的話就要跳錯，一個沒有 register 過的 msg
				//fmt.Println("Unexpect message " + respMsg.Command)
				//fail.FailedOnError(errors.New("Unexpect message "+respMsg.Command), "")
			} else {
				ch <- respMsg
			}
		}
	}()
}
