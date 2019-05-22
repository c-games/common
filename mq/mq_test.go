package mq

import (
	"encoding/json"
	"gitlab.3ag.xyz/core/backend/common/mq/msg"
	"gitlab.3ag.xyz/core/backend/common/mq/msg/user"
	"gitlab.3ag.xyz/core/backend/common/testutil"
	"os"
	"testing"
	"time"
)

const duration = 1 * time.Second

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
	err := fakeCh.Publish(msg.User, newMsg)
	testutil.TestFailIfErr(err, t)


	err = fakeCh.PublishNoWaitTo(msg.User, user.Validate, "",
		&user.ValidateData{})
	testutil.TestFailIfErr(err, t)

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
