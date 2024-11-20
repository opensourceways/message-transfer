package config

import (
	"github.com/sirupsen/logrus"

	"github.com/opensourceways/message-transfer/common/kafka"
	"github.com/opensourceways/message-transfer/utils"
)

// EurBuildConfigInstance eur build config instance.
var EurBuildConfigInstance EurBuildConfig

// EurBuildConfig definition of eur build config.
type EurBuildConfig struct {
	Kafka kafka.ConsumeConfig `yaml:"kafka"`
}

// InitEurBuildConfig init eur build config.
func InitEurBuildConfig(configFile string) {
	cfg := new(EurBuildConfig)
	if err := utils.LoadFromYaml(configFile, cfg); err != nil {
		logrus.Error("Config初始化失败, err:", err)
		return
	}
	EurBuildConfigInstance = *cfg
}
