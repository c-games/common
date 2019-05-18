package mq

import (
	"github.com/streadway/amqp"
	"gitlab.3ag.xyz/core/backend/common/fail"
)

// 這個設計不太對，因為一個 connect 可以多 channel
// 這個 scruct 記 channel 沒什麼意義
type CgAmqp struct {
	Conn *amqp.Connection
	Ch *amqp.Channel
	Queues map[string]*amqp.Queue // TODO remove
}

// Message Queue Channel
type MqChann struct {
	Ch *amqp.Channel
}

type ICgAmqp interface {
	GenChannel()
	GenRoute(string)
	GenQueue(string)
	GenBindedQueue(string)
	QOS()
	CloseAll()
}

func GenConn(url string) CgAmqp {
	conn, err := amqp.Dial(url)
	fail.FailOnError(err, "Failed to connect to RabbitMQ")
	// return CgAmqp{con, make(map[string]*amqp.Channel)}
	return CgAmqp{Conn: conn, Queues: make(map[string]*amqp.Queue)}
}


func (ctx *CgAmqp) GenChannel() {
	ch, err := ctx.Conn.Channel()
	fail.FailOnError(err, "Failed to open a channel")
	ctx.Ch = ch
}

func (ctx *CgAmqp) CloseAll() {
	_ = ctx.Ch.Close()
	_ = ctx.Conn.Close()
}

func (ctx *CgAmqp) GenExchange(name, kind string, durable, autoDelete, internal, noWait bool, args amqp.Table) {
	if ctx.Ch == nil {
		fail.FailOnError(nil, "Create a Route before the Channel")
	}
	err := ctx.Ch.ExchangeDeclare(name, kind, durable, autoDelete, internal, noWait, args)

	fail.FailOnError(err, "Failed to declare an exchange")
}

func (ctx *CgAmqp) GenRoute(exgName string) {
	ctx.GenExchange(exgName, "direct", true, false, false, false, nil)
}


func (ctx *CgAmqp) GenQueue(name string, durable, autoDelete, exclusive, noWait bool, args amqp.Table) {
	if ctx.Ch == nil {
		fail.FailOnError(nil, "Create a Queue before the Channel")
	}

	q, err := ctx.Ch.QueueDeclare(
		name,    // name
		durable, // durable
		autoDelete, // delete when usused
		exclusive,  // exclusive
		noWait, // no-wait
		args,   // arguments
	)

	fail.FailOnError(err, "Failed to declare a queue")

	ctx.Queues[name] = &q
}

// NOTE 用來給不存 ctx 用的 queue (因為是無名 queue)
// TODO 和 GenQueue 覆的要再收成一個 function
func (ctx *CgAmqp) GenRandomQueue(durable, autoDelete, exclusive, noWait bool, args amqp.Table) amqp.Queue {
	if ctx.Ch == nil {
		fail.FailOnError(nil, "Create a Queue before the Channel")
	}

	q, err := ctx.Ch.QueueDeclare(
		"",    // name
		durable, // durable
		autoDelete, // delete when usused
		exclusive,  // exclusive
		noWait, // no-wait
		args,   // arguments
	)

	fail.FailOnError(err, "Failed to declare a queue")

	return q
}

func (ctx *CgAmqp) BindQueue(queueName, bindKey, exgName string) {
	err := ctx.Ch.QueueBind(
		queueName, // queue name
		bindKey, // routing key
		exgName, // exchange
		false,
		nil,
	)

	fail.FailOnError(err, "Failed to bind a queue")
}

// TODO deprecated
func (ctx *CgAmqp) GenBindedQueue(qName, bindKey, exgName string) {
	var autoDelete bool
	// NOTE 如果 qName 是空白的時，rabbitmq 會用一個 random name，因此這個 random queue 要刪
	if qName == "" {
		autoDelete = true
	} else {
		autoDelete = false
	}
	ctx.GenQueue(qName, true, autoDelete, true, false, nil)
	ctx.BindQueue(qName, bindKey, exgName)
}

func (ctx *CgAmqp) QOS(count, size int, global bool) {
	ctx.Ch.Qos(count, size, global)
}

func (ctx *CgAmqp) Consume(queueName, comsumer string, autoAck, exclusive, noLocal, noWait bool, args amqp.Table) <-chan amqp.Delivery {
		msgs, err := ctx.Ch.Consume(
		queueName, // queue
		comsumer,     // consumer
		autoAck,   // auto ack
		exclusive,  // exclusive
		noLocal,  // no local
		noWait,  // no wait
		args,    // args
	)
	fail.FailOnError(err, "Failed to register a consumer")

	return msgs
}

//
func (ctx *CgAmqp) GetChannel() *MqChann {
	ch, err := ctx.Conn.Channel()
	fail.FailOnError(err, "Failed to open a channel")
	return &MqChann{ Ch: ch }
}

// --------------
//
//
func (chann *MqChann) GenQueue(name string, durable, autoDelete, exclusive, noWait bool, args amqp.Table) *amqp.Queue {
	q, err := chann.Ch.QueueDeclare(
		name,    // name
		durable, // durable
		autoDelete, // delete when usused
		exclusive,  // exclusive
		noWait, // no-wait
		args,   // arguments
	)

	fail.FailOnError(err, "Failed to declare a queue")

	return &q
}

func (ch *MqChann) Consume(queueName, comsumer string, autoAck, exclusive, noLocal, noWait bool, args amqp.Table) <-chan amqp.Delivery {
	msgs, err := ch.Ch.Consume(
		queueName, // queue
		comsumer,     // consumer
		autoAck,   // auto ack
		exclusive,  // exclusive
		noLocal,  // no local
		noWait,  // no wait
		args,    // args
	)
	fail.FailOnError(err, "Failed to register a consumer")

	return msgs
}

func (ch *MqChann) Close() {
	defer ch.Close()
}
