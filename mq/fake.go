package mq

import (
	"github.com/streadway/amqp"
)

// TODO Not Thread Safe
var fakeAdapterSingleton *AMQPAdapter
var allChannel map[string]chan amqp.Delivery


func CreateFakeAMQPAdapter() IAMQPAdapter {
	if fakeAdapterSingleton == nil{
		fakeAdapterSingleton = &AMQPAdapter{
			Connect: &FakeConnection{},
		}
	}

	return fakeAdapterSingleton
}

type FakeAMQPAdapter struct {
	AMQPAdapter
}

type FakeConnection struct {}

type FakeChannel struct {
	channels map[string]chan amqp.Delivery
}

// -------------------------------------------
// FakeConnection
// -------------------------------------------

func (fake *FakeConnection) Channel() (IChannel, error) {
	if allChannel == nil {
		allChannel = make(map[string]chan amqp.Delivery)
	}

	return &FakeChannel{
		channels: allChannel,
	}, nil
}

// -------------------------------------------
// FakeChannel
// -------------------------------------------

func (fake *FakeChannel) Publish(exchange, key string, mandatory, immediate bool, msg amqp.Publishing) error {
	if fake.channels[exchange] == nil {
		fake.channels[exchange] = make(chan amqp.Delivery)
	}

	go func() {
		fake.channels[exchange]<- amqp.Delivery{
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

func (fake *FakeChannel) Consume(queue, consumer string, autoAck, exclusive, noLocal, noWait bool, args amqp.Table) (<-chan amqp.Delivery, error) {
	ch := fake.channels[queue]

	if ch == nil {
		fake.channels[queue] = make(chan amqp.Delivery)
		ch = fake.channels[queue]
	}

	return ch, nil

}

func (fake *FakeChannel) QOS(count, size int, global bool) {
	// NOTE do nothing
}
