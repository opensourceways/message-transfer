package service

import (
	"github.com/opensourceways/message-transfer/common/kafka"
	"github.com/opensourceways/message-transfer/config"
	"github.com/opensourceways/message-transfer/service/transfer"
)

func SubscribeEurRaw() {
	config.InitEurBuildConfig()
	h := transfer.EurGroupHandler{}
	kafka.ConsumeGroup(config.EurBuildConfigInstance.Consume, &h)
}
