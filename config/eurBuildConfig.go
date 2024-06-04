package config

import (
	"github.com/opensourceways/message-transfer/common/kafka"
	"github.com/opensourceways/message-transfer/utils"
	"github.com/sirupsen/logrus"
)

var EurBuildConfigInstance EurBuildConfig

type EurBuildConfig struct {
	Kafka kafka.ConsumeConfig `yaml:"kafka"`
}

func InitEurBuildConfig(configFile string) {
	cfg := new(EurBuildConfig)
	if err := utils.LoadFromYaml(configFile, cfg); err != nil {
		logrus.Error("Config初始化失败, err:", err)
		return
	}
	EurBuildConfigInstance = *cfg
}
