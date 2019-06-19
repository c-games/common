package logger

import (
	"encoding/json"
	"github.com/streadway/amqp"
	"gitlab.3ag.xyz/backend/common/mq"
	"gitlab.3ag.xyz/backend/common/mq/msg"
	"os"
	"testing"
)

var fakeChannel mq.IChannelAdapter
var fakeQueue mq.IQueueAdapter

func TestLogf(t *testing.T) {
	sendData := "send to logger"
	Log(sendData)

	mqMsg := fakeQueue.Consume("", false, false, false, false, nil)

	mq.Wait(mqMsg, 1, func(d amqp.Delivery) interface {} {
		var cgData msg.CGMessage
		_ = json.Unmarshal(d.Body, &cgData)

		var loggerData LogMessage
		_ = json.Unmarshal(cgData.Data[0], &loggerData)
		if loggerData.Message != sendData {
			t.Logf("Incorrect data\nResult:%s\nExpect:%s", cgData.Data, sendData)
			t.Fail()
		}
		return nil
	}).RunIfErr(func() {
		t.Logf("timeout")
		t.Fail()
	})
}

// TestMain
func setup() {
	fakeConnect := mq.CreateFakeAMQPAdapter()
	fakeChannel = fakeConnect.GetChannel()
	fakeQueue = fakeChannel.GetQueue(msg.Logger.QueueName(), true, false, false, false)
	Init("unit test", msg.Logger.QueueName(), fakeChannel)
}
func shutdown() {}
func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	shutdown()
	os.Exit(code)
}
