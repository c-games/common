package logger

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/streadway/amqp"
	"gitlab.3ag.xyz/backend/common/fail"
	"gitlab.3ag.xyz/backend/common/mq"
	"gitlab.3ag.xyz/backend/common/mq/msg"
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
				case "print":
					body = prepareLogMessage("print", log.Data)
				default:
					fail.FailOnError(errors.New("unknown log command"), "")
				}
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
			}
		}()
	}
}

func prepareLogMessage(command string, data interface{}) []byte {

	record, err := json.Marshal(data)
	fail.FailOnError(err, "marshal failed")

	logmsg, err := json.Marshal(msg.LoggerMessage{
		Command: command,
		Record: []json.RawMessage{record},
	})

	fail.FailOnError(err, "marshal failed")
	return logmsg
}

func preparePrintMessage() []byte {
	return nil
}

func assertLoggerAvailable() {
	if logger == nil {
		fail.FailOnError( errors.New("Logger channel is nil"), "Failed to send log")
	}
}


func Logf(format string, args...interface{}) {
	msg := fmt.Sprintf(format, args...)
	Log(msg)
}

func Log(message string) {
	assertLoggerAvailable()
	logger.channel <- LogStruct{
		Type: "log",
		Command: "unknown",
		Data: message,
	}
}
func Print(serial int64, who, action, result, message string ) {
	assertLoggerAvailable()

	log := LogStruct{
		Type: "print",
		Command: "print",
		Data: msg.PrintRecord{
			Serial: serial,
			Who: who,
			Action: action,
			Result: result,
			Message: message,
		},
	}

	logger.channel <- log
}
