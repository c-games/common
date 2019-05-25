package app

import (
	"github.com/spf13/viper"
	"gitlab.3ag.xyz/backend/common/db"
	"gitlab.3ag.xyz/backend/common/logger"
	"gitlab.3ag.xyz/backend/common/mq"
	"gitlab.3ag.xyz/backend/common/mq/msg"
)

func Init() (*mq.AMQPAdapter, mq.IChannelAdapter, chan bool) {

	// RabbitMQ part
	// ----------------------------------------
	rabbitMqConf := viper.GetString("mq")

	// NOTE 當 order 一開始後，subscribe 一個 channel 是收 order 處理的
	mqAdp := mq.GenerateConnect(rabbitMqConf)
	// mqCtx := mq.GenConn(rabbitMqConf)

	mqChAdp := mqAdp.GetChannel()

	_, _ = mqChAdp.QueueDeclareByQueueConfig(msg.Wallet.GetQueueConfig())
	_, _ = mqChAdp.QueueDeclareByQueueConfig(msg.User.GetQueueConfig())
	_, _ = mqChAdp.QueueDeclareByQueueConfig(msg.Orders.GetQueueConfig())

	_, _ = mqChAdp.QueueDeclareByQueueConfig(msg.Wallet.GetResponseQueueConfig())
	_, _ = mqChAdp.QueueDeclareByQueueConfig(msg.User.GetResponseQueueConfig())
	_, _ = mqChAdp.QueueDeclareByQueueConfig(msg.Orders.GetResponseQueueConfig())


	exgName := "cg-exchanger" // TODO move

	mqChAdp.ExchangeDeclare(exgName, "route", true, false, false, false, nil)

	mqChAdp.QueueBindEasy(msg.Wallet.QueueName(), msg.Wallet.GetQueueBind(), exgName)
	mqChAdp.QueueBindEasy(msg.User.QueueName(), msg.User.GetQueueBind(), exgName)
	mqChAdp.QueueBindEasy(msg.Orders.QueueName(), msg.Orders.GetQueueBind(), exgName)

	mqChAdp.QueueBindEasy(msg.Wallet.ResponseQueueName(), msg.Wallet.GetResponseQueueBind(), exgName)
	mqChAdp.QueueBindEasy(msg.User.ResponseQueueName(), msg.User.GetResponseQueueBind(), exgName)
	mqChAdp.QueueBindEasy(msg.Orders.ResponseQueueName(), msg.Orders.GetResponseQueueBind(), exgName)


	mqChAdp.QOS(1, 0, false)

	// NOTE generate a logger channel
	chLogger := mqAdp.GetChannel()
	logger.Init(msg.Logger.QueueName(), chLogger)


	dbConf := viper.GetString("db")
	db.Init(dbConf)

	forever := make(chan bool)

	return mqAdp, mqChAdp, forever
}
