package app

import (
	"github.com/spf13/viper"
	"gitlab.3ag.xyz/backend/service"
)

type ProxyData struct {
	ReqData *service.CGMessage
	ResChan chan []byte
	Target string // target queue
}

func CheckEnvSetting(keys []string) {
	for _, key := range keys {
		if !viper.IsSet(key) {
			panic("unset environment = " + key)
		}
	}
}
