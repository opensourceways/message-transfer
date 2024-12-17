/*
Copyright (c) Huawei Technologies Co., Ltd. 2024. All rights reserved
*/

package config

import (
	"github.com/sirupsen/logrus"

	"github.com/opensourceways/message-transfer/common/kafka"
	"github.com/opensourceways/message-transfer/utils"
)

// PublishConfigInstance Publish config instance.
var PublishConfigInstance PublishConfig

// PublishConfig definition of Publish config.
type PublishConfig struct {
	Kafka kafka.ConsumeConfig `yaml:"kafka"`
}

// InitPublishConfig init Publish config.
func InitPublishConfig(configFile string) {
	cfg := new(PublishConfig)
	if err := utils.LoadFromYaml(configFile, cfg); err != nil {
		logrus.Error("Config初始化失败, err:", err)
		return
	}
	PublishConfigInstance = *cfg
}
