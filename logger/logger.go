package logger

import (
	"time"
	"encoding/json"
	"fmt"
	"errors"
	"github.com/streadway/amqp"
	"gitlab.3ag.xyz/core/backend/common/mq"
	"gitlab.3ag.xyz/core/backend/common/fail"
)

type loggerStruct struct {
	channel chan string
	mqChannel *mq.MqChann
}

type LogMessage struct {
	Message string `json:"message"`
}

var logger *loggerStruct

func Init(ch *mq.MqChann) {
	if logger == nil {
		logger = &loggerStruct{
			channel: make(chan string),
			mqChannel: ch,
		}

		ch.GenQueue("cg-log", true, false, false, false, nil)

		go func() {
			for logFromChann := range logger.channel {
				logMsg := LogMessage{ Message: logFromChann }
				logJson, err := json.Marshal(logMsg)
				fail.FailOnError(err, "Parse Log Message Failed")

				err = logger.mqChannel.Ch.Publish(
					"",
					"cg-log",
					false,
					false,
					amqp.Publishing{
						ContentType: "application/json",
						Body: []byte(logJson),
						Timestamp: time.Now(),
					})
				fail.FailOnError(err, "[logger] publish faled")
			}
		}()
	}
}

func Logf(format string, args...interface{}) {
	msg := fmt.Sprintf(format, args...)
	Log(msg)
}

func Log(message string) {
	if logger == nil {
		fail.FailOnError( errors.New("Logger channel is nil"), "Failed to send log")
	} else {
		logger.channel <- message
	}
}
