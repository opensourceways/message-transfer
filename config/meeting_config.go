/*
Copyright (c) Huawei Technologies Co., Ltd. 2024. All rights reserved
*/

package config

import (
	"github.com/opensourceways/message-transfer/common/kafka"
	"github.com/opensourceways/message-transfer/utils"
	"github.com/sirupsen/logrus"
)

// MeetingConfigInstance meeting config instance.
var MeetingConfigInstance MeetingConfig

// MeetingConfig definition of meeting config.
type MeetingConfig struct {
	Kafka kafka.ConsumeConfig `yaml:"kafka"`
}

// InitMeetingConfig init meeting config.
func InitMeetingConfig(configFile string) {
	cfg := new(MeetingConfig)
	if err := utils.LoadFromYaml(configFile, cfg); err != nil {
		logrus.Error("Config初始化失败, err:", err)
		return
	}
	MeetingConfigInstance = *cfg
}
