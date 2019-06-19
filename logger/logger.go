package logger

import (
	"encoding/json"
	"errors"
	"fmt"
	"gitlab.3ag.xyz/backend/common/fail"
	"gitlab.3ag.xyz/backend/common/mq"
	"gitlab.3ag.xyz/backend/common/mq/msg"
)

type loggerStruct struct {
	channel chan string
	mqChannel mq.IChannelAdapter
}

type LogMessage struct {
	Message string `json:"message"`
}

var logger *loggerStruct
var serviceName string

func Init(name string, loggerQueueName string, channel mq.IChannelAdapter) {
	if logger == nil {
		serviceName = name
		logger = &loggerStruct{
			channel: make(chan string),
			mqChannel: channel,
		}

		_, _ = channel.QueueDeclare(loggerQueueName, true, false, false, false, nil)

		go func() {
			for logFromChann := range logger.channel {
				logMsg := LogMessage{ Message: logFromChann }
				logJson, err := json.Marshal(logMsg)
				fail.FailOnError(err, "Parse Log Message Failed")

				err = logger.mqChannel.PublishService(
					msg.Logger,
					msg.CGMessage{
						Data: []json.RawMessage{logJson},
					},
				)
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
