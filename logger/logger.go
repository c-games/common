package logger

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/martian/log"
	"github.com/streadway/amqp"
	"gitlab.3ag.xyz/backend/common/fail"
	"gitlab.3ag.xyz/backend/common/mq"
	"gitlab.3ag.xyz/backend/common/mq/msg"
	"gitlab.3ag.xyz/backend/common/timeutil"
	"time"
)

type LogStruct struct {
	Command string
	Data    interface{}
}

type loggerStruct struct {
	channel   chan LogStruct
	mqChannel mq.IChannelAdapter
}

var logger *loggerStruct
var serviceName string
var printRemote bool

func Init(name string, loggerQueueName string, channel mq.IChannelAdapter) {
	if logger == nil {
		serviceName = name

		logger = &loggerStruct{
			channel:   make(chan LogStruct),
			mqChannel: channel,
		}

		_, _ = channel.QueueDeclare(loggerQueueName, true, false, false, false, nil)

		go func() {
			for log := range logger.channel {
				var body []byte
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
				fail.FailedOnError(err, "[logger-gorutine] publish faled")
				logdata, _ := json.Marshal(log.Data)
				fmt.Printf("[%s] %s\n", timeutil.Now(), string(logdata))
			}
		}()
	}
}

func prepareLogMessage(command string, data interface{}) []byte {
	record, err := json.Marshal(data)
	fail.FailedOnError(err, "marshal failed")

	logmsg, err := json.Marshal(msg.CGMessage{
		Command: command,
		Data:    []json.RawMessage{record},
	})

	fail.FailedOnError(err, "marshal failed")
	return logmsg
}

func assertLoggerAvailable() {
	if logger == nil {
		fail.FailedOnError(errors.New("Logger channel is nil "), "Failed to send log")
	}
}

var LOG_TYPE_INFO = 0
var LOG_TYPE_WARN = 1
var LOG_TYPE_DEBUG = 2
var LOG_TYPE_ERROR = 3
var LOG_TYPE_FATAL = 4

func GetLogLevel(num int) string {
	// 0: info 1: warn 2: debug 3: error 4: fatal
	switch num {
	case LOG_TYPE_INFO:
		return "INFO"
	case LOG_TYPE_WARN:
		return "WARN"
	case LOG_TYPE_DEBUG:
		return "DEBUG"
	case LOG_TYPE_ERROR:
		return "ERROR"
	case LOG_TYPE_FATAL:
		return "FATAL"
	default:
		panic("Undefined log level")
	}
}

func _log(serial int64, logLevel int, message string) {
	assertLoggerAvailable()
	if serviceName == "" {
		log.Errorf("ServiceName is empty, Please use logger.Init() setup common/logger")
	}
	timeStr := timeutil.Now()
	log.Infof("[%s] [%v] [%s] [%s] [%s]",
		timeStr, serial, GetLogLevel(logLevel), serviceName, message)
	if printRemote {
		logger.channel <- LogStruct{
			Command: "print",
			Data: msg.PrintRecord{
				Serial:  serial,
				Time:    timeStr,
				Service: serviceName,
				Level:   logLevel,
				Message: message,
			},
		}
	}
}

func Info(serial int64, message string) {
	_log(serial, LOG_TYPE_INFO, message)
}

func Infof(serial int64, message string, params ...interface{}) {
	Info(serial, fmt.Sprintf(message, params))
}

func Warn(serial int64, message string) {
	_log(serial, LOG_TYPE_WARN, message)
}

func Warnf(serial int64, message string, params ...interface{}) {
	Info(serial, fmt.Sprintf(message, params))
}

func Debug(serial int64, message string) {
	_log(serial, LOG_TYPE_DEBUG, message)
}

func Debugf(serial int64, message string, params ...interface{}) {
	Info(serial, fmt.Sprintf(message, params))
}

func Error(serial int64, message string) {
	_log(serial, LOG_TYPE_ERROR, message)
}

func Errorf(serial int64, message string, params ...interface{}) {
	Info(serial, fmt.Sprintf(message, params))
}

func Fatal(serial int64, message string) {
	_log(serial, LOG_TYPE_FATAL, message)
}

func Fatalf(serial int64, message string, params ...interface{}) {
	Info(serial, fmt.Sprintf(message, params))
}