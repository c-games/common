// Package mq provides wrappers for amqp libraries
package mq

import (
	"crypto/tls"
	"fmt"
	"github.com/streadway/amqp"
	"gitlab.3ag.xyz/backend/common/fail"
	"os"
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
	ConnectionState() tls.ConnectionState
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

	GetQueue(name string, durable, autoDelete, exclusive, noWait bool) IQueueAdapter
	GetQueueWithArgs(name string, durable, autoDelete, exclusive, noWait bool, args amqp.Table) IQueueAdapter
	Publish(exchange, key string, mandatory, immediate bool, msg amqp.Publishing) error
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
	if err != nil{
		fmt.Println(err)
		os.Exit(2)
	}
	fail.FailedOnError(err, "get channel failed")
	return &ChannelAdapter{
		AMQPAdapter: adp,
		Channel: ch,
	}
}

func (adp *AMQPAdapter) Close() {
	//err := adp.Channel.Qos(count, size, global)
	err := adp.Connect.Close()
	fail.FailedOnError(err, "Connect close failed")
}
func (adp *AMQPAdapter) ConnectionState() tls.ConnectionState {
	return adp.Connect.ConnectionState()
}

func (adp *ChannelAdapter) QOS(count, size int, global bool) {
	err := adp.Channel.Qos(count, size, global)
	fail.FailedOnError(err, "QOS setup failed")
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

func (adp *ChannelAdapter) QueueDeclare(name string, durable, autoDelete, exclusive, noWait bool, args amqp.Table) (IQueueAdapter, error) {
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

	fail.FailedOnError(err, "Failed to declare a queue")

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
	fail.FailedOnError(err, "Failed to declare a exchange")
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
	//fail.FailedOnError(err, "close channel failed")
}


// -------------------------------------------
// QueueAdapter
// -------------------------------------------

func (qAdp *QueueAdapter) Consume(consumer string, autoAck, exclusive, noLocal, noWait bool, args amqp.Table) <-chan amqp.Delivery {
	deliver, err := qAdp.ChannelAdapter.Channel.Consume(qAdp.name, consumer, autoAck, exclusive, noLocal, noWait, args)
	fail.FailedOnError(err, fmt.Sprintf("Consume failed: queue: %s", qAdp.name))
	return deliver
}

// -------------------------------------------
// Util Function
// -------------------------------------------

func GenerateConnect(url string) *AMQPAdapter {
	conn, err := amqp.Dial(url)
	fail.FailedOnError(err, "Connect to RabbitMq failed")

	return &AMQPAdapter{
		Connect: conn,
	}
}