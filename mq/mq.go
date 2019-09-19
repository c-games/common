// Package mq provides wrappers for amqp libraries
package mq

import (
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"gitlab.3ag.xyz/backend/common/fail"
	"gitlab.3ag.xyz/backend/common/mq/msg"
	"time"
)

// -------------------------------------------
// Error type
// -------------------------------------------

type NoResponseErr struct {}

func (err NoResponseErr) Error() string {
	return "Empty Response Configuration"
}

// -------------------------------------------
// AMQP Interface Adapter
// -------------------------------------------

type IConnection interface {
	Channel() (*amqp.Channel, error)
	Close() error
}

type IChannel interface {
	// 原本 amqp Channel 的 method
	// TODO  還有沒用到的 Function 沒補
	Publish(exchange, key string, mandatory, immediate bool, msg amqp.Publishing) error
	Consume(queue, consumer string, autoAck, exclusive, noLocal, noWait bool, args amqp.Table) (<-chan amqp.Delivery, error)
	QueueDeclare(name string, durable, autoDelete, exclusive, noWait bool, args amqp.Table) (amqp.Queue, error)
	ExchangeDeclare(name, kind string, durable, autoDelete, internal, noWait bool, args amqp.Table) error
	QueueBind(name, key, exchange string, noWait bool, args amqp.Table) error
	Qos(prefetchCount, prefetchSize int, global bool) error
	Close() error
}

// TODO Queue 尚未使用

// -------------------------------------------
// Adapter Interface
// -------------------------------------------

type IAMQPAdapter interface {
	GetChannel() IChannelAdapter
}

type IChannelAdapter interface {
	QueueDeclare(name string, durable, autoDelete, exclusive, noWait bool, args amqp.Table) (IQueueAdapter, error)
	QueueDeclareByQueueConfig(config msg.QueueConfig) (IQueueAdapter, error)
	GetQueue(name string, durable, autoDelete, exclusive, noWait bool) IQueueAdapter
	GetQueueWithArgs(name string, durable, autoDelete, exclusive, noWait bool, args amqp.Table) IQueueAdapter
	Publish(exchange, key string, mandatory, immediate bool, msg amqp.Publishing) error
	PublishService(targetService msg.Service, cgMsg msg.CGMessage) error
	PublishServiceNoWaitTo(serviceName msg.Service, command msg.ServiceCommand, serial string, data msg.IServiceData) error
	ResponseService(targetService msg.Service, cgMsg msg.CGMessage) error
	ResponseServiceNoWaitTo(serviceName msg.Service, command msg.ServiceCommand, serial string, data msg.IServiceData) error
	ExchangeDeclare(name, kind string, durable, autoDelete, internal, noWait bool, args amqp.Table)
	QueueBind(queueName, bindKey, exchangeName string, noWait bool, args amqp.Table)
	QueueBindEasy(queueName, bindKey, exchangeName string)
	QOS(count, size int, global bool)
	Consume(queue, consumer string, autoAck, exclusive, noLocal, noWait bool, args amqp.Table) (<-chan amqp.Delivery, error)
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
	fail.FailOnError(err, "get channel failed")
	return &ChannelAdapter{
		AMQPAdapter: adp,
		Channel: ch,
	}
}

func (adp *AMQPAdapter) Close() {
	//err := adp.Channel.Qos(count, size, global)
	err := adp.Connect.Close()
	fail.FailOnError(err, "Connect close failed")
}

func (adp *ChannelAdapter) QOS(count, size int, global bool) {
	err := adp.Channel.Qos(count, size, global)
	fail.FailOnError(err, "QOS setup failed")
}

// -------------------------------------------
// ChannelAdapter
// -------------------------------------------

// NOTE low level api
func (adp *ChannelAdapter) Publish(exchange, key string, mandatory, immediate bool, msg amqp.Publishing) error {
	err := adp.Channel.Publish(
		exchange,
		key,
		mandatory,
		immediate,
		msg)

	return err
}

func (adp *ChannelAdapter) PublishService(targetService msg.Service, cgmsg msg.CGMessage) error {
	if (cgmsg.WaitResponse && &cgmsg.ResponseQueue == nil) ||
		(cgmsg.WaitResponse) { //&& cgmsg.Serial == nil) {
		return NoResponseErr{}
	}

	// appId := cgmsg.Serial
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
			// AppId:       appId,   // NOTE appid 是 amqp 的 serial，目前還沒用到
		})

	return err
}

