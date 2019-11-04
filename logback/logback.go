package logback

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
	"regexp"
	"runtime"
	"strings"
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
var printRemote = false

func Init(name string, loggerQueueName string, channel mq.IChannelAdapter, isPrintRemote bool, logLevel int) {

	printRemote = isPrintRemote
	log.SetLevel(logLevel)

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
				fail.FailedOnError(err, "[logback-gorutine] publish faled")
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



func _log(serial int64, logLevel int, format string, args ...interface{}) {
	assertLoggerAvailable()
	pc, file, line, ok := runtime.Caller(2)
	details := runtime.FuncForPC(pc)
	var fnName = ""
	if !ok {
		file = "???"
		line = 0
	} else {
		re := regexp.MustCompile(`.*\/gitlab\.3ag\.xyz\/[^\/]*`)
		file = re.ReplaceAllString(file, "")
		p := strings.Split(details.Name(), "/")
		fnName = p[len(p)-1]
	}

	format = "{%v} {%s:%s:%v} " + format
	args = append([]interface{}{serial, file, fnName, line}, args...)

	if serviceName == "" {
		log.Errorf("ServiceName is empty, Please use logback.Init() setup common/logback")
	}
	timeStr := timeutil.Now()

	switch logLevel {
	case log.Info:
		log.Infof(format, args...)
	case log.Error:
		log.Errorf(format, args...)
	case log.Debug:
		log.Debugf(format, args...)
	default:
		panic("Unknown log level")
	}

	if printRemote {
		logger.channel <- LogStruct{
			Command: "print",
			Data: msg.PrintRecord{
				Serial:  serial,
				Time:    timeStr,
				Service: serviceName,
				Level:   logLevel,
				Message: fmt.Sprintf(format, args...),
			},
		}
	}
}

func Infof(serial int64, message string, params ...interface{}) {
	_log(serial, log.Info, fmt.Sprintf(message, params))
}

func Debugf(serial int64, format string, args ...interface{}) {
	_log(serial, log.Debug, format, args...)
}

func Errorf(serial int64, format string, args ...interface{}) {
	_log(serial, log.Error, format, args)
}

// TODO
func Fatalf(serial int64, format string, args ...interface{}) {
	_log(serial, 9999, format, args)
}

// TODO
func SQL() {
	panic("implement me")
}