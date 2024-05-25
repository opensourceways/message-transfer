package service

import (
	"github.com/opensourceways/message-transfer/common/kafka"
	"github.com/opensourceways/message-transfer/service/transfer"
	"github.com/opensourceways/message-transfer/utils"
	"github.com/sirupsen/logrus"
)

func SubscribeEurRaw() {
	cfg := new(kafka.ConsumeConfig)
	if err := utils.LoadFromYaml("config/eur_build_conf.yaml", cfg); err != nil {
		logrus.Error("Config初始化失败, err:", err)
		return
	}

	h := transfer.EurGroupHandler{}
	kafka.ConsumeGroup(*cfg, &h)
}
