package service

import (
	kfklib "github.com/opensourceways/kafka-lib/agent"
	"github.com/opensourceways/message-transfer/config"
	"github.com/opensourceways/message-transfer/service/transfer"
)

func SubscribeEurRaw() {
	config.InitEurBuildConfig()
	_ = kfklib.Subscribe(config.EurBuildConfigInstance.Kafka.Group, transfer.Handle, []string{config.EurBuildConfigInstance.Kafka.Topic})
}
