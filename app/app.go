package app

import (
	"github.com/spf13/viper"
	"gitlab.3ag.xyz/core/backend/common/mq"
	"gitlab.3ag.xyz/core/backend/common/db"
	"gitlab.3ag.xyz/core/backend/common/logger"
)

func Init() (*mq.CgAmqp, chan bool) {

	// RabbitMQ part
	// ----------------------------------------
	rabbitMqConf := viper.GetString("mq")

	// NOTE 當 order 一開始後，subscribe 一個 channel 是收 order 處理的
	mqCtx := mq.GenConn(rabbitMqConf)

	mqCtx.GenChannel()
	mqCtx.GenQueue("cg-wallet", true, false, false, false, nil)
	mqCtx.GenQueue("cg-user", true, false, false, false, nil)
	mqCtx.GenQueue("cg-orders", true, false, false, false, nil)
	mqCtx.GenQueue("cg-wallet-pipeline-response", true, false, false, false, nil)
	mqCtx.GenQueue("cg-user-pipeline-response", true, false, false, false, nil)
	mqCtx.GenQueue("cg-orders-pipeline-response", true, false, false, false, nil)
	// TODO other Queue
	mqCtx.QOS(1, 0, false)

	// NOTE generate a logger channel
	chLogger := mqCtx.GetChannel()
	logger.Init(chLogger)


	dbConf := viper.GetString("db")
	db.Init(dbConf)

	forever := make(chan bool)

	return &mqCtx, forever

}
