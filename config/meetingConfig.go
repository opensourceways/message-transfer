package config

import (
	"github.com/opensourceways/message-transfer/common/kafka"
	"github.com/opensourceways/message-transfer/utils"
	"github.com/sirupsen/logrus"
)

var MeetingConfigInstance MeetingConfig

type MeetingConfig struct {
	Kafka kafka.ConsumeConfig `yaml:"kafka"`
}

func InitMeetingConfig(configFile string) {
	cfg := new(MeetingConfig)
	if err := utils.LoadFromYaml(configFile, cfg); err != nil {
		logrus.Error("Config初始化失败, err:", err)
		return
	}
	MeetingConfigInstance = *cfg
}
