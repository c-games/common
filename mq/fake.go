package mq

import (
	"fmt"
	"github.com/streadway/amqp"
	"gitlab.3ag.xyz/backend/common/fail"
	"regexp"
)

// TODO Not Thread Safe
var fakeAdapterSingleton *AMQPAdapter
var allChannel map[string]chan amqp.Delivery
var fakeChanelSinglton *FakeChannel

func CreateFakeAMQPAdapter() IAMQPAdapter {
	if fakeAdapterSingleton == nil{
		fakeAdapterSingleton = &AMQPAdapter{
			Connect: &FakeConnection{},
		}
	}

	return fakeAdapterSingleton
}

func CreateFakeChannel() IChannelAdapter {

	fc := FakeConnection{}
	_, _ = fc.Channel()
	ch := GetFakeChannel()

	return &ChannelAdapter{
		AMQPAdapter: &AMQPAdapter{Connect: &fc},
		Channel: ch,
	}
}

type FakeAMQPAdapter struct {
	AMQPAdapter
}

type FakeConnection struct {}

type ExchangeStruct struct {
	exchangeType string
	channelMatchCheck map[string] func(queueName string) bool
}

type FakeChannel struct {
	*amqp.Channel
	channels map[string]chan amqp.Delivery
	exchange map[string]*ExchangeStruct
}

// -------------------------------------------
// FakeConnection
// -------------------------------------------
func (fake *FakeConnection) Channel() (*amqp.Channel, error) {
	if allChannel == nil {
		allChannel = make(map[string]chan amqp.Delivery)
	}

	fakeChanelSinglton = &FakeChannel{
		channels: allChannel,
		exchange: make(map[string] *ExchangeStruct),
	}
	return &amqp.Channel{}, nil
}

func (fake *FakeConnection) Close() error {
	// NOTE do nothing
	return nil
}

func GetFakeChannel() *FakeChannel {
	return fakeChanelSinglton
}

// -------------------------------------------
// FakeChannel
// -------------------------------------------

func (fake *FakeChannel) Publish(exchange, key string, mandatory, immediate bool, msg amqp.Publishing) error {
	targetQueue := "unknow"
	if key == "" && exchange == "" {
		targetQueue = exchange
	} else if exchange == "" {
		targetQueue = key
	} else {
		exg := fake.exchange[exchange]
		if exg == nil {
			panic("not exist exchange " + exchange)
		}
		for qName, chkFn := range exg.channelMatchCheck {
			if chkFn(key) {
				targetQueue = qName
			}
		}
	}

	fake.checkQueueExist(targetQueue)

	go func() {
		fake.channels[targetQueue]<- amqp.Delivery{
			Body: msg.Body,
			AppId: msg.AppId,
			Timestamp: msg.Timestamp,
			ContentType: msg.ContentType,
		}
	}()
	return nil
}

func (fake *FakeChannel) QueueDeclare(name string, durable, autoDelete, exclusive, noWait bool, args amqp.Table) (amqp.Queue, error) {
	return amqp.Queue{
		Name: name,
	}, nil
}

func (fake *FakeChannel) ExchangeDeclare(name, kind string, durable, autoDelete, internal, noWait bool, args amqp.Table) error {
	fake.exchange[name] = &ExchangeStruct{
		exchangeType: kind,
		channelMatchCheck: make(map[string]func(queueName string) bool),
	}

	return nil
}

func (fake *FakeChannel) checkQueueExist(name string) {
	if fake.channels[name] == nil {
		fake.channels[name] = make(chan amqp.Delivery)
	}
}

func (fake *FakeChannel) QueueBind(name, key, exchange string, noWait bool, args amqp.Table) error {
	exg := fake.exchange[exchange]

	exg.channelMatchCheck[name] = func(pubKey string) bool {
		fake.checkQueueExist(pubKey)
		if exg.exchangeType == "direct" {
			return pubKey == key
		} else if exg.exchangeType == "topic" {
			isMatch, _ := regexp.MatchString(key, pubKey)
			return isMatch

		} else if exg.exchangeType == "fanout" {
			// TODO
		} else if exg.exchangeType == "headers" {
			// TODO
		} else {
			fail.FailedOnError(amqp.Error{
				Reason: fmt.Sprintf("type %s unexpect", exg.exchangeType),
				Code: 999,
			},
			"Unexpect exchange type")
		}
		return false
	}

	return nil
}

func (fake *FakeChannel) Consume(queue, consumer string, autoAck, exclusive, noLocal, noWait bool, args amqp.Table) (<-chan amqp.Delivery, error) {
	ch := fake.channels[queue]

	if ch == nil {
		fake.channels[queue] = make(chan amqp.Delivery)
		ch = fake.channels[queue]
	}

	return ch, nil

}

func (fake *FakeChannel) Qos(count, size int, global bool) error {
	// NOTE do nothing
	return nil
}
