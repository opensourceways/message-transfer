package config

import (
	"github.com/opensourceways/message-transfer/common/kafka"
	"github.com/opensourceways/message-transfer/utils"
	"github.com/sirupsen/logrus"
)

var CveConfigInstance CVEConfig

type CVEConfig struct {
	Kafka kafka.ConsumeConfig `yaml:"kafka"`
}

func InitCVEConfig(configFile string) {
	cfg := new(CVEConfig)
	if err := utils.LoadFromYaml(configFile, cfg); err != nil {
		logrus.Error("Config初始化失败, err:", err)
		return
	}
	CveConfigInstance = *cfg
}
