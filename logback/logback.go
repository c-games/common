package logback

import (
	"fmt"
	"github.com/google/martian/log"
	"regexp"
	"runtime"
	"strings"
)

type LogStruct struct {
	Command string
	Data    interface{}
}

func Init(logLevel int) {
	log.SetLevel(logLevel)
}

func _log(serial int, logLevel int, format string, args ...interface{}) {
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
}

func Infof(message string, params ...interface{}) {
	_log(0, log.Info, fmt.Sprintf(message, params...))
}

func Debugf(format string, args ...interface{}) {
	_log(0, log.Debug, format, args...)
}

func Errorf(format string, args ...interface{}) {
	_log(0, log.Error, format, args)
}

func Infof_s(serial int, message string, params ...interface{}) {
	_log(serial, log.Info, fmt.Sprintf(message, params...))
}

func Debugf_s(serial int, format string, args ...interface{}) {
	_log(serial, log.Debug, format, args...)
}

func Errorf_s(serial int, format string, args ...interface{}) {
	_log(serial, log.Error, format, args)
}
