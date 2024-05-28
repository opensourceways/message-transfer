package config

import (
	"github.com/opensourceways/message-transfer/common/kafka"
	"github.com/opensourceways/message-transfer/utils"
	"github.com/sirupsen/logrus"
)

var EurBuildConfigInstance EurBuildConfig

type EurBuildConfig struct {
	Consume kafka.ConsumeConfig `yaml:"consume"`
	Publish string              `yaml:"publish"`
}

func InitEurBuildConfig() {
	cfg := new(EurBuildConfig)
	if err := utils.LoadFromYaml("config/eur_build_conf.yaml", cfg); err != nil {
		logrus.Error("Config初始化失败, err:", err)
		return
	}
	EurBuildConfigInstance = *cfg
}
