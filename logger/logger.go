package logger

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/streadway/amqp"
	"gitlab.3ag.xyz/backend/common/fail"
	"gitlab.3ag.xyz/backend/common/mq"
	"gitlab.3ag.xyz/backend/common/mq/msg"
	time2 "gitlab.3ag.xyz/backend/common/timeutil"
	"time"
)

type LogStruct struct {
	Type string
	Command string
	Data interface{}
}

type loggerStruct struct {
	channel chan LogStruct
	mqChannel mq.IChannelAdapter
}

var logger *loggerStruct
var serviceName string

func Init(name string, loggerQueueName string, channel mq.IChannelAdapter) {
	if logger == nil {
		serviceName = name

		logger = &loggerStruct{
			channel: make(chan LogStruct),
			mqChannel: channel,
		}

		_, _ = channel.QueueDeclare(loggerQueueName, true, false, false, false, nil)

		go func() {
			for log := range logger.channel {
				var body []byte
				switch log.Type {
				case "log":
					body = prepareLogMessage(log.Command, log.Data)

					err := logger.mqChannel.Publish(
					"",
					msg.Logger.QueueName(),
					false,
					false,
					amqp.Publishing{
						ContentType: "application/json",
						Body:        body,
						Timestamp:   time.Now(),
					})
					fail.FailOnError(err, "[logger-gorutine] publish faled")
					logdata, _ := json.Marshal(log.Data)

					fmt.Printf("[%s] %s\n", time2.Now(), string(logdata))
				// TODO 分更多的類型
				case "print":
					fmt.Printf("[%s] %s\n", time2.Now(), log.Data)
				default:
					fail.FailOnError(errors.New("unknown log command"), "")
				}
			}
		}()
	}
}

func prepareLogMessage(command string, data interface{}) []byte {

	record, err := json.Marshal(data)
	fail.FailOnError(err, "marshal failed")

	logmsg, err := json.Marshal(msg.CGMessage{
		Command: command,
		Data: []json.RawMessage{record},
	})

	fail.FailOnError(err, "marshal failed")
	return logmsg
}

func assertLoggerAvailable() {
	if logger == nil {
		fail.FailOnError( errors.New("Logger channel is nil "), "Failed to send log")
	}
}

func Logf(format string, args...interface{}) {
	msg := fmt.Sprintf(format, args...)
	Log(msg)
}

func Log(message string) {
	assertLoggerAvailable()
	logger.channel <- LogStruct{
		Type: "print",
		Data: message,
	}
}

// TODO remove who
func PrintRemote(serial int64, who, action, result, message string ) {
	assertLoggerAvailable()
	log := LogStruct{
		Type: "log",
		Command: "print",
		Data: msg.PrintRecord{
			Serial: serial,
			Who: serviceName,
			Action: action,
			Result: result,
			Message: message,
		},
	}
	logger.channel <- log
}
