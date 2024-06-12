package config

import (
	"github.com/opensourceways/message-transfer/common/kafka"
	"github.com/opensourceways/message-transfer/utils"
	"github.com/sirupsen/logrus"
)

var GiteeConfigInstance GiteeConfig

type GiteeConfig struct {
	Issue kafka.ConsumeConfig `yaml:"issue"`
	Push  kafka.ConsumeConfig `yaml:"push"`
	PR    kafka.ConsumeConfig `yaml:"pr"`
	Note  kafka.ConsumeConfig `yaml:"note"`
}

func InitGiteeConfig(configFile string) {
	cfg := new(GiteeConfig)
	if err := utils.LoadFromYaml(configFile, cfg); err != nil {
		logrus.Error("Config初始化失败, err:", err)
		return
	}
	GiteeConfigInstance = *cfg
}
