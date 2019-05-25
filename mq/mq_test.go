package mq

import (
	"encoding/json"
	"github.com/streadway/amqp"
	"gitlab.3ag.xyz/backend/common/mq/msg"
	"gitlab.3ag.xyz/backend/common/mq/msg/user"
	"gitlab.3ag.xyz/backend/common/testutil"
	"os"
	"testing"
	"time"
)

const duration = 10 * time.Millisecond

// Channel
func TestAMQPAdapter_GetChannel(t *testing.T) {
	fake := CreateFakeAMQPAdapter()
	fake.GetChannel()
}

// Test the channel struct
// Test Publish and Consume
func TestChannelAdapter_PublishAndConsume(t *testing.T) {
	fakeCh := initFakeChannel()

	srvData, _ := json.Marshal(&user.ValidateData{})
	newMsg := msg.CGMessage{
		Serial: "a-test-serial-number",
		Command: string(user.Validate.GetCommand()),
		ResponseName: msg.User,
		WaitResponse: false,
		Data: []json.RawMessage{srvData},
	}
	err := fakeCh.PublishService(msg.User, newMsg)
	testutil.TestFailIfErr(t, err, "")


	err = fakeCh.PublishServiceNoWaitTo(msg.User, user.Validate, "",
		&user.ValidateData{})
	testutil.TestFailIfErr(t, err, "")

	fakeQ := initFakeQueue(msg.User.QueueName())
	resMsg := fakeQ.Consume("", false, false, false, false, nil)

	timer := time.NewTimer(duration)

	select {
	case d := <-resMsg:
		var resData msg.CGMessage
		_ = json.Unmarshal(d.Body, &resData)
		if resData.Serial != d.AppId {
			t.Error("Unexpect Value")
		} else {
			t.Logf("pass, AppId = data.Serial = %s", d.AppId)
		}

		break;
	case <-timer.C:
		t.Error("time out")
		break
	}

}

func TestChannelAdapter_ExchangeDeclareAndBindQueue_Direct(t *testing.T) {
	fakeCh := initFakeChannel()

	q1, _ := fakeCh.QueueDeclare("queue-1", true, false, false, false, nil)
	q2, _ := fakeCh.QueueDeclare("queue-2", true, false, false, false, nil)
	fakeCh.ExchangeDeclare("exg", "direct", false, false, false, false, nil)

	fakeCh.QueueBindEasy("queue-1", "q1", "exg")
	fakeCh.QueueBindEasy("queue-2", "q2", "exg")


	msg1 := q1.Consume("", false, false, false, false, nil)
	msg2 := q2.Consume("", false, false, false, false, nil)

	_ = fakeCh.Publish("exg", "q1", false, false,
		amqp.Publishing{
			Body: []byte("test"),
		})

	timer := time.NewTimer(duration)
	select {
	case d := <-msg1:
		t.Log("get msg successful: ", d)
	case d := <- msg2:
		t.Error("wrong channel: ", d)
	case <-timer.C:
		t.Error("time out")
	}


	_ = fakeCh.Publish("exg", "q2", false, false,
		amqp.Publishing{
			Body: []byte("test"),
		})
	timer.Reset(duration)
	select {
	case d := <-msg1:
		t.Error("wrong channel: ", d)
	case d := <- msg2:
		t.Log("get msg successful: ", d)
	case <-timer.C:
		t.Error("time out")
	}

	_ = fakeCh.Publish("exg", "q3", false, false,
		amqp.Publishing{
			Body: []byte("test"),
		})
	timer.Reset(duration)
	select {
	case d := <-msg1:
		t.Error("wrong channel: ", d)
	case d := <- msg2:
		t.Error("wrong channel: ", d)
	case <-timer.C:
		t.Log("never receive")
	}

}

func TestChannelAdapter_ExchangeDeclareAndBindQueue_Topic(t *testing.T) {
	fakeCh := initFakeChannel()

	q1, _ := fakeCh.QueueDeclare("queue-1", true, false, false, false, nil)
	q2, _ := fakeCh.QueueDeclare("queue-2", true, false, false, false, nil)
	fakeCh.ExchangeDeclare("exg-topic", "topic", false, false, false, false, nil)

	fakeCh.QueueBindEasy("queue-1", "q1.*", "exg-topic")
	fakeCh.QueueBindEasy("queue-2", "q2.*", "exg-topic")


	msg1 := q1.Consume("", false, false, false, false, nil)
	msg2 := q2.Consume("", false, false, false, false, nil)

	_ = fakeCh.Publish("exg-topic", "q1.abc", false, false,
		amqp.Publishing{
			Body: []byte("test"),
		})

	timer := time.NewTimer(duration)
	select {
	case d := <-msg1:
		t.Log("get msg successful: ", d)
	case d := <- msg2:
		t.Error("wrong channel: ", d)
	case <-timer.C:
		t.Error("time out")
	}


	_ = fakeCh.Publish("exg", "q2.abc", false, false,
		amqp.Publishing{
			Body: []byte("test"),
		})
	timer.Reset(duration)
	select {
	case d := <-msg1:
		t.Error("wrong channel: ", d)
	case d := <- msg2:
		t.Log("get msg successful: ", d)
	case <-timer.C:
		t.Error("time out")
	}

	_ = fakeCh.Publish("exg", "q3.abc", false, false,
		amqp.Publishing{
			Body: []byte("test"),
		})
	timer.Reset(duration)
	select {
	case d := <-msg1:
		t.Error("wrong channel: ", d)
	case d := <- msg2:
		t.Error("wrong channel: ", d)
	case <-timer.C:
		t.Log("never receive")
	}

}

// TODO
// func TestChannelAdapter_ExchangeDeclareAndBindQueue_Fanout(t *testing.T) {}
// func TestChannelAdapter_ExchangeDeclareAndBindQueue_Headers(t *testing.T) {}

// utils
func initFakeChannel() IChannelAdapter {
	// NOTE 實際使用時，會用 GenerateConnect ，Fake 只有在 Test 中使用。
	fakemq := CreateFakeAMQPAdapter()
	chAdp := fakemq.GetChannel()
	return chAdp
}

func initFakeQueue(queueName string) IQueueAdapter {
	chAdp := initFakeChannel()
	return chAdp.GetQueue(queueName, true, false, false, false)
}


// TestMain
func setup() {}
func shutdown() {}
func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	shutdown()
	os.Exit(code)
}
