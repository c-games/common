// Package mq provides wrappers for amqp libraries
package mq

import (
	"encoding/json"
	"github.com/streadway/amqp"
	"gitlab.3ag.xyz/core/backend/common/fail"
	"gitlab.3ag.xyz/core/backend/common/mq/msg"
	"time"
)

// -------------------------------------------
// AMQP Interface Facade
// -------------------------------------------

type IConnection interface {
	// 原本 amqp Connection 的 method
	// TODO 還有沒用到的 Function 沒補
	Channel() (IChannel, error)
}

type IChannel interface {
	// 原本 amqp Channel 的 method
	// TODO  還有沒用到的 Function 沒補
	Publish(exchange, key string, mandatory, immediate bool, msg amqp.Publishing) error
	Consume(queue, consumer string, autoAck, exclusive, noLocal, noWait bool, args amqp.Table) (<-chan amqp.Delivery, error)
	QueueDeclare(name string, durable, autoDelete, exclusive, noWait bool, args amqp.Table) (amqp.Queue, error)
	QOS(count, size int, global bool)
}

// TODO Queue 尚未使用

// -------------------------------------------
// Adapter Interface
// -------------------------------------------

type IAMQPAdapter interface {
	GetChannel() IChannelAdapter
}

type IChannelAdapter interface {
	QueueDeclare(name string, durable, autoDelete, exclusive, noWait bool) (amqp.Queue, error)
	QueueDeclareByQueueConfig(config msg.QueueConfig) (amqp.Queue, error)
	GetQueue(name string, durable, autoDelete, exclusive, noWait bool) IQueueAdapter
	GetQueueWithArgs(name string, durable, autoDelete, exclusive, noWait bool, args amqp.Table) IQueueAdapter
	Publish(targetService msg.Service, cgMsg msg.CGMessage) error
	PublishNoWaitTo(serviceName msg.Service, command msg.ServiceCommand, serial string, data msg.IServiceData) error
	QOS(count, size int, global bool)
	// TODO 需要一個直接指定 queue 的 publish
	Close()
}

type IQueueAdapter interface {
	Consume(consumer string, autoAck, exclusive, noLocal, noWait bool, args amqp.Table) <-chan amqp.Delivery
}

// -------------------------------------------
// Adapter Structure
// -------------------------------------------

type AMQPAdapter struct {
	Connect IConnection
}

type ChannelAdapter struct {
	AMQPAdapter *AMQPAdapter
	Channel IChannel
}

type QueueAdapter struct {
	name string
	ChannelAdapter *ChannelAdapter
	Queue *amqp.Queue
}


// -------------------------------------------
// AMQPAdapter
// -------------------------------------------

func (adp *AMQPAdapter) GetChannel() IChannelAdapter {
	ch, err := adp.Connect.Channel()
	fail.FailOnError(err, "")
	return &ChannelAdapter{
		AMQPAdapter: adp,
		Channel: ch,
	}
}

// -------------------------------------------
// ChannelAdapter
// -------------------------------------------

func (adp *ChannelAdapter) Publish(targetService msg.Service, cgmsg msg.CGMessage) error {
	if (cgmsg.WaitResponse && &cgmsg.ResponseName == nil) ||
		(cgmsg.WaitResponse && cgmsg.Serial == "") {
		return NoResponseErr{}
	}

	appId := cgmsg.Serial
	data, err := json.Marshal(cgmsg)
	fail.FailOnError(err, "parse cg message error")
	err = adp.Channel.Publish(
		targetService.QueueName(),
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        data,
			Timestamp:   time.Now(),
			AppId:       appId,
		})

	return err
}

type NoResponseErr struct {}

func (err NoResponseErr) Error() string {
	return "Empty Response Configuration"
}

func (adp *ChannelAdapter) PublishNoWaitTo(service msg.Service, command msg.ServiceCommand,
	serial string, data msg.IServiceData) error {
		err := adp.Channel.Publish(
		service.QueueName(),
		"", // route key
		false, //
		false, //
		amqp.Publishing{
			ContentType: "application/json",
			Body:        msg.ToByteArray(data),
			Timestamp:   time.Now(),
			AppId: serial,
		})
		return err
}

func (adp *ChannelAdapter) QueueDeclare(name string, durable, autoDelete, exclusive, noWait bool) (amqp.Queue, error) {
	q, err := adp.Channel.QueueDeclare(name, durable, autoDelete, exclusive, noWait, nil)
	return q, err
}

func (adp *ChannelAdapter) QueueDeclareByQueueConfig(config msg.QueueConfig) (amqp.Queue, error) {
	name, durable, autoDelete, exclusive, noWait, args := config.Spread()
	q, err := adp.Channel.QueueDeclare(name, durable, autoDelete, exclusive, noWait, args)
	return q, err
}

func (adp *ChannelAdapter) GetQueue(name string, durable, autoDelete, exclusive, noWait bool) IQueueAdapter {
	return adp.GetQueueWithArgs(name, durable, autoDelete, exclusive, noWait, nil)
}

func (adp *ChannelAdapter) GetQueueWithArgs(name string, durable, autoDelete, exclusive, noWait bool, args amqp.Table) IQueueAdapter {
	q, err := adp.Channel.QueueDeclare(
		name,    // name
		durable, // durable
		autoDelete, // delete when usused
		exclusive,  // exclusive
		noWait, // no-wait
		nil,   // arguments
	)

	fail.FailOnError(err, "Failed to declare a queue")

	return &QueueAdapter{
		name: name,
		ChannelAdapter: adp,
		Queue: &q,
	}

}

func (adp *ChannelAdapter) QOS(count, size int, global bool) {
	adp.Channel.QOS(count, size, global)
}

func (adp *ChannelAdapter) Close() {
	panic("implement me")
}


// -------------------------------------------
// QueueAdapter
// -------------------------------------------

func (chAdp *QueueAdapter) Consume(consumer string, autoAck, exclusive, noLocal,
	noWait bool, args amqp.Table) <-chan amqp.Delivery {
	// real version
	deliver, err := chAdp.ChannelAdapter.Channel.Consume(chAdp.name, consumer, autoAck, exclusive, noLocal, noWait, args)
	fail.FailOnError(err, "Consume failed")
	return deliver

}

// -------------------------------------------
// Util Function
// -------------------------------------------

func GenerateConnect(url string) *AMQPAdapter {
	return &AMQPAdapter{}
}