func (adp *ChannelAdapter) PublishServiceNoWaitTo(service msg.Service, command msg.ServiceCommand,
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

func (adp *ChannelAdapter) ResponseService(targetService msg.Service, cgmsg msg.CGMessage) error {
	if (cgmsg.WaitResponse && &cgmsg.ResponseQueue == nil) ||
		(cgmsg.WaitResponse) { // && cgmsg.Serial == nil) {
		return NoResponseErr{}
	}

	// appId := cgmsg.Serial
	data, err := json.Marshal(cgmsg)
	fail.FailOnError(err, "parse cg message error")
	err = adp.Channel.Publish(
		targetService.ResponseQueueName(),
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        data,
			Timestamp:   time.Now(),
			// AppId:       appId, // NOTE amqp 用的 serial，目前沒用到
		})

	return err
}


func (adp *ChannelAdapter) ResponseServiceNoWaitTo(service msg.Service, command msg.ServiceCommand,
	serial string, data msg.IServiceData) error {
		err := adp.Channel.Publish(
		service.ResponseQueueName(),
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

func (adp *ChannelAdapter) QueueDeclare(name string, durable, autoDelete, exclusive, noWait bool, args amqp.Table) (IQueueAdapter, error) {
	q, err := adp.Channel.QueueDeclare(name, durable, autoDelete, exclusive, noWait, args)
	return &QueueAdapter{Queue: &q, ChannelAdapter: adp, name: name}, err
}

func (adp *ChannelAdapter) QueueDeclareByQueueConfig(config msg.QueueConfig) (IQueueAdapter, error) {
	name, durable, autoDelete, exclusive, noWait, args := config.Spread()
	q, err := adp.Channel.QueueDeclare(name, durable, autoDelete, exclusive, noWait, args)
	return &QueueAdapter{Queue: &q, ChannelAdapter: adp, name: name}, err
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

func (adp *ChannelAdapter) ExchangeDeclare(name, kind string, durable, autoDelete, internal, noWait bool, args amqp.Table) {
	err := adp.Channel.ExchangeDeclare(
		name,    // name
		kind,    // kind
		durable, // durable
		autoDelete, // delete when usused
		internal,  // exclusive
		noWait, // no-wait
		nil,   // arguments
	)
	fail.FailOnError(err, "Failed to declare a exchange")
}

func (adp *ChannelAdapter) QueueBind(queueName, bindKey, exchangeName string, noWait bool, args amqp.Table) {
	adp.Channel.QueueBind(queueName, bindKey, exchangeName, false, nil)
}

func (adp *ChannelAdapter) QueueBindEasy(queueName, bindKey, exchangeName string) {
	adp.QueueBind(queueName, bindKey, exchangeName, false, nil)
}

func (adp *ChannelAdapter) Consume(queue, consumer string, autoAck, exclusive, noLocal, noWait bool, args amqp.Table) (<-chan amqp.Delivery, error) {
	return adp.Channel.Consume(queue, consumer, autoAck, exclusive, noLocal, noWait, nil)
}

func (adp *ChannelAdapter) Close() {
	// TODO fix bug
	// close channel failed: Exception (504) Reason: "channel/connection is not open"
	_ = adp.Channel.Close()
	//fail.FailOnError(err, "close channel failed")
}


// -------------------------------------------
// QueueAdapter
// -------------------------------------------

func (qAdp *QueueAdapter) Consume(consumer string, autoAck, exclusive, noLocal, noWait bool, args amqp.Table) <-chan amqp.Delivery {
	deliver, err := qAdp.ChannelAdapter.Channel.Consume(qAdp.name, consumer, autoAck, exclusive, noLocal, noWait, args)
	fail.FailOnError(err, fmt.Sprintf("Consume failed: queue: %s", qAdp.name))
	return deliver
}

// -------------------------------------------
// Util Function
// -------------------------------------------

func GenerateConnect(url string) *AMQPAdapter {
	conn, err := amqp.Dial(url)
	fail.FailOnError(err, "Connect to RabbitMq failed")

	return &AMQPAdapter{
		Connect: conn,
	}
}