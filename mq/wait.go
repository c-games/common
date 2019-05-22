package mq

import (
	"github.com/streadway/amqp"
	"time"
)

type deliveryHandler func(delivery amqp.Delivery) interface {}

type TimeoutErr struct {}

func (err TimeoutErr) Error() string {
	return "Timeout"
}

type Hold struct {
	holdFn func() (interface{}, error)
}

func WaitChannel(msg <-chan amqp.Delivery, timeout int64, handleFn deliveryHandler) (interface{}, error) {
	duration := time.Duration(timeout) * time.Second
	timer := time.NewTicker(duration)
	select {
	case d := <-msg:
		return handleFn(d), nil
	case <- timer.C:
		return nil, TimeoutErr{}
	}
}


func Wait(msg <-chan amqp.Delivery, timeout int64, handleFn deliveryHandler) *Hold {
	return &Hold{
		holdFn: func() (interface{}, error) {
			return WaitChannel(msg, timeout, handleFn)
		}}
}

func (h *Hold) RunIfErr(errFn func()) interface {} {
	rlt, err := h.holdFn()
	if err != nil {
		errFn()
		return nil
	} else {
		return rlt
	}

}
